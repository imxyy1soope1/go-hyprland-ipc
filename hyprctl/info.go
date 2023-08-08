package hyprctl

func (c *Client) InfoVersion() (res string, err error) {
	cmd := c.newCmd("version", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoMonitors() (res string, err error) {
	cmd := c.newCmd("monitors", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoWorkspaces() (res string, err error) {
	cmd := c.newCmd("workspaces", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoActiveWorkspace() (res string, err error) {
	cmd := c.newCmd("activeworkspace", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoClients() (res string, err error) {
	cmd := c.newCmd("clients", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoDevices() (res string, err error) {
	cmd := c.newCmd("devices", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoInfBinds() (res string, err error) {
	cmd := c.newCmd("binds", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoActiveWindow() (res string, err error) {
	cmd := c.newCmd("activewindow", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoLayers() (res string, err error) {
	cmd := c.newCmd("layers", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoSplash() (res string, err error) {
	cmd := c.newCmd("splash", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoGetOption(opt string) (res string, err error) {
	cmd := c.newCmd("getoption", "", []string{opt})
	res, err = cmd.res()
	return
}

func (c *Client) InfoCursorPos() (res string, err error) {
	cmd := c.newCmd("version", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoAnimations() (res string, err error) {
	cmd := c.newCmd("animations", "", []string{})
	res, err = cmd.res()
	return
}

func (c *Client) InfoInstances() (res string, err error) {
	cmd := c.newCmd("instances", "", []string{})
	res, err = cmd.res()
	return
}
