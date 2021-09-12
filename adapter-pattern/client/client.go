package client

import (
	"fmt"

	"design-pattern/adapter-pattern/service"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com service.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
