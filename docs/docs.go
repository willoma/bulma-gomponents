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
			documentRoot,
			elements,
			classes,
			extending,
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
		"Helpers and features",
		[]*c.Page{
			skeletons,
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
	content := el.Div(
		html.ID("bgd-content"),
		b.Style("margin-left", "11.25rem"),
		b.Padding(b.Spacing4),
		b.Title(
			el.A(html.Name("top")),
			p.Title,
		),
	)
	if p.BulmaURL != "" {
		content.With(b.Content(b.AHref(p.BulmaURL, html.Target("_blank"), "Bulma documentation")))
	}
	content.With(p.Children, p.InternalMenu())

	return b.HTML(
		b.Script(path.Join(p.BaseURL, "htmx.min.js")),
		b.HTitle(p.Title+" | Bulma-Gomponents"),
		b.CSSPath(path.Join(p.BaseURL, "bulma.css")),
		fa.CSSHead(path.Join(p.BaseURL, "fa")),
		gomponents.Attr("hx-on:htmx:load", `if(event.detail.elt.id ==='bgd-content')window.scrollTo(0,0)`),
		b.TopNavbar(
			b.Shadow,
			b.NavbarBrand(
				b.NavbarItem(
					b.Title(
						"Bulma-Gomponents",
					),
				),
			),
			b.NavbarEnd(
				b.NavbarAHref(
					"https://pkg.go.dev/github.com/willoma/bulma-gomponents",
					b.Tags(html.Span, b.InlineFlex, b.Addons, b.Tag(b.Dark, "Go"), b.Tag(b.Info, "Reference")),
				),
				b.NavbarAHref(
					"https://github.com/willoma/bulma-gomponents",
					b.Tags(html.Span, b.InlineFlex, b.Addons, b.Tag(b.Dark, "GitHub"), b.Tag(b.Success, "Repository")),
				),
				b.NavbarAHref(
					"https://bulma.io/documentation",
					b.Tags(html.Span, b.InlineFlex, b.Addons, b.Tag(b.Dark, "Bulma"), b.Tag(b.Warning, "Official documentation")),
				),
			),
		),
		b.Box(
			b.Style(
				"position", "fixed",
				"height", "100vh",
				"overflow-y", "auto",
				"width", "11.25rem",
				"top", "0",
				"left", "0",
				"padding-top", "3.25rem",
			),
			b.PaddingHorizontal(b.Spacing0),
			navMenu(p.Path),
		),
		content,
	)
}

func navMenu(currentPath string) b.Element {
	navmenu := b.Menu(
		html.ID("bgd-menu"),
		gomponents.Attr("hx-select-oob", "#bgd-menu,#bgd-content"),
		gomponents.Attr("hx-push-url", "true"),
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
