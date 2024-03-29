package components

import (
	"path"
	"strings"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/el"
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
}

func (p *Page) URL() string {
	return path.Join(p.BaseURL, p.Path)
}

func NewPage(menuentry, title, path, bulmaURL string, content ...any) *Page {
	return &Page{
		menuentry:          menuentry,
		Title:              title,
		Path:               path,
		BaseURL:            "/",
		BulmaURL:           bulmaURL,
		Children:           content,
		internalMenuScript: "acnc='has-background-link-light'", // ACtive Navigation Class
	}
}

func (p *Page) MenuEntry(activePath string) b.Element {
	a := b.AHref(
		p.URL(),
		gomponents.Attr("hx-get", p.URL()),
		p.menuentry,
	)
	if activePath == p.Path {
		a.With(b.Active)
	}

	return b.MenuEntry(a)
}

func (p *Page) Section(title, bulmaURL string, content ...any) *Page {
	return p.section(1, title, bulmaURL, content...)
}

func (p *Page) Subsection(title, bulmaURL string, content ...any) *Page {
	return p.section(2, title, bulmaURL, content...)
}

func (p *Page) section(level int, title, bulmaURL string, content ...any) *Page {
	titleSlug := slug(title)

	section := b.Block(html.ID(titleSlug))

	switch level {
	case 1:
		section.With(
			b.Title4(
				html.H2,
				title,
				" ",
				b.AHref(
					"#"+titleSlug,
					b.FontSize5,
					fa.Icon(fa.Solid, "link"),
				),
			),
		)

		p.internalMenu = append(
			p.internalMenu,
			b.MenuEntry(
				b.AHref(
					"#"+titleSlug,
					html.ID("bgd-nav-"+titleSlug),
					title,
				),
			),
		)
	case 2:
		section.With(
			b.Title5(
				html.ID(titleSlug),
				html.H3,
				title,
				" ",
				b.AHref(
					"#"+titleSlug,
					b.FontSize6,
					html.Name(titleSlug),
					fa.Icon(fa.Solid, "link"),
				),
			),
		)

		p.internalMenu = append(
			p.internalMenu,
			b.MenuEntry(
				b.AHref("#"+titleSlug,
					html.ID("bgd-nav-"+titleSlug),
					fa.Icon(
						fa.Solid, "caret-right",
						b.Style("height", "1em", "width", "1em"),
						fa.SizeSm,
					),
					title,
				),
			),
		)
	}
	if bulmaURL != "" {
		section.With(
			b.Content(b.AHref(bulmaURL, html.Target("_blank"), "Bulma documentation")),
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
		html.Script(gomponents.Raw(p.internalMenuScript)),
		b.Style("margin-right", "9rem"),
		b.Box(
			b.Style(
				"position", "fixed",
				"height", "100vh",
				"overflow-y", "auto",
				"width", "9rem",
				"top", "0",
				"right", "0",
				"padding-top", "3.25rem",
			),
			b.PaddingHorizontal(b.Spacing0),
			b.Menu(
				b.FontSize7,
				b.MenuList(
					b.MenuEntry(
						el.Strong(
							b.AHref(
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
