package docs

import (
	"fmt"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
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
	colorPopupCSS = el.Style(gomponents.Raw(".colorsample>.popup{display:none}.colorsample:hover>.popup{display:block}"))
)

func colorSample(col, inv b.Color, name string) b.Element {
	box := b.Box(
		b.Class("colorsample"),
		b.Padding(b.Spacing0),
		b.Relative,
		b.Style(
			"width", "4em",
			"height", "2em",
			"border-radius", "0.5em",
		),
		col,
	)

	popup := b.Box(
		b.Class("popup"),
		b.Padding(b.Spacing2),
		b.Style(
			"position", "absolute",
			"top", "2.5em",
			"left", "0.5em",
			"z-index", "50",
		),
	)

	if inv != nil {
		box.With(
			el.Div(
				b.Style(
					"position", "absolute",
					"top", "0.25em",
					"right", "0.25em",
					"width", "1.5em",
					"height", "1.5em",
					"border-radius", "0.25em",
				),
				inv,
			),
		)

		popup.With(
			el.Pre(
				b.Padding(b.Spacing2),
				b.Margin(b.Spacing0),
				b.Style("border-radius", "var(--bulma-radius-medium) var(--bulma-radius-medium) 0 0"),
				col.Text(),
				inv,
				"b.Text"+name+"\nb.Background"+name+"Invert",
			),
			el.Pre(
				b.Padding(b.Spacing2),
				b.Margin(b.Spacing0),
				b.Style("border-radius", "0 0 var(--bulma-radius-medium) var(--bulma-radius-medium)"),
				col,
				inv.Text(),
				"b.Text"+name+"Invert\nb.Background"+name,
			),
		)
	} else {
		popup.With(
			el.Pre(
				b.Padding(b.Spacing2),
				b.Style("border-radius", "var(--bulma-radius-medium) var(--bulma-radius-medium) 0 0"),
				col.Text(),
				b.BackgroundWhite,
				"b.Text"+name,
			),
			el.Pre(
				b.Padding(b.Spacing2),
				b.Style("border-radius", "0 0 var(--bulma-radius-medium) var(--bulma-radius-medium)"),
				col,
				b.TextWhite,
				"b.Background"+name,
			),
		)
	}
	return box.With(popup)
}

func cellSampleTitle(title string) b.Element {
	return b.TCell(
		b.TextCentered,
		b.Style(
			"white-space", "nowrap",
		),
		el.Div(
			b.InlineBlock,
			b.Style(
				"writing-mode", "vertical-lr",
				"transform", "rotate(180deg)",
			),
			title,
		),
		el.Div(
			b.InlineBlock,
			b.Style(
				"writing-mode", "vertical-lr",
				"transform", "rotate(180deg)",
			),
			"...invert",
		),
	)
}

