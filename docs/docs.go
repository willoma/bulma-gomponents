package docs

import (
	_ "embed"
	"path"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

//go:embed htmx.min.js
var HtmxJS []byte

type docSection struct {
	title string
	Pages []*c.Page
}

var Sections = []docSection{
	{
		"",
		[]*c.Page{
			intro,
		},
	},
	{
		"Columns",
		[]*c.Page{
			columnsBasics,
			columnsSizes,
			columnsResponsiveness,
			columnsNesting,
			columnsGap,
			columnsOptions,
		},
	},
	{
		"Elements",
		[]*c.Page{
			block,
			box,
			button,
			content,
			delete,
			icon,
			image,
			notification,
			progress,
			table,
			tag,
			title,
		},
	},
	{
		"Components",
		[]*c.Page{
			breadcrumb,
			card,
			dropdown,
			menu,
			message,
			modal,
			navbar,
			pagination,
			panel,
			tabs,
		},
	},
	{
		"Form",
		[]*c.Page{
			formGeneral,
			formInput,
			formTextarea,
			formSelect,
			formCheckbox,
			formRadio,
			formFile,
		},
	},
	{
		"Layout",
		[]*c.Page{
			container,
			level,
			mediaObject,
			hero,
			section,
			footer,
			tiles,
		},
	},
	{
		"Helpers",
		[]*c.Page{
			color,
			spacing,
			typography,
			visibility,
			flexbox,
			other,
		},
	},
}

func Layout(p *c.Page) gomponents.Node {
	return b.HTML(
		b.Script(path.Join(p.BaseURL, "htmx.min.js")),
		b.HTitle("Bulma-Gomponents - "+p.Title),
		b.CSSPath(path.Join(p.BaseURL, "bulma.css")),
		fa.CSSPath(path.Join(p.BaseURL, "fa")),
		b.Box(
			b.Style(
				"position", "fixed",
				"height", "100vh",
				"overflow-y", "auto",
				"width", "11.25rem",
				"top", "0",
				"left", "0",
			),
			b.PaddingHorizontal(b.Spacing0),
			navMenu(p.Path),
		),
		el.Div(
			html.ID("content"),
			b.Style("margin-left", "11.25rem"),
			b.Padding(b.Spacing4),
			b.Title(
				el.A(html.Name("top")),
				p.Title,
			),
			b.Content(b.AHref(p.BulmaURL, html.Target("_blank"), "Bulma documentation")),
		).Withs(p.Children).With(p.InternalMenu()),
	)
}

func navMenu(currentPath string) *b.Element {
	navmenu := b.Menu(
		html.ID("menu"),
		b.MenuLabel(
			html.H1,
			b.MarginHorizontal(b.Spacing2),
			el.Strong("Bulma-Gomponents"),
		),
	)

	for _, section := range Sections {
		if section.title != "" {
			navmenu.With(b.MenuLabel(
				b.PaddingLeft(b.Spacing2),
				section.title,
			))
		}
		thislist := b.MenuList()
		for _, page := range section.Pages {
			thislist.With(
				page.MenuEntry(currentPath),
			)
		}
		navmenu.With(thislist)
	}

	return navmenu
}
