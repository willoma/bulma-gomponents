package components

import (
	"strings"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"

	"github.com/maragudk/gomponents"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
)

const warningDark = "hsl(48, 100%, 29%)"

var (
	colorFormatter = html.New(
		html.PreventSurroundingPre(true),
	)
	colorStyle = styles.Get("catppuccin-latte")

	colorGoLexer   = lexers.Get("go")
	colorHTMLLexer = lexers.Get("html")
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
		b.Style("border-left", "1em solid "+warningDark),
		b.Style("border-right", "0.25em solid "+warningDark),
		b.Style("border-radius", "1.2em 0.5em 0.5em 1.2em"),
		b.Style("position", "relative"),
		b.Style("padding-left", "0.2em"),
		b.Style("padding-right", "0.2em"),
		b.PaddingHorizontal(b.Spacing2),
		el.Div(
			b.FontSize7,
			b.WeightBold,
			b.TextWhite,
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
				ExamplePre(code),
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
			ExamplePre(code),
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

func ExamplePre(code string) b.Element {
	return pre(colorGo(code))
}

func pre(children ...any) b.Element {
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
		if elem, ok := e.(gomponents.Node); ok {
			elem.Render(&result)
		}
		result.WriteByte('\n')
	}

	return b.Box(
		b.Hidden,
		b.Class("html"),
		b.Padding(b.Spacing2),
		b.Style(
			"position", "absolute",
			"width", "100%",
			"top", "0",
			"left", "0",
			"z-index", "50",
		),
		pre(
			b.Style(
				"white-space", "pre-wrap",
				"border-radius", "var(--bulma-radius-small)",
			),
			colorHTML(result.String()),
		),
	)
}

func colorGo(source string) any {
	iterator, err := colorGoLexer.Tokenise(nil, source)
	if err != nil {
		return source
	}

	var buf strings.Builder
	if err := colorFormatter.Format(&buf, colorStyle, iterator); err != nil {
		return source
	}

	return gomponents.Raw(buf.String())
}

func colorHTML(source string) any {
	iterator, err := colorHTMLLexer.Tokenise(nil, source)
	if err != nil {
		return source
	}

	var buf strings.Builder
	if err := colorFormatter.Format(&buf, colorStyle, iterator); err != nil {
		return source
	}

	return gomponents.Raw(buf.String())
}