func rowSample(name string, variantMaker func(b.Color) b.Color, hasInvert bool) b.Element {
	row := b.Row()
	if name == "" {
		row.With(b.TCell(html.Th, "-"))
	} else {
		row.With(b.TCell(html.Th, el.Code(name)))
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

func paletteTitleRow(rowFn func(children ...any) b.Element) b.Element {
	row := rowFn("Variant")
	for _, color := range primaryColors {
		row.With(cellSampleTitle(color.name))
	}
	return row
}

func colorPalette() b.Element {
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

func myDumbComponent(color b.Color, content ...any) b.Element {
	return el.Div(
		color.Light().Background(),
		color.Light().Text().Invert(),
		el.P(content...),
	)
}

var color = c.NewPage(
	"Color", "Color helpers", "/color",
	"",

	colorPopupCSS,

	b.Content(
		el.P(el.Em("Bulma-Gomponents"), " provides predefined modifier variables for all colors available in ", el.Em("Bulma"), ". These modifiers names (corresponding to ", el.Em("Bulma"), " classes) are formatted as:", el.Br(), el.Code("b.[Target][Color][Variant][Invert]"), ", where:"),
		el.Ul(
			el.Li(
				el.Code("Target"), " may be one of:",
				b.DList(
					el.Code("Text"),
					"Set the text color",

					el.Code("Background"),
					"Set the background color",

					"No target",
					[]any{
						"Set the component color - it should work on ",
						el.Code("b.Button*"), ", ",
						el.Code("b.Icon"), ", ",
						el.Code("b.IconText"), ", ",
						el.Code("fa.FA"), ", ",
						el.Code("fa.Icon"), ", ",
						el.Code("fa.Li"), ", ",
						el.Code("b.Notification"), ", ",
						el.Code("b.Progress"), ", ",
						el.Code("b.ProgressIndeterminate"), ", ",
						el.Code("b.TCell"), ", ",
						el.Code("b.Tag"), ", ",
						el.Code("b.Message"), ", ",
						el.Code("b.Navbar"), ", ",
						el.Code("b.Panel"), ", ",
						el.Code("b.Help"), ", ",
						el.Code("b.Input*"), ", ",
						el.Code("b.Textarea"), ", ",
						el.Code("b.Select"), ", ",
						el.Code("b.File"), " and ",
						el.Code("b.Hero"), ".",
						" Their use on icons and icon-related elements is specific to ", el.Em("Bulma-Gomponents"), " which transforms them to their text-color equivalent for you.",
					},
				),
			),
			el.Li(
				el.Code("Color"), " may be one of ",
				el.Code("White"), ", ",
				el.Code("Black"), ", ",
				el.Code("Dark"), ", ",
				el.Code("Light"), ", ",
				el.Code("Text"), ", ",
				el.Code("Link"), ", ",
				el.Code("Primary"), ", ",
				el.Code("Info"), ", ",
				el.Code("Success"), ", ",
				el.Code("Warning"), ", ",
				el.Code("Danger"), ", ",
				el.Code("BlackBis"), ", ",
				el.Code("BlackTer"), ", ",
				el.Code("GreyDarker"), ", ",
				el.Code("GreyDark"), ", ",
				el.Code("Grey"), ", ",
				el.Code("GreyLight"), ", ",
				el.Code("GreyLighter"), ", ",
				el.Code("WhiteTer"), " or ",
				el.Code("WhiteBis"), ".",
			),
			el.Li(
				el.Code("Variant"), " may be one of:",
				b.DList(
					"No variant",
					"The base color",

					[]any{
						el.Code("00"), ", ",
						el.Code("05"), ", ",
						el.Code("10"), ", ",
						el.Code("15"), ", ",
						el.Code("20"), ", ",
						el.Code("25"), ", ",
						el.Code("30"), ", ",
						el.Code("35"), ", ",
						el.Code("40"), ", ",
						el.Code("45"), ", ",
						el.Code("50"), ", ",
						el.Code("55"), ", ",
						el.Code("60"), ", ",
						el.Code("65"), ", ",
						el.Code("70"), ", ",
						el.Code("75"), ", ",
						el.Code("80"), ", ",
						el.Code("85"), ", ",
						el.Code("90"), ", ",
						el.Code("95"), " or ",
						el.Code("100"),
					},
					"The corresponding shade",

					[]any{
						el.Code("Light"), ", ",
						el.Code("Dark"), ", ",
						el.Code("Soft"), ", ",
						el.Code("Bold"), " or ",
						el.Code("OnScheme"),
					},
					"The corresponding variant",
				),
			),
			el.Li(
				el.Code("Invert"), " may be ignored for the base color, or provided for the inverted equivalent.",
			),
		),
		el.P("Examples:"),
		b.DList(
			el.Code("b.Primary"),
			"Primary color for a component or element which supports colors",

			el.Code("b.TextSuccess60"),
			"Success color with 60% opacity on the text",

			el.Code("b.BackgroundDangerLight"),
			"Danger color in its light variant on the background",

			el.Code("b.BackgroundInfo, b.TextInfoInvert"),
			"Info color on the background and a decent-looking color for the text",
		),
	),
).Section(
	"Variant generation", "",

	b.Content(
		el.P(
			"In addition to the predefined modifier variables in the ", el.Code("b"), " package, the color modifiers provide methods to generate corresponding variants:",
		),
		b.DList(
			el.Code(".Base()"),
			"Return the equivalent base color (for use within components or elements that support color)",

			el.Code(".Text()"),
			"Return the equivalent text color",

			el.Code(".Background()"),
			"Return the equivalent background color",

			el.Code(".Normal()"),
			"Remove any variant or shade on the color",

			el.Code(".Light()"),
			"Return the equivalent light variant",

			el.Code(".Dark()"),
			"Return the equivalent dark variant",

			el.Code(".Soft()"),
			"Return the equivalent soft variant",

			el.Code(".Bold()"),
			"Return the equivalent bold variant",

			el.Code(".OnScheme()"),
			"Return the equivalent on-scheme variant",

			el.Code(".Shade(int)"),
			"Return the equivalent shade variant (the integer value must be one of 0, 5, 10, ... 95, 100)",

			el.Code(".Invert()"),
			"Return the equivalent inverted color (for inverted colors, it returns the non-inverted equivalent; for base colors, it returns the same color because base colors have no inverted equivalent)",
		),
		"These methods allow easily generate custom components with adequate colors, for example:",
	),
	c.Example(
		`func myDumbComponent(color b.Color, content ...any) b.Element {
	return el.Div(
		color.Light().Background(),
		color.Light().Text().Invert(),
		el.P(content...),
	),
}

func successAndDangerDumbComponent() []b.Element {
	return []b.Element{
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
			el.Code("White"),
			colorSample(b.BackgroundWhite, b.BackgroundWhiteInvert, "White"),
		),
		b.Row(
			el.Code("Black"),
			colorSample(b.BackgroundBlack, b.BackgroundBlackInvert, "Black"),
		),
		b.Row(
			el.Code("Light"),
			colorSample(b.BackgroundLight, b.BackgroundLightInvert, "Light"),
		),
		b.Row(
			el.Code("Dark"),
			colorSample(b.BackgroundDark, b.BackgroundDarkInvert, "Dark"),
		),
		b.Row(
			el.Code("Text"),
			colorSample(b.BackgroundText, b.BackgroundTextInvert, "Text"),
		),
		b.Row(
			el.Code("Link"),
			colorSample(b.BackgroundLink, b.BackgroundLinkInvert, "Link"),
		),
		b.Row(
			el.Code("Primary"),
			colorSample(b.BackgroundPrimary, b.BackgroundPrimaryInvert, "Primary"),
		),
		b.Row(
			el.Code("Info"),
			colorSample(b.BackgroundInfo, b.BackgroundInfoInvert, "Info"),
		),
		b.Row(
			el.Code("Success"),
			colorSample(b.BackgroundSuccess, b.BackgroundSuccessInvert, "Success"),
		),
		b.Row(
			el.Code("Warning"),
			colorSample(b.BackgroundWarning, b.BackgroundWarningInvert, "Warning"),
		),
		b.Row(
			el.Code("Danger"),
			colorSample(b.BackgroundDanger, b.BackgroundDangerInvert, "Danger"),
		),
	),
	b.Content(
		"The 7 primary colors also have the following variants, which you can check in the color palette below:",
		b.DList(
			el.Code("Light"),
			"The primary color at 90% lightness",

			el.Code("Dark"),
			"The primary color at 20% lightness",

			el.Code("Soft"),
			"Light in light mode, dark in dark mode",

			el.Code("Bold"),
			"Dark in light mode, light in dark mode",

			el.Code("OnScheme"),
			"A variant that looks good on the scheme main color (white by default, which is used as the page background) - this variant does not have an inverted equivalent.",
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
			el.Code("BlackBis"),
			colorSample(b.BackgroundBlackBis, nil, "BlackBis"),
		),
		b.Row(
			el.Code("BlackTer"),
			colorSample(b.BackgroundBlackTer, nil, "BlackTer"),
		),
		b.Row(
			el.Code("GreyDarker"),
			colorSample(b.BackgroundGreyDarker, nil, "GreyDarker"),
		),
		b.Row(
			el.Code("GreyDark"),
			colorSample(b.BackgroundGreyDark, nil, "GreyDark"),
		),
		b.Row(
			el.Code("Grey"),
			colorSample(b.BackgroundGrey, nil, "Grey"),
		),
		b.Row(
			el.Code("GreyLight"),
			colorSample(b.BackgroundGreyLight, nil, "GreyLight"),
		),
		b.Row(
			el.Code("GreyLighter"),
			colorSample(b.BackgroundGreyLighter, nil, "GreyLighter"),
		),
		b.Row(
			el.Code("WhiteTer"),
			colorSample(b.BackgroundWhiteTer, nil, "WhiteTer"),
		),
		b.Row(
			el.Code("WhiteBis"),
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
