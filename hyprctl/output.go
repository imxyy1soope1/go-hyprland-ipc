package hyprctl

type OutputBackend string

const (
	OutputWayland  OutputBackend = "wayland"
	OutputX11      OutputBackend = "x11"
	OutputHeadless OutputBackend = "headless"
	OutputAuto     OutputBackend = "auto"
)

func (c *Client) OutputCreate(backend OutputBackend) error {
	cmd := c.newCmd("output", "", []string{"create", string(backend)})
	err := cmd.err()
	return err
}

func (c *Client) OutputRemove(name string) error {
	cmd := c.newCmd("output", "", []string{"remove", name})
	err := cmd.err()
	return err
}
