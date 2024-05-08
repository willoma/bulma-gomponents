package bulma

import e "github.com/willoma/gomplements"

type baseColor struct {
	name    string
	variant string
}

func (c baseColor) Classes() []e.Class {
	if c.name == "" {
		return []e.Class{}
	}

	classes := []e.Class{e.Class("is-" + c.name)}

	if c.variant != "" {
		classes = append(classes, e.Class("is-"+c.variant))
	}

	return classes
}

func (c baseColor) If(cond bool) Color {
	if cond {
		return c
	}

	return baseColor{}
}

func (c baseColor) Base() Color {
	return c
}

func (c baseColor) Text() Color {
	return partColor{
		name:    c.name,
		text:    true,
		variant: c.variant,
	}
}

func (c baseColor) Background() Color {
	return partColor{
		name:    c.name,
		text:    false,
		variant: c.variant,
	}
}

func (c baseColor) Normal() Color {
	return baseColor{
		name: c.name,
	}
}

func (c baseColor) Light() Color {
	return baseColor{
		name:    c.name,
		variant: colorLight,
	}
}

func (c baseColor) Dark() Color {
	return baseColor{
		name:    c.name,
		variant: colorDark,
	}
}

func (c baseColor) Soft() Color {
	return baseColor{
		name:    c.name,
		variant: colorSoft,
	}
}

func (c baseColor) Bold() Color {
	return baseColor{
		name:    c.name,
		variant: colorBold,
	}
}

func (c baseColor) OnScheme() Color {
	return baseColor{
		name:    c.name,
		variant: colorOnScheme,
	}
}

func (c baseColor) Shade(int) Color {
	return c
}

func (c baseColor) Invert() Color {
	return c
}

var (
	White Color = baseColor{name: "white"}
	Black Color = baseColor{name: "black"}
	Light Color = baseColor{name: "light"}
	Dark  Color = baseColor{name: "dark"}

	Text Color = baseColor{name: "text"}

	Primary      Color = baseColor{name: "primary"}
	PrimaryLight Color = baseColor{name: "primary", variant: colorLight}
	PrimaryDark  Color = baseColor{name: "primary", variant: colorDark}
	PrimarySoft  Color = baseColor{name: "primary", variant: colorSoft}
	PrimaryBold  Color = baseColor{name: "primary", variant: colorBold}

	Link      Color = baseColor{name: "link"}
	LinkLight Color = baseColor{name: "link", variant: colorLight}
	LinkDark  Color = baseColor{name: "link", variant: colorDark}
	LinkSoft  Color = baseColor{name: "link", variant: colorSoft}
	LinkBold  Color = baseColor{name: "link", variant: colorBold}

	Info      Color = baseColor{name: "info"}
	InfoLight Color = baseColor{name: "info", variant: colorLight}
	InfoDark  Color = baseColor{name: "info", variant: colorDark}
	InfoSoft  Color = baseColor{name: "info", variant: colorSoft}
	InfoBold  Color = baseColor{name: "info", variant: colorBold}

	Success      Color = baseColor{name: "success"}
	SuccessLight Color = baseColor{name: "success", variant: colorLight}
	SuccessDark  Color = baseColor{name: "success", variant: colorDark}
	SuccessSoft  Color = baseColor{name: "success", variant: colorSoft}
	SuccessBold  Color = baseColor{name: "success", variant: colorBold}

	Warning      Color = baseColor{name: "warning"}
	WarningLight Color = baseColor{name: "warning", variant: colorLight}
	WarningDark  Color = baseColor{name: "warning", variant: colorDark}
	WarningSoft  Color = baseColor{name: "warning", variant: colorSoft}
	WarningBold  Color = baseColor{name: "warning", variant: colorBold}

	Danger      Color = baseColor{name: "danger"}
	DangerLight Color = baseColor{name: "danger", variant: colorLight}
	DangerDark  Color = baseColor{name: "danger", variant: colorDark}
	DangerSoft  Color = baseColor{name: "danger", variant: colorSoft}
	DangerBold  Color = baseColor{name: "danger", variant: colorBold}

	// These do not exist as colors for elements, but they may be used for icons with the automatic color type changing, with b.Icon*.
	BlackBis    Color = baseColor{name: "black-bis"}
	BlackTer    Color = baseColor{name: "black-ter"}
	GreyDarker  Color = baseColor{name: "grey-darker"}
	GreyDark    Color = baseColor{name: "grey-dark"}
	Grey        Color = baseColor{name: "grey"}
	GreyLight   Color = baseColor{name: "grey-light"}
	GreyLighter Color = baseColor{name: "grey-lighter"}
	WhiteTer    Color = baseColor{name: "white-ter"}
	WhiteBis    Color = baseColor{name: "white-bis"}
)
