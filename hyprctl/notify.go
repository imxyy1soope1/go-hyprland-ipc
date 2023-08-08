package hyprctl

import (
	"strconv"
	"time"
)

type Icon string

func (i Icon) getIcon() string {
	if i != "" {
		return string(i)
	} else {
		return "-1"
	}
}

func (c *Client) Notify(icon Icon, t time.Duration, color Color, msg string) error {
	cmd := c.newCmd("notify", "", []string{
		icon.getIcon(),
		strconv.FormatInt(t.Microseconds(), 10),
		color.getColor(),
		msg,
	})
	err := cmd.err()
	return err
}
