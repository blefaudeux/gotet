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

func (c *Client) send(message []byte) error {
	if err := c.socket.SendBytes(message); err != nil {
		return err
	}
	return nil
}

// sendReq sends a Request-formatted json object to the server
func (c *Client) sendReq(cat string, val string) (map[string]interface{}, error) {
	req := Request{
		Category: cat,
		Request:  "get",
		Values:   []string{val}}

	reqb, _ := json.Marshal(req)
	if err := c.send(reqb); err != nil {
		return nil, err
	}

	reply := c.readB("request", "get")
	return reply, nil
}

func (c *Client) readB(field string, value string) map[string]interface{} {
	return c.socket.ReadBlock(field, value)
}

// Version reports the protocol version currently used
func (c *Client) Version() (int, error) {
	reply, err := c.sendReq("tracker", "version")
	if err != nil {
		return 0, err
	}

	vers := reply["values"].(map[string]interface{})
	return int(vers["version"].(float64)), nil
}

// IsCalibrated reports whether the server is currently calibrated
func (c *Client) IsCalibrated() (bool, error) {
	reply, err := c.sendReq("tracker", "iscalibrated")
	if err != nil {
		return false, err
	}

	vers := reply["values"].(map[string]interface{})
	return vers["iscalibrated"].(bool), nil
}

// IsCalibrating reports whether the server is currently calibrated
func (c *Client) IsCalibrating() (bool, error) {
	reply, err := c.sendReq("tracker", "iscalibrating")
	if err != nil {
		return false, err
	}

	vers := reply["values"].(map[string]interface{})
	return vers["iscalibrating"].(bool), nil
}
