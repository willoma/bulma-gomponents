package bulma

import e "github.com/willoma/gomplements"

type ResponsiveClass string

func (c ResponsiveClass) Class() e.Class {
	return e.Class(c)
}

func (c ResponsiveClass) If(cond bool) e.ParentModifier {
	return &conditionalResponsiveClass{class: c, cond: cond}
}

type conditionalResponsiveClass struct {
	class ResponsiveClass
	cond  bool
}

func (c *conditionalResponsiveClass) ModifyParent(p e.Element) {
	if c.cond {
		p.With(c.class)
	}
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

type ResponsiveClasses struct {
	Responsive    []string
	NonResponsive []string
}

func (c *ResponsiveClasses) Classes() []e.Class {
	classes := make([]e.Class, 0, len(c.Responsive)+len(c.NonResponsive))
	for _, cl := range c.Responsive {
		classes = append(classes, e.Class(cl))
	}
	for _, cl := range c.NonResponsive {
		classes = append(classes, e.Class(cl))
	}
	return classes
}

func (c *ResponsiveClasses) If(cond bool) e.ParentModifier {
	return &conditionalResponsiveClasses{classes: c, cond: cond}
}

type conditionalResponsiveClasses struct {
	classes *ResponsiveClasses
	cond    bool
}

func (c *conditionalResponsiveClasses) ModifyParent(p e.Element) {
	if c.cond {
		p.With(c.classes)
	}
}

func (c *ResponsiveClasses) responsify(suffix string) e.Classeser {
	classes := make(e.Classes, 0, len(c.Responsive)+len(c.NonResponsive))
	for _, cl := range c.Responsive {
		classes = append(classes, e.Class(cl+suffix))
	}
	for _, cl := range c.NonResponsive {
		classes = append(classes, e.Class(cl))
	}
	return classes
}

// Mobile makes the class apply on mobile screens.
func (c *ResponsiveClasses) Mobile() e.Classeser {
	return c.responsify("-mobile")
}

// Tablet makes the class apply on tablet and larger screens.
func (c *ResponsiveClasses) Tablet() e.Classeser {
	return c.responsify("-tablet")
}

// TabletOnly makes the class apply on tablet screens only.
func (c *ResponsiveClasses) TabletOnly() e.Classeser {
	return c.responsify("-tablet-only")
}

// Touch makes the class apply on touch screens (mobile and tablet) only.
func (c *ResponsiveClasses) Touch() e.Classeser {
	return c.responsify("-touch")
}

// Desktop makes the class apply on desktop and larger screens.
func (c ResponsiveClasses) Desktop() e.Classeser {
	return c.responsify("-desktop")
}

// DesktopOnly makes the class apply on desktop screens only.
func (c *ResponsiveClasses) DesktopOnly() e.Classeser {
	return c.responsify("-desktop-only")
}

// Widescreen makes the class apply on widescreen and larger screens.
func (c *ResponsiveClasses) Widescreen() e.Classeser {
	return c.responsify("-widescreen")
}

// WidescreenOnly makes the class apply on widescreen screens only.
func (c *ResponsiveClasses) WidescreenOnly() e.Classeser {
	return c.responsify("-widescreen-only")
}

// FullHD makes the class apply on full hd screens only.
func (c *ResponsiveClasses) FullHD() e.Classeser {
	return c.responsify("-fullhd")
}
