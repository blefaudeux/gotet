package gotet

import (
	"encoding/json"
	"github.com/blefaudeux/dealer"
)

// ---- Define needed structures

// Client is a socket client to a The Eye Tribe server
type Client struct {
	addr, port   string
	socket       dealer.Socket
	pointsToPush int
}

// Request is an interface{} serving as a JSON skeleleton for all requests
type Request struct {
	Category string   `json:"category"`
	Request  string   `json:"request"`
	Values   []string `json:"values"`
}

// RequestComp is an interface{} serving as a JSON skeleleton for all complex requests
type RequestComp struct {
	Category string         `json:"category"`
	Request  string         `json:"request"`
	Values   map[string]int `json:"values"`
}

// Response is an interface{} serving as a JSON skeleleton for all server responses
type Response struct {
	Category string                 `json:"category"`
	Request  string                 `json:"request"`
	Values   map[string]interface{} `json:"values"`
}

// ---- Helpers (private)

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

// sendReqComp sends a RequestComp-formatted json object to the server
func (c *Client) sendReqComp(cat string, req string, val map[string]int) (map[string]interface{}, error) {
	message := RequestComp{
		Category: cat,
		Request:  req,
		Values:   val}

	messB, _ := json.Marshal(message)
	if err := c.send(messB); err != nil {
		return nil, err
	}

	reply := c.readB("request", req)
	return reply, nil
}

// readB waits for a given json struct (described by a given field and value) to appear
func (c *Client) readB(field string, value string) map[string]interface{} {
	return c.socket.ReadBlock(field, value)
}

// handleReq is a proxy for the whole "send request and grab server feedback" loop
func (c *Client) handleReq(cat string, val string) (interface{}, error) {
	reply, err := c.sendReq(cat, val)

	if err != nil {
		return "", err
	}
	res := reply["values"].(map[string]interface{})

	return res[val], nil
}

// handleReqComp is a proxy for the whole "send request and grab server feedback" loop
func (c *Client) handleReqComp(cat string, req string, val map[string]int) (int, error) {
	reply, err := c.sendReqComp(cat, req, val)

	if err != nil {
		return 0, err
	}

	return int(reply["statuscode"].(float64)), nil
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

// ---- Implement all the API calls as methods

// -- Tracker

// Close client connection
func (c *Client) Close() {
	c.socket.Close()
}

// Version reports the protocol version currently used
func (c *Client) Version() (int, error) {
	res, err := c.handleReq("tracker", "version")

	if err != nil {
		return 0, err
	}

	return int(res.(float64)), nil
}

// Trackerstate reports the current tracker state
// See http://dev.theeyetribe.com/api/#tracker_state for a proper description
func (c *Client) Trackerstate() (int, error) {
	res, err := c.handleReq("tracker", "trackerstate")

	if err != nil {
		return 0, err
	}

	return int(res.(float64)), nil
}

// FrameData reports the last gaze estimation results
func (c *Client) FrameData() (map[string]interface{}, error) {
	reply, err := c.sendReq("tracker", "frame")

	if err != nil {
		return nil, err
	}

	return reply["values"].(map[string]interface{}), nil
}

// Framerate reports the current tracker framerate
func (c *Client) Framerate() (int, error) {
	res, err := c.handleReq("tracker", "framerate")

	if err != nil {
		return 0, err
	}

	return int(res.(float64)), nil
}

// IsCalibrated reports whether the server is currently calibrated
func (c *Client) IsCalibrated() (bool, error) {
	res, err := c.handleReq("tracker", "iscalibrated")

	if err != nil {
		return false, err
	}

	return res.(bool), nil
}

// IsCalibrating reports whether the server is currently calibrated
func (c *Client) IsCalibrating() (bool, error) {
	res, err := c.handleReq("tracker", "iscalibrating")

	if err != nil {
		return false, err
	}

	return res.(bool), nil
}

// CalibResult reports the calibration results
func (c *Client) CalibResult() (map[string]interface{}, error) {
	reply, err := c.sendReq("tracker", "calibresult")

	if err != nil {
		return nil, err
	}

	return reply["values"].(map[string]interface{}), nil
}

// ScreenIndex reports the index of the screen currently in use
func (c *Client) ScreenIndex() (int, error) {
	res, err := c.handleReq("tracker", "screenindex")

	if err != nil {
		return 0, err
	}

	return int(res.(float64)), nil
}

// ScreenResH reports the number of vertical lines on the screen
func (c *Client) ScreenResH() (int, error) {
	res, err := c.handleReq("tracker", "screenresh")

	if err != nil {
		return 0, err
	}

	return int(res.(float64)), nil
}

// ScreenResW reports the number of horizontal lines on the screen
func (c *Client) ScreenResW() (int, error) {
	res, err := c.handleReq("tracker", "screenresw")

	if err != nil {
		return 0, err
	}

	return int(res.(float64)), nil
}

// ScreenPsyW reports the physical (meters) width of the screen
func (c *Client) ScreenPsyW() (float64, error) {
	res, err := c.handleReq("tracker", "screenpsyh")

	if err != nil {
		return 0, err
	}

	return res.(float64), nil
}

// ScreenPsyH reports the physical (meters) height of the screen
func (c *Client) ScreenPsyH() (float64, error) {
	res, err := c.handleReq("tracker", "screenpsyh")

	if err != nil {
		return 0, err
	}

	return res.(float64), nil
}

// -- Calibration

// CalibStart prepares the tracker for a new calibration
func (c *Client) CalibStart(pts int) (bool, error) {
	values := map[string]int{"pointcount": pts}

	res, err := c.handleReqComp("calibration", "start", values)

	if err != nil {
		return false, err
	}

	c.pointsToPush = pts
	return res == 200, nil
}

// CalibPointStart sends the coordinate of a new point
func (c *Client) CalibPointStart(ptx, pty int) (bool, error) {
	values := map[string]int{"x": ptx, "y": pty}

	res, err := c.handleReqComp("calibration", "pointstart", values)

	if err != nil {
		return false, err
	}

	return res == 200, nil
}

// CalibPointEnd ends the calibration process for the current point
func (c *Client) CalibPointEnd() (bool, error) {
	res, err := c.handleReqComp("calibration", "pointend", map[string]int{})
	if err != nil {
		return false, err
	}

	c.pointsToPush--

	// Calibration sequence is done, grab the result
	if c.pointsToPush == 0 {
		calibRes := c.readB("tracker", "calibresult")
		return calibRes["result"].(bool), nil
	}

	return res == 200, nil
}

// CalibAbort aborts the calibration process
func (c *Client) CalibAbort() (bool, error) {
	res, err := c.handleReqComp("calibration", "abort", map[string]int{})

	c.pointsToPush = 0

	if err != nil {
		return false, err
	}

	return res == 200, nil
}

// CalibClear clear the server calibration status
func (c *Client) CalibClear() (bool, error) {
	res, err := c.handleReqComp("calibration", "clear", map[string]int{})

	if err != nil {
		return false, err
	}

	return res == 200, nil
}
