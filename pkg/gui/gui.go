package gui

import (
	"fmt"
	"grpchat/pkg/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Client struct {
	a        fyne.App
	w        fyne.Window
	i        *widget.Entry
	Ctrl     controller.Sendable
	messages []fyne.CanvasObject
	username string
}

func (c *Client) Run() {
	c.a = app.New()
	c.w = c.a.NewWindow("Chat")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter username")
	onLogin := func() {
		c.username = input.Text
		go c.Ctrl.Run(c.username)
		c.i = widget.NewEntry()
		c.i.SetPlaceHolder("Enter text...")
		c.i.OnSubmitted = func(string) {
			c.Ctrl.Send(c.i.Text)
			c.i.Text = ""
			c.i.Refresh()
		}
		c.w.SetContent(container.NewVBox(append(c.messages, c.i)...))
	}
	submit := widget.NewButton("Login", onLogin)

	input.OnSubmitted = func(string) {
		onLogin()
	}

	c.w.SetContent(container.NewVBox(input, submit))
	c.w.ShowAndRun()
}

func (c *Client) Update(text string, sender string) error {
	c.messages = append(c.messages, widget.NewLabel(fmt.Sprintf("%s: %s", sender, text)))
	c.w.SetContent(container.NewVBox(append(c.messages, c.i)...))
	return nil
}
