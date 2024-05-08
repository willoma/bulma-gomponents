package components

import (
	"path"
	"strings"
	"sync"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

type Page struct {
	menuentry          string
	Title              string
	Path               string
	BaseURL            string
	BulmaURL           string
	Children           []any
	internalMenu       []any
	internalMenuScript string

	prepared    gomponents.Node
	prepareOnce sync.Once
}

type DocSection struct {
	Title string
	Pages []*Page
}

func (p *Page) URL() string {
	return path.Join(p.BaseURL, p.Path)
}

func NewPage(menuentry, title, path, bulmaURL string, children ...any) *Page {
	return &Page{
		menuentry:          menuentry,
		Title:              title,
		Path:               path,
		BaseURL:            "/",
		BulmaURL:           bulmaURL,
		Children:           children,
		internalMenuScript: "acnc='has-background-link-light'", // ACtive Navigation Class
	}
}

func (p *Page) Prepare(sections []DocSection) gomponents.Node {
	p.prepareOnce.Do(func() {
		content := e.Div(
			e.ID("bgd-content"),
			e.Styles{"margin-left": "11.25rem"},
			b.Padding(b.Spacing4),
			b.Title(
				e.A(e.Name("top")),
				p.Title,
			),
		)
		if p.BulmaURL != "" {
			content.With(b.Content(e.AHref(p.BulmaURL, e.Target("_blank"), "Bulma documentation")))
		}
		content.With(p.Children, p.InternalMenu())

		navbar := b.Navbar(
			e.Styles{"z-index": "100"},
			b.Shadow,
			b.FixedTop,
			b.NavbarBrand(
				b.NavbarItem(
					b.Title(
						"Bulma-Gomponents",
					),
				),
			),
			b.NavbarEnd(
				b.NavbarItem(
					b.Buttons(
						b.Small,
						b.ButtonAHref(
							"https://pkg.go.dev/github.com/willoma/bulma-gomponents",
							b.Primary,
							fa.Icon(fa.Brand, "golang"),
							"Reference",
						),
						b.ButtonAHref(
							"https://github.com/willoma/bulma-gomponents",
							b.Info,
							fa.Icon(fa.Brand, "github"),
							"Repository",
						),
						b.ButtonAHref(
							"https://bulma.io/documentation",
							b.Link,
							b.Icon(
								e.ImgSrc(
									"https://bulma.io/assets/brand/Bulma%20Icon%20White.svg",
									e.Styles{"height": "1em"},
								),
							),
							"Bulma doc",
						),
					),
				),
			),
		)

		menu := b.Box(
			e.Styles{
				"position":    "fixed",
				"height":      "100vh",
				"overflow-y":  "auto",
				"width":       "11.25rem",
				"top":         "0",
				"left":        "0",
				"padding-top": "3.25rem",
			},
			b.PaddingHorizontal(b.Spacing0),
			navMenu(sections, p.Path),
		)

		p.prepared = b.HTML(
			b.Script(path.Join(p.BaseURL, "htmx.min.js")),
			b.HTitle(p.Title+" | Bulma-Gomponents"),
			b.CSSPath(path.Join(p.BaseURL, "bulma.css")),
			fa.CSSHead(path.Join(p.BaseURL, "fa")),
			gomponents.Attr("hx-on:htmx:load", `if(event.detail.elt.id ==='bgd-content')window.scrollTo(0,0)`),
			e.Prepare(navbar),
			e.Prepare(menu),
			e.Prepare(content),
		)
	})

	return p.prepared
}

func (p *Page) MenuEntry(activePath string) e.Element {
	a := e.AHref(
		p.URL(),
		gomponents.Attr("hx-get", p.URL()),
		p.menuentry,
	)
	if activePath == p.Path {
		a.With(b.Active)
	}

	return b.MenuEntry(a)
}

func (p *Page) Section(title, peerURL string, content ...any) *Page {
	return p.section(1, title, peerURL, content...)
}

func (p *Page) Subsection(title, peerURL string, content ...any) *Page {
	return p.section(2, title, peerURL, content...)
}

func (p *Page) section(level int, title, peerURL string, content ...any) *Page {
	titleSlug := slug(title)

	section := b.Block(e.ID(titleSlug))

	peerName := "Bulma"

	if strings.Contains(peerURL, "::") {
		splat := strings.SplitN(peerURL, "::", 2)
		peerName = splat[0]
		peerURL = splat[1]
	}

	switch level {
	case 1:
		section.With(
			b.Title4(
				html.H2,
				title,
				" ",
				e.AHref(
					"#"+titleSlug,
					b.FontSize5,
					fa.Icon(fa.Solid, "link"),
				),
			),
		)

		p.internalMenu = append(
			p.internalMenu,
			b.MenuEntry(
				e.AHref(
					"#"+titleSlug,
					e.ID("bgd-nav-"+titleSlug),
					title,
				),
			),
		)
	case 2:
		section.With(
			b.Title5(
				e.ID(titleSlug),
				html.H3,
				title,
				" ",
				e.AHref(
					"#"+titleSlug,
					b.FontSize6,
					e.Name(titleSlug),
					fa.Icon(fa.Solid, "link"),
				),
			),
		)

		p.internalMenu = append(
			p.internalMenu,
			b.MenuEntry(
				e.AHref("#"+titleSlug,
					e.ID("bgd-nav-"+titleSlug),
					fa.Icon(
						fa.Solid, "caret-right",
						e.Styles{"height": "1em", "width": "1em"},
						fa.SizeSm,
					),
					title,
				),
			),
		)
	}
	if peerURL != "" {
		section.With(
			b.Content(e.AHref(peerURL, e.Target("_blank"), peerName+" documentation")),
		)
	}
	section.With(content...)

	p.internalMenuScript += ";new IntersectionObserver(([entry])=>{d=document.getElementById('bgd-nav-" + titleSlug + "');if(!d)return;entry.isIntersecting?d.classList.add(acnc):d.classList.remove(acnc)}).observe(document.getElementById('" + titleSlug + "'))"
	p.Children = append(p.Children, section)

	return p
}

func (p *Page) InternalMenu() []any {
	if len(p.internalMenu) == 0 {
		return nil
	}

	return []any{
		e.Script(gomponents.Raw(p.internalMenuScript)),
		e.Styles{"margin-right": "9rem"},
		b.Box(
			e.Styles{
				"position":    "fixed",
				"height":      "100vh",
				"overflow-y":  "auto",
				"width":       "9rem",
				"top":         "0",
				"right":       "0",
				"padding-top": "3.25rem",
			},
			b.PaddingHorizontal(b.Spacing0),
			b.Menu(
				b.FontSize7,
				b.MenuList(
					b.MenuEntry(
						e.Strong(
							e.AHref(
								"#top",
								p.Title,
							),
						),
					),
					p.internalMenu,
				),
			),
		),
	}
}

func slug(src string) string {
	dst := strings.TrimSpace(src)
	dst = strings.ToLower(dst)
	dst = strings.ReplaceAll(dst, " ", "-")
	return dst
}

func navMenu(sections []DocSection, currentPath string) e.Element {
	navmenu := b.Menu(
		e.ID("bgd-menu"),
		gomponents.Attr("hx-select-oob", "#bgd-menu,#bgd-content"),
		gomponents.Attr("hx-push-url", "true"),
	)

	for _, section := range sections {
		if section.Title != "" {
			navmenu.With(b.MenuLabel(
				b.PaddingLeft(b.Spacing2),
				section.Title,
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