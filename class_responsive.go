package bulma

type ResponsiveClass string

func (c ResponsiveClass) Class() Class {
	return Class(c)
}

// Mobile makes the class apply on mobile screens.
func (c ResponsiveClass) Mobile() Class {
	return Class(c) + "-mobile"
}

// Tablet makes the class apply on tablet and larger screens.
func (c ResponsiveClass) Tablet() Class {
	return Class(c) + "-tablet"
}

// TabletOnly makes the class apply on tablet screens only.
func (c ResponsiveClass) TabletOnly() Class {
	return Class(c) + "-tablet-only"
}

// Touch makes the class apply on touch screens (mobile and tablet) only.
func (c ResponsiveClass) Touch() Class {
	return Class(c) + "-touch"
}

// Desktop makes the class apply on desktop and larger screens.
func (c ResponsiveClass) Desktop() Class {
	return Class(c) + "-desktop"
}

// DesktopOnly makes the class apply on desktop screens only.
func (c ResponsiveClass) DesktopOnly() Class {
	return Class(c) + "-desktop-only"
}

// Widescreen makes the class apply on widescreen and larger screens.
func (c ResponsiveClass) Widescreen() Class {
	return Class(c) + "-widescreen"
}

// WidescreenOnly makes the class apply on widescreen screens only.
func (c ResponsiveClass) WidescreenOnly() Class {
	return Class(c) + "-widescreen-only"
}

// FullHD makes the class apply on full hd screens only.
func (c ResponsiveClass) FullHD() Class {
	return Class(c) + "-fullhd"
}

type ResponsiveClasses struct {
	Responsive    []string
	NonResponsive []string
}

func (c ResponsiveClasses) Classes() []Class {
	classes := make([]Class, 0, len(c.Responsive)+len(c.NonResponsive))
	for _, cl := range c.Responsive {
		classes = append(classes, Class(cl))
	}
	for _, cl := range c.NonResponsive {
		classes = append(classes, Class(cl))
	}
	return classes
}

func (c ResponsiveClasses) responsify(suffix string) Classeser {
	classes := make(Classes, 0, len(c.Responsive)+len(c.NonResponsive))
	for _, cl := range c.Responsive {
		classes = append(classes, Class(cl+suffix))
	}
	for _, cl := range c.NonResponsive {
		classes = append(classes, Class(cl))
	}
	return classes
}

// Mobile makes the class apply on mobile screens.
func (c ResponsiveClasses) Mobile() Classeser {
	return c.responsify("-mobile")
}

// Tablet makes the class apply on tablet and larger screens.
func (c ResponsiveClasses) Tablet() Classeser {
	return c.responsify("-tablet")
}

// TabletOnly makes the class apply on tablet screens only.
func (c ResponsiveClasses) TabletOnly() Classeser {
	return c.responsify("-tablet-only")
}

// Touch makes the class apply on touch screens (mobile and tablet) only.
func (c ResponsiveClasses) Touch() Classeser {
	return c.responsify("-touch")
}

// Desktop makes the class apply on desktop and larger screens.
func (c ResponsiveClasses) Desktop() Classeser {
	return c.responsify("-desktop")
}

// DesktopOnly makes the class apply on desktop screens only.
func (c ResponsiveClasses) DesktopOnly() Classeser {
	return c.responsify("-desktop-only")
}

// Widescreen makes the class apply on widescreen and larger screens.
func (c ResponsiveClasses) Widescreen() Classeser {
	return c.responsify("-widescreen")
}

// WidescreenOnly makes the class apply on widescreen screens only.
func (c ResponsiveClasses) WidescreenOnly() Classeser {
	return c.responsify("-widescreen-only")
}

// FullHD makes the class apply on full hd screens only.
func (c ResponsiveClasses) FullHD() Classeser {
	return c.responsify("-fullhd")
}
