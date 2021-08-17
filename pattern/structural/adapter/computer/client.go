package computer

import "fmt"

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Icomputer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
