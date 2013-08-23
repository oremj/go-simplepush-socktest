package main

import (
	"fmt"
	"log"
	"code.google.com/p/go.net/websocket"
)

type Client struct {
	conn *websocket.Conn
	uaid string
	notificationHandler func(string, int)
}

func NewClient(url string) (c *Client, err error) {
	ws, err := websocket.Dial(url, "", url)
	if err != nil {
		return
	}
	c = &Client{ws, "", nil}
	c.hello()
	c.listen()
	return
}

func (c *Client) Register(channelID string) (err error) {
	data := Message{"messageType": "register",
					"channelID": channelID}
	err = c.send(data)
	return
}

func (c *Client) send(v Message) (err error) {
	err = websocket.JSON.Send(c.conn, v)
	return
}

func (c *Client) recv() (data Message, err error) {
	err = websocket.JSON.Receive(c.conn, &data)
	return
}

func (c *Client) listen() {
	go func() {
		log.Println(c.recv())
	}()
}

func (c *Client) hello() (err error) {
	data := Message{"messageType": "hello",
					"uaid": c.uaid,
					"channelIDs": []string{}}
	err = c.send(data)
	if err != nil {
		return
	}
	resp, err := c.recv()
	if err != nil {
		return
	}
	if resp["messageType"].(string) != "hello" {
		return fmt.Errorf("MessageType was not hello: %s", resp["messageType"])
	}
	c.uaid = resp["uaid"].(string)
	return
}
