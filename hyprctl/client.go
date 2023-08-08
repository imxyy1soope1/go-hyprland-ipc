package hyprctl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type Client struct {
	socket net.Conn
}
type cmd struct {
	client        *Client
	command, flag string
	args          []string
}

func NewClient() (c *Client, err error) {
	his := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if his == "" {
		err = errors.New("failed to get HIS (is Hyprland runnning?)")
		return
	}
	listener, e := net.Listen("unix", "/tmp/hypr/"+his+"/.socket.sock")
	if e != nil {
		err = errors.New("failed to listen socket: " + e.Error())
		return
	}
	socket, e := listener.Accept()
	if e != nil {
		err = errors.New("failed to connect to socket: " + e.Error())
		return
	}
	c = &Client{socket}
	return
}

func (c *Client) newCmd(command, flag string, args []string) *cmd {
	return &cmd{
		client:  c,
		command: command,
		flag:    flag,
		args:    args,
	}
}

func (r cmd) res() (response string, err error) {
	writer := bufio.NewWriter(r.client.socket)
	_, err = writer.Write(([]byte)(fmt.Sprintf("%s/%s %s\n", r.flag, r.command, strings.Join(r.args, " "))))
	if err != nil {
		return
	}
	var res []byte
	reader := bufio.NewReaderSize(r.client.socket, 8192)
	var e error
	var buf []byte
	for e != io.EOF {
		if e != bufio.ErrBufferFull {
			err = e
			break
		}
		buf, e = reader.ReadSlice('\n')
		res = append(res, buf...)
	}
	if err != nil {
		return
	}
	response = string(res)
	return
}

func (r cmd) err() error {
	res, err := r.res()
	if res != "" && res != "ok" {
		err = errors.New(res)
	}
	return err
}
