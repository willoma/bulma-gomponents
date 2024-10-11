package components

import (
	"strings"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents"

	b "github.com/willoma/bulma-gomponents"
)

// Color code from https://bulma.io/documentation/helpers/color-helpers/
const colorInfo = "hsl(217, 71%, 45%)"

var (
	colorFormatter = html.New(
		html.PreventSurroundingPre(true),
	)
	colorStyle = styles.Get("catppuccin-latte")

	colorGoLexer   = lexers.Get("go")
	colorHTMLLexer = lexers.Get("html")
)

var CellStyle = []any{
	b.PaddingVertical(3), b.PaddingHorizontal(4),
	b.BackgroundPrimary, b.TextPrimaryInvert,
	b.RadiusNormal,
}

func ColParagraph(children ...any) e.Element {
	return e.P(
		CellStyle,
		b.TextCentered,
		children,
	)
}

func exampleContainer(children ...any) e.Element {
	return b.Block(
		e.Styles{
			"border-left":   "1em solid " + colorInfo,
			"border-right":  "0.25em solid " + colorInfo,
			"border-radius": "1.2em 0.5em 0.5em 1.2em",
			"position":      "relative",
			"padding-left":  "0.2em",
			"padding-right": "0.2em",
		},
		b.PaddingHorizontal(2),
		e.Div(
			b.FontSize(7),
			b.WeightBold,
			b.TextWhite,
			e.Styles{
				"position":     "absolute",
				"top":          "50%",
				"left":         "-1.5em",
				"writing-mode": "vertical-lr",
				"transform":    "translateY(-50%) rotate(180deg)",
			},
			"Example",
		),
		children,
	)
}

func Example(code string, result ...any) e.Element {
	return exampleContainer(
		b.Columns(
			b.ColumnGap(1),
			b.Column(
				b.Clipped,
				b.FlexGrow(1), b.FlexShrink(1),
				b.PaddingVertical(0),
				codeTag(),
				ExamplePre(code),
			),
			b.Column(
				e.Class("exampleresult"),
				b.FlexGrow(1), b.FlexShrink(1),
				b.PaddingVertical(0),
				resultTags(),
				e.Div(
					e.Styles{"position": "relative"},
					e.Div(e.Class("result"), result),
					htmlPre(result),
				),
			),
		),
	)
}

func HorizontalExample(code string, result ...any) e.Element {
	return exampleContainer(
		e.Div(
			e.Class("exampleresult"),
			b.MarginBottom(1),
			resultTags(),
			e.Div(
				e.Styles{"position": "relative"},
				e.Div(e.Class("result"), result),
				htmlPre(result),
			),
		),
		e.Div(
			b.MarginTop(2),
			b.Clipped,
			codeTag(),
			ExamplePre(code),
		),
	)
}

func codeTag() e.Element {
	return b.Tags(
		b.MarginBottom(0),
		b.Tag(
			b.GreyLight, "Code", b.WeightBold, b.MarginBottom(0),
			e.Styles{
				"border-bottom-left-radius":  "0",
				"border-bottom-right-radius": "0",
			},
		),
	)
}

func resultTags() e.Element {
	return b.Tags(
		b.MarginBottom(1),
		b.Tag(
			b.Info, "Result",
			b.MarginBottom(0),
			b.Clickable,
			e.OnClick(`
				this.closest(".exampleresult").getElementsByClassName("html")[0].classList.toggle("is-hidden")
				this.closest(".exampleresult").getElementsByClassName("result")[0].classList.toggle("is-invisible")
				this.classList.toggle("is-light")
				this.nextSibling.classList.toggle("is-light")
			`),
		),
		b.Tag(
			b.WarningLight, "HTML",
			b.MarginBottom(0),
			b.Clickable,
			e.OnClick(`
				this.closest(".exampleresult").getElementsByClassName("html")[0].classList.toggle("is-hidden")
				this.closest(".exampleresult").getElementsByClassName("result")[0].classList.toggle("is-invisible")
				this.previousSibling.classList.toggle("is-light")
				this.classList.toggle("is-light")
			`),
		),
	)
}

func ExamplePre(code string) e.Element {
	return pre(colorGo(code))
}

func pre(children ...any) e.Element {
	return e.Pre(
		b.Padding(2),
		b.FontSize(7),
		e.Styles{"tab-size": "4"},
		children,
	)
}

func htmlPre(elements []any) e.Element {
	var result strings.Builder
	for _, e := range elements {
		if elem, ok := e.(gomponents.Node); ok {
			elem.Render(&result)
		}
		result.WriteByte('\n')
	}

	return b.Box(
		b.Hidden,
		e.Class("html"),
		b.Padding(2),
		b.BackgroundWarning,
		e.Styles{
			"position": "absolute",
			"width":    "100%",
			"top":      "0",
			"left":     "0",
			"z-index":  "50",
		},
		pre(
			e.Styles{
				"white-space":   "pre-wrap",
				"border-radius": "var(--bulma-radius-small)",
			},
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
