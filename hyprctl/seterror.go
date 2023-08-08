package hyprctl

import "fmt"

type Color interface {
	getColor() string
}

type RGB struct {
	R, G, B int
}

func (r RGB) getColor() string {
	return fmt.Sprintf("rgb(%x%x%x)", r.R, r.G, r.B)
}

type RGBA struct {
	R, G, B, A int
}

func (r RGBA) getColor() string {
	return fmt.Sprintf("rgb(%x%x%x%x)", r.R, r.G, r.B, r.A)
}

type EmptyColor struct {
}

func (c EmptyColor) getColor() string {
	return ""
}

func (c *Client) SetError(color Color, msg string) error {
	cmd := c.newCmd("seterror", "", []string{color.getColor(), msg})
	err := cmd.err()
	return err
}

func (c *Client) SetErrorDisable() error {
	cmd := c.newCmd("seterror", "", []string{"disable"})
	err := cmd.err()
	return err
}
