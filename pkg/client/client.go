package client

import (
	"flag"
	"fmt"
	"grpchat/pkg/controller"
	"grpchat/pkg/pb"
	"log"

	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	Ctrl     controller.Updatable
	client   pb.ChatServiceClient
	wait     *sync.WaitGroup
	username string
}

func (c *Client) init() {
	c.wait = &sync.WaitGroup{}
}

func (c *Client) connect(user *pb.User) error {
	var streamerror error

	stream, err := c.client.CreateStream(context.Background(), &pb.Connect{
		User:   user,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	c.wait.Add(1)
	go func(str pb.ChatService_CreateStreamClient) {
		defer c.wait.Done()

		for {
			msg, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("error reading message: %v", err)
				break
			}
			c.Ctrl.Update(msg.Text, msg.From.Name)
		}
	}(stream)

	return streamerror
}

func (c *Client) Send(text string) error {
	_, err := c.client.BroadcastMessage(context.Background(), &pb.Message{
		Text: text,
		From: &pb.User{Name: c.username},
	})
	return err
}

func (c *Client) Run(username string) {
	c.username = username
	done := make(chan int)
	c.init()

	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldnt connect to service: %v", err)
	}

	c.client = pb.NewChatServiceClient(conn)
	user := &pb.User{
		Name: username,
	}

	c.connect(user)

	go func() {
		c.wait.Wait()
		close(done)
	}()

	<-done
}
