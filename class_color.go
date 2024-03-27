package bulma

type colorVariant int

const (
	colorDefault colorVariant = iota
	colorLight
	colorDark
)

// ColorClass defines a color for an element.
type ColorClass struct {
	class   string
	variant colorVariant
}

var (
	// Base colors
	White   = ColorClass{class: "white"}
	Black   = ColorClass{class: "black"}
	Light   = ColorClass{class: "light"}
	Dark    = ColorClass{class: "dark"}
	Primary = ColorClass{class: "primary"}
	Link    = ColorClass{class: "link"}
	Info    = ColorClass{class: "info"}
	Success = ColorClass{class: "success"}
	Warning = ColorClass{class: "warning"}
	Danger  = ColorClass{class: "danger"}

	// Shades of grey
	BlackBis    = ColorClass{class: "black-bis"}
	BlackTer    = ColorClass{class: "black-ter"}
	GreyDarker  = ColorClass{class: "grey-darker"}
	GreyDark    = ColorClass{class: "grey-dark"}
	Grey        = ColorClass{class: "grey"}
	GreyLight   = ColorClass{class: "grey-light"}
	GreyLighter = ColorClass{class: "grey-lighter"}
	WhiteTer    = ColorClass{class: "white-ter"}
	WhiteBis    = ColorClass{class: "white-bis"}

	// Light variants
	PrimaryLight = ColorClass{class: "primary", variant: colorLight}
	LinkLight    = ColorClass{class: "link", variant: colorLight}
	InfoLight    = ColorClass{class: "info", variant: colorLight}
	SuccessLight = ColorClass{class: "success", variant: colorLight}
	WarningLight = ColorClass{class: "warning", variant: colorLight}
	DangerLight  = ColorClass{class: "danger", variant: colorLight}

	// Dark variants
	PrimaryDark = ColorClass{class: "primary", variant: colorDark}
	LinkDark    = ColorClass{class: "link", variant: colorDark}
	InfoDark    = ColorClass{class: "info", variant: colorDark}
	SuccessDark = ColorClass{class: "success", variant: colorDark}
	WarningDark = ColorClass{class: "warning", variant: colorDark}
	DangerDark  = ColorClass{class: "danger", variant: colorDark}
)

// Text returns the "has-text-*" version of the color
func (c ColorClass) Text() Class {
	cl := Class("has-text-" + c.class)

	switch c.variant {
	case colorLight:
		return cl + "-light"
	case colorDark:
		return cl + "-dark"
	default:
		return cl
	}
}

// TextLight returns the "has-text-*-light" version of the color
func (c ColorClass) TextLight() Class {
	return Class("has-text-" + c.class + "-light")
}

// TextDark returns the "has-text-*-dark" version of the color
func (c ColorClass) TextDark() Class {
	return Class("has-text-" + c.class + "-dark")
}

// IconLight returns the "light" variant of the
func (c ColorClass) IconLight() Class {
	return c.TextLight()
}

func (c ColorClass) IconDark() Class {
	return c.TextDark()
}

// Background returns the "has-background-*" version of the color
func (c ColorClass) Background() Class {
	cl := Class("has-background-" + c.class)

	switch c.variant {
	case colorLight:
		return cl + "-light"
	case colorDark:
		return cl + "-dark"
	default:
		return cl
	}
}

// BackgroundLight returns the "has-background-*-light" version of the color
func (c ColorClass) BackgroundLight() Class {
	return Class("has-background-" + c.class + "-light")
}

// BackgroundDark returns the "has-background-*-dark" version of the color
func (c ColorClass) BackgroundDark() Class {
	return Class("has-background-" + c.class + "-dark")
}
