package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var documentRoot = c.NewPage(
	"Document root", "HTML document root", "/root",
	"",

	b.Content(
		el.P("In order to generate a document root, use ", el.Code("b.HTML(...any)"), ". The allowed options are:"),
		b.DList(
			el.Code("b.CSSPath(string)"),
			[]any{"path to the CSS file, automatically applied in a ", el.Code(`<link rel="stylesheet">`), " node (if it is not defined, the default CDN is used)"},
			el.Code("b.Language(string)"),
			"define the document language",
			el.Code("b.HTitle(string)"),
			"define the document title",
			el.Code("b.Description(string)"),
			"define the document description",
			el.Code("b.Head(...gomponents.Node)"),
			[]any{"add one or multiple children to the ", el.Code("<head>")},
			"any other value",
			[]any{"add the node to the ", el.Code("<body>")},
		),
		el.H2("Example"),
		el.Pre(`b.HTML(
	b.Script("/htmx.min.js"),
	b.HTitle(p.Title+" | Bulma-Gomponents"),
	b.CSSPath("/bulma.css"),
	fa.CSSHead("/fa"),
	b.Language("en"),
	b.Head(html.Meta(html.Charset("utf-8"))),
	b.Content(
		el.H1("Hello"),
		el.P("Hello world"),
	),
)`),
	),
)
