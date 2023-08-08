package hyprctl

import "fmt"

type Prop interface {
	getProp() string
}

type PropAnimationStyle string

func (p PropAnimationStyle) getProp() string {
	return fmt.Sprintf("animationstyle %s", p)
}

type PropRounding int

func (p PropRounding) getProp() string {
	return fmt.Sprintf("rounding %d", p)
}

type PropBorderSize int

func (p PropBorderSize) getProp() string {
	return fmt.Sprintf("bordersize %d", p)
}

type PropForceNoBlur bool

func (p PropForceNoBlur) getProp() string {
	if p {
		return "forcenoblur 1"
	} else {
		return "forcenoblur 0"
	}
}

type PropForceOpaque bool

func (p PropForceOpaque) getProp() string {
	if p {
		return "forceopaque 1"
	} else {
		return "forceopaque 0"
	}
}

type PropForceOpaqueOverriden bool

func (p PropForceOpaqueOverriden) getProp() string {
	if p {
		return "forcenobluroverriden 1"
	} else {
		return "forcenobluroverriden 0"
	}
}

type PropForceAllowsInput bool

func (p PropForceAllowsInput) getProp() string {
	if p {
		return "forceallowsinput 1"
	} else {
		return "forceallowsinput 0"
	}
}

type PropForceNoAnims bool

func (p PropForceNoAnims) getProp() string {
	if p {
		return "forcenoanims 1"
	} else {
		return "forcenoanims 0"
	}
}

type PropForceNoBorder bool

func (p PropForceNoBorder) getProp() string {
	if p {
		return "forcenoborder 1"
	} else {
		return "forcenoborder 0"
	}
}

type PropForceNoDim bool

func (p PropForceNoDim) getProp() string {
	if p {
		return "forcenodim 1"
	} else {
		return "forcenodim 0"
	}
}

type PropForceNoShadow bool

func (p PropForceNoShadow) getProp() string {
	if p {
		return "forcenoshadow 1"
	} else {
		return "forcenoshadow 0"
	}
}

type PropWindowDanceCompat bool

func (p PropWindowDanceCompat) getProp() string {
	if p {
		return "windowdancecompat 1"
	} else {
		return "windowdancecompat 0"
	}
}

type PropNoMaxSize bool

func (p PropNoMaxSize) getProp() string {
	if p {
		return "nomaxsize 1"
	} else {
		return "nomaxsize 0"
	}
}

type PropDimAround bool

func (p PropDimAround) getProp() string {
	if p {
		return "dimaround 1"
	} else {
		return "dimaround 0"
	}
}

type PropKeepAspectRatio bool

func (p PropKeepAspectRatio) getProp() string {
	if p {
		return "keepaspectratio 1"
	} else {
		return "keepaspectratio 0"
	}
}

type PropAlphaOverride bool

func (p PropAlphaOverride) getProp() string {
	if p {
		return "alphaoverride 1"
	} else {
		return "alphaoverride 0"
	}
}

type PropAlpha float64

func (p PropAlpha) getProp() string {
	return fmt.Sprintf("alpha %f", p)
}

type PropAlphaInactiveOverride bool

func (p PropAlphaInactiveOverride) getProp() string {
	if p {
		return "alphainactiveoverride 1"
	} else {
		return "alphainactiveoverride 0"
	}
}

type PropAlphaInactive float64

func (p PropAlphaInactive) getProp() string {
	return fmt.Sprintf("alphainactive %f", p)
}

type PropActiveBorderColor struct {
	Color
	NotSet bool
}

func (p PropActiveBorderColor) getProp() string {
	if p.NotSet {
		return "activebordercolor -1"
	} else {
		return fmt.Sprintf("activebordercolor %s", p.Color.getColor())
	}
}

type PropInactiveBorderColor struct {
	Color
	NotSet bool
}

func (p PropInactiveBorderColor) getProp() string {
	if p.NotSet {
		return "inactivebordercolor -1"
	} else {
		return fmt.Sprintf("inactivebordercolor %s", p.Color.getColor())
	}
}

func (c *Client) SetProp(window Window, prop Prop) error {
	cmd := c.newCmd("setprop", "", []string{window.getWindow(), prop.getProp()})
	err := cmd.err()
	return err
}
