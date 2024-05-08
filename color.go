package bulma

import e "github.com/willoma/gomplements"

//go:generate go run cmd/colors_generator/*.go

type Color interface {
	// Base returns the base color equivalent.
	Base() Color

	// Text returns the text color equivalent.
	Text() Color

	// Background returns the background color equivalent.
	Background() Color

	// Normal returns the normal color equivalent (no variant).
	Normal() Color

	// Light returns the light color equivalent.
	Light() Color

	// Dark returns the dark color equivalent.
	Dark() Color

	// Soft returns the soft color equivalent.
	Soft() Color

	// Bold returns the bold color equivalent.
	Bold() Color

	// OnScheme returns the on-scheme color equivalent.
	OnScheme() Color

	// Shade returns the shaded color equivalent (it only works with shades
	// defined by Bulma: 0, 5, 10, 15, ... 100).
	//
	// For Base colors, it returns itself because base colors have no shade equivalent.
	Shade(shade int) Color

	// Invert returns the inverted color equivalent.
	//
	//   - For Invert colors, it returns the non-inverted color equivalent.
	//   - For Base colors, it returns itself because base colors have no inverted equivalent.
	Invert() Color

	// If applies the color conditionally
	If(cond bool) e.ParentModifier
}

const (
	colorLight    = "light"
	colorDark     = "dark"
	colorSoft     = "soft"
	colorBold     = "bold"
	colorOnScheme = "on-scheme"
)
