package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

func colorSample(class b.Class) *b.Element {
	return b.Box(
		b.Style("display", "inline-block"),
		b.Style("vertical-align", "bottom"),
		b.PaddingHorizontal(b.Spacing5),
		b.PaddingVertical(b.Spacing3),
		class,
	)
}

var color = c.NewPage(
	"Color", "Color helpers", "/color",
	"https://bulma.io/documentation/helpers/color-helpers/",
	b.Content(el.P("The following modifiers may be used to define an element or component ", el.Strong("modifier"), ":")),
	b.Columns(
		b.Multiline,
		b.Table(
			b.HeadRow("Color modifier", "Example"),
			b.Row(el.Code("b.White"), colorSample(b.White.Background())),
			b.Row(el.Code("b.Black"), colorSample(b.Black.Background())),
			b.Row(el.Code("b.Light"), colorSample(b.Light.Background())),
			b.Row(el.Code("b.Dark"), colorSample(b.Dark.Background())),
			b.Row(el.Code("b.Primary"), colorSample(b.Primary.Background())),
			b.Row(el.Code("b.Link"), colorSample(b.Link.Background())),
			b.Row(el.Code("b.Info"), colorSample(b.Info.Background())),
			b.Row(el.Code("b.Success"), colorSample(b.Success.Background())),
			b.Row(el.Code("b.Warning"), colorSample(b.Warning.Background())),
			b.Row(el.Code("b.Danger"), colorSample(b.Danger.Background())),
		),
		b.Table(
			b.HeadRow("Shade modifier", "Example"),
			b.Row(el.Code("b.BlackBis"), colorSample(b.BlackBis.Background())),
			b.Row(el.Code("b.BlackTer"), colorSample(b.BlackTer.Background())),
			b.Row(el.Code("b.GreyDarker"), colorSample(b.GreyDarker.Background())),
			b.Row(el.Code("b.GreyDark"), colorSample(b.GreyDark.Background())),
			b.Row(el.Code("b.Grey"), colorSample(b.Grey.Background())),
			b.Row(el.Code("b.GreyLight"), colorSample(b.GreyLight.Background())),
			b.Row(el.Code("b.GreyLighter"), colorSample(b.GreyLighter.Background())),
			b.Row(el.Code("b.WhiteTer"), colorSample(b.WhiteTer.Background())),
			b.Row(el.Code("b.WhiteBis"), colorSample(b.WhiteBis.Background())),
		),
		b.Table(
			b.HeadRow("Light modifier", "Example"),
			b.Row(el.Code("b.PrimaryLight"), colorSample(b.Primary.BackgroundLight())),
			b.Row(el.Code("b.LinkLight"), colorSample(b.Link.BackgroundLight())),
			b.Row(el.Code("b.InfoLight"), colorSample(b.Info.BackgroundLight())),
			b.Row(el.Code("b.SuccessLight"), colorSample(b.Success.BackgroundLight())),
			b.Row(el.Code("b.WarningLight"), colorSample(b.Warning.BackgroundLight())),
			b.Row(el.Code("b.DangerLight"), colorSample(b.Danger.BackgroundLight())),
		),
		b.Table(
			b.HeadRow("Dark modifier", "Example"),
			b.Row(el.Code("b.Primary.XXXDark()"), colorSample(b.Primary.BackgroundDark())),
			b.Row(el.Code("b.Link.XXXDark()"), colorSample(b.Link.BackgroundDark())),
			b.Row(el.Code("b.Info.XXXDark()"), colorSample(b.Info.BackgroundDark())),
			b.Row(el.Code("b.Success.XXXDark()"), colorSample(b.Success.BackgroundDark())),
			b.Row(el.Code("b.Warning.XXXDark()"), colorSample(b.Warning.BackgroundDark())),
			b.Row(el.Code("b.Danger.XXXDark()"), colorSample(b.Danger.BackgroundDark())),
		),
	),
).Section(
	"Text color",
	"https://bulma.io/documentation/helpers/color-helpers/#text-color",
	b.Content(
		"You can change the text color by calling one of the following functions on a color value:",
		el.Ul(
			el.Li(el.Code(".Text()"), " to set the text to the named color"),
			el.Li(el.Code(".TextLight()"), " to set the text to the light color"),
			el.Li(el.Code(".TextDark()"), " to set the text to the dark color"),
		),
	),
	c.Example(
		`b.Content(
	el.P(b.Info.Text(), "Info text"),
	el.P(b.Danger.TextDark(), "Danger dark text"),
)`,
		b.Content(
			el.P(b.Info.Text(), "Info text"),
			el.P(b.Danger.TextDark(), "Danger dark text"),
		),
	),
).Section(
	"Background color",
	"https://bulma.io/documentation/helpers/color-helpers/#background-color",
	b.Content(
		"You can change the background color by calling one of the following functions on a color value:",
		el.Ul(
			el.Li(el.Code(".Background()"), " to set the background to the named color"),
			el.Li(el.Code(".BackgroundLight()"), " to set the background to the light color"),
			el.Li(el.Code(".BackgroundDark()"), " to set the background to the dark color"),
		),
	),
	c.Example(
		`b.Content(
	el.P(b.Warning.Background(), "Warning background"),
	el.P(b.Danger.BackgroundLight(), "Danger light background"),
)`,
		b.Content(
			el.P(b.Warning.Background(), "Warning background"),
			el.P(b.Danger.BackgroundLight(), "Danger light background"),
		),
	),
)
