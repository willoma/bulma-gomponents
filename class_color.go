package bulma

// ColorClass defines a color for an element.
type ColorClass struct {
	class string
	light bool
}

var (
	// Base colors
	White   = ColorClass{"white", false}
	Black   = ColorClass{"black", false}
	Light   = ColorClass{"light", false}
	Dark    = ColorClass{"dark", false}
	Primary = ColorClass{"primary", false}
	Link    = ColorClass{"link", false}
	Info    = ColorClass{"info", false}
	Success = ColorClass{"success", false}
	Warning = ColorClass{"warning", false}
	Danger  = ColorClass{"danger", false}

	// Shades of grey
	BlackBis    = ColorClass{"black-bis", false}
	BlackTer    = ColorClass{"black-ter", false}
	GreyDarker  = ColorClass{"grey-darker", false}
	GreyDark    = ColorClass{"grey-dark", false}
	Grey        = ColorClass{"grey", false}
	GreyLight   = ColorClass{"grey-light", false}
	GreyLighter = ColorClass{"grey-lighter", false}
	WhiteTer    = ColorClass{"white-ter", false}
	WhiteBis    = ColorClass{"white-bis", false}

	// Light variants
	PrimaryLight = ColorClass{"primary", true}
	LinkLight    = ColorClass{"link", true}
	InfoLight    = ColorClass{"info", true}
	SuccessLight = ColorClass{"success", true}
	WarningLight = ColorClass{"warning", true}
	DangerLight  = ColorClass{"danger", true}
)

// Text returns the "has-text-*" version of the color
func (c ColorClass) Text() Class {
	cl := Class("has-text-" + c.class)

	if c.light {
		return cl + "-light"
	}

	return cl
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

	if c.light {
		return cl + "-light"
	}

	return cl
}

// BackgroundLight returns the "has-background-*-light" version of the color
func (c ColorClass) BackgroundLight() Class {
	return Class("has-background-" + c.class + "-light")
}

// BackgroundDark returns the "has-background-*-dark" version of the color
func (c ColorClass) BackgroundDark() Class {
	return Class("has-background-" + c.class + "-dark")
}
