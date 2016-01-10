/*
Here we're demonstrating inline interfaces and simple mocks.

The app below depends on a client that it relies on for sending messages.  The only requirement for the client is that it must have a `Send(string) error` method.
*/
package main

import (
	"fmt"
)

// standard client
type Client struct {}

func (c *Client) Send(msg string) error {
	fmt.Printf("client is sending message %q\n", msg) 
	return nil
}

// test client
type TestClient struct {
	sendfunc func(string) error
}

func (c *TestClient) Send(msg string) error {
	return c.sendfunc(msg)
}

// our main app
type App struct {
	client interface {
		Send(string) error
	}
}

func (a *App) Send(msg string) error {
	return a.client.Send(msg)
}

func main() {
	client := new(Client)
	app := &App{client}
	app.Send("hi!")

	var got string
	send := func(msg string) error {
		got = msg
		fmt.Printf("test client is now sending message %q\n", msg)
		return nil
	}
	want := "bye!"
	app.client = &TestClient{send}
	app.Send(want)

	fmt.Println(got == want)
}
