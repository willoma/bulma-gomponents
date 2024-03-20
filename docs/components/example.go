package components

import (
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

func ColParagraph(children ...any) b.Element {
	return el.P(
		b.Style("outline", "1px dashed #36b6e0"),
		b.Style("border-radius", "0.25em"),
		b.TextCentered,
		children,
	)
}

func exampleContainer(children ...any) b.Element {
	return b.Block(
		b.Style("border-left", "1em solid hsl(48, 100%, 29%)"), // warning-dark
		b.Style("border-right", "0.25em solid hsl(48, 100%, 29%)"),
		b.Style("border-radius", "1.2em 0.5em 0.5em 1.2em"),
		b.Style("position", "relative"),
		b.Style("padding-left", "0.2em"),
		b.Style("padding-right", "0.2em"),
		b.PaddingHorizontal(b.Spacing2),
		el.Div(
			b.FontSize7,
			b.Bold,
			b.White.Text(),
			b.Style("position", "absolute"),
			b.Style("top", "50%"),
			b.Style("left", "-1.5em"),
			b.Style("writing-mode", "vertical-lr"),
			b.Style("transform", "translateY(-50%) rotate(180deg)"),
			"Example",
		),
		children,
	)
}

func Example(code string, result ...any) b.Element {
	return exampleContainer(
		b.Columns(
			b.Gap1,
			b.Column(
				b.Clipped,
				b.FlexGrow1, b.FlexShrink1,
				b.PaddingVertical(b.Spacing0),
				b.Tag(b.Primary, "Code", b.MarginBottom(b.Spacing1)),
				el.Pre(
					b.Padding(b.Spacing2),
					b.FontSize7,
					b.Style("tab-size", "4"),
					code,
				),
			),
			b.Column(
				b.FlexGrow1, b.FlexShrink1,
				b.PaddingVertical(b.Spacing0),
				b.Tag(b.Info, "Result", b.MarginBottom(b.Spacing1)),
				el.Div(result...),
			),
		),
	)
}

func HorizontalExample(code string, result ...any) b.Element {
	return exampleContainer(
		el.Div(
			b.Tag(b.Info, "Result", b.MarginBottom(b.Spacing1)),
			el.Div(result...),
			b.MarginBottom(b.Spacing1),
		),
		el.Div(
			b.Clipped,
			b.Tag(b.Primary, "Code", b.MarginBottom(b.Spacing1)),
			el.Pre(
				b.Padding(b.Spacing2),
				b.FontSize7,
				b.Style("tab-size", "4"),
				code,
			),
		),
	)
}
