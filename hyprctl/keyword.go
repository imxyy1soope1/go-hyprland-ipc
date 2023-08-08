package hyprctl

func (c *Client) KeywordSet(key, value string) error {
	cmd := c.newCmd("keyword", "", []string{key, value})
	err := cmd.err()
	return err
}
