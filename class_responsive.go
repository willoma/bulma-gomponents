package bulma

// Mobile makes the class apply on mobile screens
// (only for classes supporting responsiveness).
func (c Class) Mobile() Class {
	return c + "-mobile"
}

// Tablet makes the class apply on tablet and larger screens
// (only for classes supporting responsiveness).
func (c Class) Tablet() Class {
	return c + "-tablet"
}

// TabletOnly makes the class apply on tablet screens only
// (only for classes supporting responsiveness).
func (c Class) TabletOnly() Class {
	return c + "-tablet-only"
}

// Touch makes the class apply on touch screens (mobile and tablet) only
// (only for classes supporting responsiveness).
func (c Class) Touch() Class {
	return c + "-touch"
}

// Desktop makes the class apply on desktop and larger screens
// (only for classes supporting responsiveness).
func (c Class) Desktop() Class {
	return c + "-desktop"
}

// DesktopOnly makes the class apply on desktop screens only
// (only for classes supporting responsiveness).
func (c Class) DesktopOnly() Class {
	return c + "-desktop-only"
}

// Widescreen makes the class apply on widescreen and larger screens
// (only for classes supporting responsiveness).
func (c Class) Widescreen() Class {
	return c + "-widescreen"
}

// WidescreenOnly makes the class apply on widescreen screens only
// (only for classes supporting responsiveness).
func (c Class) WidescreenOnly() Class {
	return c + "-widescreen-only"
}

// FullHD makes the class apply on full hd screens only
// (only for classes supporting responsiveness).
func (c Class) FullHD() Class {
	return c + "-fullhd"
}

// Mobile makes the class apply on mobile and larger screens
// (only for classes supporting responsiveness).
func (c MultiClass) Mobile() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-mobile")
	}
	return newC
}

// Tablet makes the class apply on tablet and larger screens
// (only for classes supporting responsiveness).
func (c MultiClass) Tablet() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-tablet")
	}
	return newC
}

// TabletOnly makes the class apply on tablet screens only
// (only for classes supporting responsiveness).
func (c MultiClass) TabletOnly() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-tablet-only")
	}
	return newC
}

// Touch makes the class apply on touch screens (mobile and tablet) only
// (only for classes supporting responsiveness).
func (c MultiClass) Touch() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-touch")
	}
	return newC
}

// Desktop makes the class apply on desktop and larger screens
// (only for classes supporting responsiveness).
func (c MultiClass) Desktop() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-desktop")
	}
	return newC
}

// DesktopOnly makes the class apply on desktop screens only
// (only for classes supporting responsiveness).
func (c MultiClass) DesktopOnly() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-desktop-only")
	}
	return newC
}

// Widescreen makes the class apply on widescreen and larger screens
// (only for classes supporting responsiveness).
func (c MultiClass) Widescreen() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-widescreen")
	}
	return newC
}

// WidescreenOnly makes the class apply on widescreen screens only
// (only for classes supporting responsiveness).
func (c MultiClass) WidescreenOnly() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-widescreen-only")
	}
	return newC
}

// FullHD makes the class apply on full hd screens only
// (only for classes supporting responsiveness).
func (c MultiClass) FullHD() MultiClass {
	newC := MultiClass{
		Static: c.Static,
	}
	for _, rc := range c.Responsive {
		newC.Static = append(newC.Static, rc+"-fullhd")
	}
	return newC
}
