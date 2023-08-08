package hyprctl

import "strconv"

func (c *Client) SetCursor(name string, size int) error {
	cmd := c.newCmd("keyword", "", []string{name, strconv.FormatInt(int64(size), 10)})
	err := cmd.err()
	return err
}
