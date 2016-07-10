package gotet

import (
	"dealer"
	"encoding/json"
)

// Client a socket client to a The Eye Tribe server
type Client struct {
	addr, port string
	socket     dealer.Socket
}

// Connect to a running server
func (c *Client) Connect(addr, port string) error {
	c.addr = addr
	c.port = port

	if err := c.socket.Connect(addr, port); err != nil {
		return err
	}

	return nil
}

// Close client connection
func (c *Client) Close() {
	c.socket.Close()
}

// IO helpers
type Request struct {
	Category string   `json:"category"`
	Request  string   `json:"request"`
	Values   []string `json:"values"`
	// TODO: Implement JSON exports
}

type Response struct {
	Category string                 `json:"category"`
	Request  string                 `json:"request"`
	Values   map[string]interface{} `json:"values"`
	// TODO: Implement JSON exports
}

func (c *Client) send(message []byte) {
	c.socket.SendBytes(message)
}

func (c *Client) readB(field string, value string) map[string]interface{} {
	return c.socket.ReadBlock(field, value)
}

// Version reports the protocol version currently used
func (c *Client) Version() (float64, error) {
	req := Request{
		Category: "tracker",
		Request:  "get",
		Values:   []string{"version"}}

	reqb, _ := json.Marshal(req)
	c.send(reqb)

	reply := c.readB("request", "get")
	vers := reply["values"].(map[string]interface{})
	return vers["version"].(float64), nil
}
