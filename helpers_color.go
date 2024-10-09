package docs

import (
	"fmt"

	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var (
	primaryColors = []struct {
		color b.Color
		name  string
	}{
		{b.Text, "Text"},
		{b.Link, "Link"},
		{b.Primary, "Primary"},
		{b.Info, "Info"},
		{b.Success, "Success"},
		{b.Warning, "Warning"},
		{b.Danger, "Danger"},
	}
	colorShades   = []int{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}
	colorPopupCSS = e.Style(gomponents.Raw(".colorsample>.popup{display:none}.colorsample:hover>.popup{display:block}"))
)

func colorSample(col, inv b.Color, name string) e.Element {
	box := b.Box(
		e.Class("colorsample"),
		b.Padding(0),
		b.Relative,
		e.Styles{
			"width":         "4em",
			"height":        "2em",
			"border-radius": "0.5em",
		},
		col,
	)

	popup := b.Box(
		e.Class("popup"),
		b.Padding(2),
		e.Styles{
			"position": "absolute",
			"top":      "2.5em",
			"left":     "0.5em",
			"z-index":  "50",
		},
	)

	if inv != nil {
		box.With(
			e.Div(
				e.Styles{
					"position":      "absolute",
					"top":           "0.25em",
					"right":         "0.25em",
					"width":         "1.5em",
					"height":        "1.5em",
					"border-radius": "0.25em",
				},
				inv,
			),
		)

		popup.With(
			e.Pre(
				b.Padding(2),
				b.Margin(0),
				e.Styles{"border-radius": "var(--bulma-radius-medium) var(--bulma-radius-medium) 0 0"},
				col.Text(),
				inv,
				"b.Text"+name+"\nb.Background"+name+"Invert",
			),
			e.Pre(
				b.Padding(2),
				b.Margin(0),
				e.Styles{"border-radius": "0 0 var(--bulma-radius-medium) var(--bulma-radius-medium)"},
				col,
				inv.Text(),
				"b.Text"+name+"Invert\nb.Background"+name,
			),
		)
	} else {
		popup.With(
			e.Pre(
				b.Padding(2),
				e.Styles{"border-radius": "var(--bulma-radius-medium) var(--bulma-radius-medium) 0 0"},
				col.Text(),
				b.BackgroundWhite,
				"b.Text"+name,
			),
			e.Pre(
				b.Padding(2),
				e.Styles{"border-radius": "0 0 var(--bulma-radius-medium) var(--bulma-radius-medium)"},
				col,
				b.TextWhite,
				"b.Background"+name,
			),
		)
	}
	return box.With(popup)
}

func cellSampleTitle(title string) e.Element {
	return b.TCell(
		b.TextCentered,
		e.Styles{"white-space": "nowrap"},
		e.Div(
			b.InlineBlock,
			e.Styles{
				"writing-mode": "vertical-lr",
				"transform":    "rotate(180deg)",
			},
			title,
		),
		e.Div(
			b.InlineBlock,
			e.Styles{
				"writing-mode": "vertical-lr",
				"transform":    "rotate(180deg)",
			},
			"...invert",
		),
	)
}

func rowSample(name string, variantMaker func(b.Color) b.Color, hasInvert bool) e.Element {
	row := b.Row()
	if name == "" {
		row.With(b.TCell(html.Th, "-"))
	} else {
		row.With(b.TCell(html.Th, e.Code(name)))
	}

	for _, base := range primaryColors {
		color := variantMaker(base.color.Background())
		if hasInvert {
			row.With(colorSample(color, color.Invert(), base.name+name))
		} else {
			row.With(colorSample(color, nil, base.name+name))
		}
	}
	return row
}

func paletteTitleRow(rowFn func(children ...any) e.Element) e.Element {
	row := rowFn("Variant")
	for _, color := range primaryColors {
		row.With(cellSampleTitle(color.name))
	}
	return row
}

func colorPalette() e.Element {
	table := b.Table(
		paletteTitleRow(b.HeadRow),
		paletteTitleRow(b.FootRow),
		rowSample("", func(c b.Color) b.Color { return c }, true),
	)
	for _, shade := range colorShades {
		table.With(rowSample(
			fmt.Sprintf("%02d", shade),
			func(c b.Color) b.Color {
				return c.Shade(shade)
			},
			true,
		))
	}
	table.With(
		rowSample("Dark", func(c b.Color) b.Color { return c.Dark() }, true),
		rowSample("Light", func(c b.Color) b.Color { return c.Light() }, true),
		rowSample("Soft", func(c b.Color) b.Color { return c.Soft() }, true),
		rowSample("Bold", func(c b.Color) b.Color { return c.Bold() }, true),
		rowSample("OnScheme", func(c b.Color) b.Color { return c.OnScheme() }, false),
	)
	return table
}

func myDumbComponent(color b.Color, content ...any) e.Element {
	return e.Div(
		color.Light().Background(),
		color.Light().Text().Invert(),
		e.P(content...),
	)
}

var color = c.NewPage(
	"Color", "Color helpers", "/color",
	"",

	colorPopupCSS,

	b.Content(
		e.P(e.Em("Bulma-Gomponents"), " provides predefined modifier variables for all colors available in ", e.Em("Bulma"), ". These modifiers names (corresponding to ", e.Em("Bulma"), " classes) are formatted as:", e.Br(), e.Code("b.[Target][Color][Variant][Invert]"), ", where:"),
		c.Table(
			b.HeadRow(b.FontSize(6), b.TCell(e.Code("Target"), "..."), "means..."),
			c.Row("Text", "Set text color"),
			c.Row("Background", "Set background color"),
			c.RowNoCode(
				"No target",
				"Set component color - it should work on ",
				e.Code("b.Button*"), ", ",
				e.Code("b.Icon"), ", ",
				e.Code("b.IconText"), ", ",
				e.Code("fa.FA"), ", ",
				e.Code("fa.Icon"), ", ",
				e.Code("fa.Li"), ", ",
				e.Code("b.Notification"), ", ",
				e.Code("b.Progress"), ", ",
				e.Code("b.ProgressIndeterminate"), ", ",
				e.Code("b.TCell"), ", ",
				e.Code("b.Tag"), ", ",
				e.Code("b.Message"), ", ",
				e.Code("b.Navbar"), ", ",
				e.Code("b.Panel"), ", ",
				e.Code("b.Help"), ", ",
				e.Code("b.Input*"), ", ",
				e.Code("b.Textarea"), ", ",
				e.Code("b.Select"), ", ",
				e.Code("b.File"), " and ",
				e.Code("b.Hero"),
				". Its use on icons and icon-related elements is specific to ", e.Em("Bulma-Gomponents"), " which transforms the color to the text-color equivalent for you.",
			),
		),
		c.Table(
			b.HeadRow(b.FontSize(6), b.TCell(e.Code("Color"), "..."), "means..."),
			c.RowNoCode(
				b.TCell(
					e.Code("White"), ", ",
					e.Code("Black"), ", ",
					e.Code("Dark"), ", ",
					e.Code("Light"), ", ",
					e.Code("Text"), ", ",
					e.Code("Link"), ", ",
					e.Code("Primary"), ", ",
					e.Code("Info"), ", ",
					e.Code("Success"), ", ",
					e.Code("Warning"), ", ",
					e.Code("Danger"), ", ",
					e.Code("BlackBis"), ", ",
					e.Code("BlackTer"), ", ",
					e.Code("GreyDarker"), ", ",
					e.Code("GreyDark"), ", ",
					e.Code("Grey"), ", ",
					e.Code("GreyLight"), ", ",
					e.Code("GreyLighter"), ", ",
					e.Code("WhiteTer"), " or ",
					e.Code("WhiteBis"), ".",
				),
				"Set the base color",
			),
		),
		c.Table(
			b.HeadRow(b.FontSize(6), b.TCell(e.Code("Variant"), "..."), "means..."),
			c.RowNoCode("No variant", "The base color"),
			c.RowMultiple(
				[]string{"00", "05", "10", "15", "20", "25", "30", "35", "40", "45", "50", "55", "60", "65", "70", "75", "80", "85", "90", "95", "100"},
				"The corresponding shade",
			),
			c.RowMultiple(
				[]string{
					"Light",
					"Dark",
					"Soft",
					"Bold",
					"OnScheme",
				},
				"The corresponding variant",
			),
		),
		c.Table(
			b.HeadRow(b.FontSize(6), b.TCell(e.Code("Invert"), "..."), "means..."),
			c.RowNoCode("No invert", "The base color"),
			c.Row("Invert", "Invert the color"),
		),
		e.P("Examples:"),
		c.Table(
			b.HeadRow("Example", "Color"),
			c.Row("b.Primary", "Primary color for a component or element which supports colors"),
			c.Row("b.TextSuccess60", "Success color with 60% opacity on the text"),
			c.Row("b.BackgroundDangerLight", "Danger color in its light variant on the background"),
			c.Row("b.BackgroundInfo, b.TextInfoInvert", "Info color on the background and a decent-looking color for the text"),
		),
	),
).Section(
	"Variant generation", "",

	b.Content(
		e.P(
			"In addition to the predefined modifier variables in the ", e.Code("b"), " package, the color modifiers provide methods to generate corresponding variants:",
		),
		c.Table(
			b.HeadRow("Method", "Action"),
			c.Row(".Base()", "Return the equivalent base color (for use within components or elements that support color)"),
			c.Row(".Text()", "Return the equivalent text color"),
			c.Row(".Background()", "Return the equivalent background color"),
			c.Row(".Normal()", "Remove any variant or shade on the color"),
			c.Row(".Light()", "Return the equivalent light variant"),
			c.Row(".Dark()", "Return the equivalent dark variant"),
			c.Row(".Soft()", "Return the equivalent soft variant"),
			c.Row(".Bold()", "Return the equivalent bold variant"),
			c.Row(".OnScheme()", "Return the equivalent on-scheme variant"),
			c.Row(".Shade(int)", "Return the equivalent shade variant (the integer value must be one of 0, 5, 10, ... 95, 100)"),
			c.Row(".Invert()", "Return the equivalent inverted color (for inverted colors, it returns the non-inverted equivalent; for base colors, it returns the same color because base colors have no inverted equivalent)"),
		),
		"These methods allow easily generate custom components with adequate colors, for example:",
	),
	c.Example(
		`func myDumbComponent(color b.Color, content ...any) e.Element {
	return e.Div(
		color.Light().Background(),
		color.Light().Text().Invert(),
		e.P(content...),
	),
}

func successAndDangerDumbComponent() []e.Element {
	return []e.Element{
		myDumbComponent(b.Success, "This is a success message"),
		myDumbComponent(b.Danger, "This is a danger message"),
	}
}`,
		myDumbComponent(b.Success, "This is a success message"),
		myDumbComponent(b.Danger, "This is a danger message"),
	),
).Section(
	"11 main colors", "https://bulma.io/documentation/features/color-palettes/",

	b.Content(
		`Bulma provides 4 base white/black colors and 7 primary colors, with "inverted" equivalents (which will look decent on the primary color):`,
	),
	b.Table(
		b.HeadRow("Color", "Sample"),
		b.Row(
			e.Code("White"),
			colorSample(b.BackgroundWhite, b.BackgroundWhiteInvert, "White"),
		),
		b.Row(
			e.Code("Black"),
			colorSample(b.BackgroundBlack, b.BackgroundBlackInvert, "Black"),
		),
		b.Row(
			e.Code("Light"),
			colorSample(b.BackgroundLight, b.BackgroundLightInvert, "Light"),
		),
		b.Row(
			e.Code("Dark"),
			colorSample(b.BackgroundDark, b.BackgroundDarkInvert, "Dark"),
		),
		b.Row(
			e.Code("Text"),
			colorSample(b.BackgroundText, b.BackgroundTextInvert, "Text"),
		),
		b.Row(
			e.Code("Link"),
			colorSample(b.BackgroundLink, b.BackgroundLinkInvert, "Link"),
		),
		b.Row(
			e.Code("Primary"),
			colorSample(b.BackgroundPrimary, b.BackgroundPrimaryInvert, "Primary"),
		),
		b.Row(
			e.Code("Info"),
			colorSample(b.BackgroundInfo, b.BackgroundInfoInvert, "Info"),
		),
		b.Row(
			e.Code("Success"),
			colorSample(b.BackgroundSuccess, b.BackgroundSuccessInvert, "Success"),
		),
		b.Row(
			e.Code("Warning"),
			colorSample(b.BackgroundWarning, b.BackgroundWarningInvert, "Warning"),
		),
		b.Row(
			e.Code("Danger"),
			colorSample(b.BackgroundDanger, b.BackgroundDangerInvert, "Danger"),
		),
	),
	b.Content(
		"The 7 primary colors also have the following variants, which you can check in the color palette below:",
		b.DList(
			c.Row("Light", "The primary color at 90% lightness"),

			c.Row("Dark", "The primary color at 20% lightness"),

			c.Row("Soft", "Light in light mode, dark in dark mode"),

			c.Row("Bold", "Dark in light mode, light in dark mode"),

			c.Row("OnScheme", "A variant that looks good on the scheme main color (white by default, which is used as the page background) - this variant does not have an inverted equivalent."),
		),
	),
).Section(
	"Shades of grey", "https://bulma.io/documentation/helpers/color-helpers/#text-color",

	b.Content(
		"The text and background colors may also be one of the following 9 shades of grey:",
	),
	b.Table(
		b.HeadRow("Color", "Sample"),
		b.Row(
			e.Code("BlackBis"),
			colorSample(b.BackgroundBlackBis, nil, "BlackBis"),
		),
		b.Row(
			e.Code("BlackTer"),
			colorSample(b.BackgroundBlackTer, nil, "BlackTer"),
		),
		b.Row(
			e.Code("GreyDarker"),
			colorSample(b.BackgroundGreyDarker, nil, "GreyDarker"),
		),
		b.Row(
			e.Code("GreyDark"),
			colorSample(b.BackgroundGreyDark, nil, "GreyDark"),
		),
		b.Row(
			e.Code("Grey"),
			colorSample(b.BackgroundGrey, nil, "Grey"),
		),
		b.Row(
			e.Code("GreyLight"),
			colorSample(b.BackgroundGreyLight, nil, "GreyLight"),
		),
		b.Row(
			e.Code("GreyLighter"),
			colorSample(b.BackgroundGreyLighter, nil, "GreyLighter"),
		),
		b.Row(
			e.Code("WhiteTer"),
			colorSample(b.BackgroundWhiteTer, nil, "WhiteTer"),
		),
		b.Row(
			e.Code("WhiteBis"),
			colorSample(b.BackgroundWhiteBis, nil, "WhiteBis"),
		),
	),
).Section(
	"Color palette", "https://bulma.io/documentation/helpers/palette-helpers/",

	b.Content(
		"For each of the 7 primary colors, when used on text or background, Bulma provides 21 shades with their inverted equivalents, 4 variants with their inverted equivalents and 1 variant without an inverted equivalent:",
	),
	colorPalette(),
)
