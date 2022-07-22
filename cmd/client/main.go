package main

import (
	"grpchat/pkg/client"
	"grpchat/pkg/gui"
)

func main() {
	gui := &gui.Client{}
	c := &client.Client{
		Ctrl: gui,
	}
	gui.Ctrl = c
	gui.Run()
}
