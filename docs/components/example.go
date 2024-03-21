package components

import (
	"strings"

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
				codeTag(),
				examplePre(code),
			),
			b.Column(
				b.Class("exampleresult"),
				b.FlexGrow1, b.FlexShrink1,
				b.PaddingVertical(b.Spacing0),
				resultTags(),
				el.Div(
					b.Style("position", "relative"),
					el.Div(result...),
					htmlPre(result),
				),
			),
		),
	)
}

func HorizontalExample(code string, result ...any) b.Element {
	return exampleContainer(
		el.Div(
			b.Class("exampleresult"),
			b.MarginBottom(b.Spacing1),
			resultTags(),
			el.Div(
				b.Style("position", "relative"),
				el.Div(result...),
				htmlPre(result),
			),
		),
		el.Div(
			b.Clipped,
			codeTag(),
			examplePre(code),
		),
	)
}

func codeTag() b.Element {
	return b.Tags(
		b.MarginBottom(b.Spacing1),
		b.Tag(b.Primary, "Code", b.MarginBottom(b.Spacing0)),
	)
}

func resultTags() b.Element {
	return b.Tags(
		b.MarginBottom(b.Spacing1),
		b.Tag(b.Info, "Result", b.MarginBottom(b.Spacing0)),
		b.Tag(
			b.WarningLight, "HTML",
			b.MarginBottom(b.Spacing0),
			b.Style("cursor", "pointer"),
			b.OnClick(`
			this.closest(".exampleresult").getElementsByClassName("html")[0].classList.toggle("is-hidden")
			this.classList.toggle("is-light")
			`),
		),
	)
}

func examplePre(children ...any) b.Element {
	return el.Pre(
		b.Padding(b.Spacing2),
		b.FontSize7,
		b.Style("tab-size", "4"),
		children,
	)
}

func htmlPre(elements []any) b.Element {
	var result strings.Builder
	for _, e := range elements {
		if elem, ok := e.(b.Element); ok {
			elem.Render(&result)
		}
		result.WriteByte('\n')
	}

	return examplePre(
		b.Class("html"),
		b.Hidden,
		b.Style(
			"position", "absolute",
			"width", "100%",
			"top", "0",
			"left", "0",
			"white-space", "pre-wrap",
			"z-index", "10",
		),
		result.String(),
	)
}
