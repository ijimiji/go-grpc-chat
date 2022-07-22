package server

import (
	"context"
	"grpchat/pkg/pb"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type Connection struct {
	stream pb.ChatService_CreateStreamServer
	id     string
	active bool
	error  chan error
}

type Server struct {
	pb.UnimplementedChatServiceServer
	Connection []*Connection
}

func (s *Server) CreateStream(connection *pb.Connect, stream pb.ChatService_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     connection.User.Name,
		active: true,
		error:  make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *pb.Message) (*pb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *pb.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)

				if err != nil {
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)

	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &pb.Close{}, nil
}

func Run() {
	var connections []*Connection

	server := &Server{
		Connection: connections,
	}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	pb.RegisterChatServiceServer(grpcServer, server)
	grpcServer.Serve(listener)
}
