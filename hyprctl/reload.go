package hyprctl

func (c *Client) Reload() error {
	cmd := c.newCmd("reload", "", []string{})
	err := cmd.err()
	return err
}
