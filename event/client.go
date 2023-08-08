package event

import (
	"errors"
	"fmt"
	"net"
	"os"
)

type RegFunc func(Data)
type Client struct {
	socket net.Conn
	reg    map[Event]map[string]RegFunc
}

func NewClient() (c *Client, err error) {
	his := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if his == "" {
		err = errors.New("failed to get HIS (is Hyprland runnning?)")
		return
	}
	listener, e := net.Listen("unix", "/tmp/hypr/"+his+"/.socket2.sock")
	if e != nil {
		err = errors.New("failed to listen Hyprland socket2: " + e.Error())
		return
	}
	socket, e := listener.Accept()
	if e != nil {
		err = errors.New("failed to connect to Hyprland socket2: " + e.Error())
		return
	}
	c = &Client{socket, make(map[Event]map[string]RegFunc)}
	return
}

func (c *Client) RegisterEvent(event Event, f RegFunc, name string) error {
	_, ok := c.reg[event][name]
	if ok {
		return fmt.Errorf("func \"%s\" already registered", name)
	}
	c.reg[event][name] = f
	return nil
}

func (c *Client) UnregisterEvent(event Event, name string) error {
	delete(c.reg[event], name)
	return nil
}
