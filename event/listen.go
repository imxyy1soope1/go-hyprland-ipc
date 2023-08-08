package event

import (
	"bufio"
	"io"
	"strings"
)

func (c *Client) Listen() {
	go c.listen()
}

func (c *Client) listen() {
	reader := bufio.NewReader(c.socket)
	var line string
	for {
		buf, err := reader.ReadSlice('\n')
		if err == io.EOF {
			continue
		} else if err != nil {
			panic("read Hyprland socket2 failed")
		}
		line = strings.TrimRight(string(buf), "\n")
		parts := strings.Split(line, ">>")
		event := stringToEvent[parts[0]]
		for _, f := range c.reg[event] {
			parser := eventToParser[event]
			go f(parser(parts[1]))
		}
	}
}
