package bulma

import e "github.com/willoma/gomplements"

type ResponsiveClass string

func (c ResponsiveClass) Class() e.Class {
	return e.Class(c)
}

func (c ResponsiveClass) If(cond bool) ResponsiveClass {
	if cond {
		return c
	}

	return ""
}

// Mobile makes the class apply on mobile screens.
func (c ResponsiveClass) Mobile() e.Class {
	return e.Class(c) + "-mobile"
}

// Tablet makes the class apply on tablet and larger screens.
func (c ResponsiveClass) Tablet() e.Class {
	return e.Class(c) + "-tablet"
}

// TabletOnly makes the class apply on tablet screens only.
func (c ResponsiveClass) TabletOnly() e.Class {
	return e.Class(c) + "-tablet-only"
}

// Touch makes the class apply on touch screens (mobile and tablet) only.
func (c ResponsiveClass) Touch() e.Class {
	return e.Class(c) + "-touch"
}

// Desktop makes the class apply on desktop and larger screens.
func (c ResponsiveClass) Desktop() e.Class {
	return e.Class(c) + "-desktop"
}

// DesktopOnly makes the class apply on desktop screens only.
func (c ResponsiveClass) DesktopOnly() e.Class {
	return e.Class(c) + "-desktop-only"
}

// Widescreen makes the class apply on widescreen and larger screens.
func (c ResponsiveClass) Widescreen() e.Class {
	return e.Class(c) + "-widescreen"
}

// WidescreenOnly makes the class apply on widescreen screens only.
func (c ResponsiveClass) WidescreenOnly() e.Class {
	return e.Class(c) + "-widescreen-only"
}

// FullHD makes the class apply on full hd screens only.
func (c ResponsiveClass) FullHD() e.Class {
	return e.Class(c) + "-fullhd"
}
