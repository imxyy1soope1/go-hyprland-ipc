package hyprctl

import "strconv"

type XkbLayout interface {
	getLayout() string
}

type XkbLayoutNext struct {
}

func (x XkbLayoutNext) getLayout() string {
	return "next"
}

type XkbLayoutPrev struct {
}

func (x XkbLayoutPrev) getLayout() string {
	return "prev"
}

type XkbLayoutWithId struct {
	id uint
}

func (x XkbLayoutWithId) getLayout() string {
	return strconv.FormatUint(uint64(x.id), 10)
}

func (c *Client) SwitchXkbLayout(device string, layout XkbLayout) error {
	cmd := c.newCmd("switchxkblayout", "", []string{device, layout.getLayout()})
	err := cmd.err()
	return err
}
