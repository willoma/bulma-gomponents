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
	menuentry    string
	Title        string
	Path         string
	BaseURL      string
	BulmaURL     string
	Children     []any
	internalMenu []any
}

func (p *Page) URL() string {
	if p.BaseURL == "" {
		return p.Path
	}
	return path.Join(p.BaseURL, p.Path)
}

func NewPage(menuentry, title, path, bulmaURL string, content ...any) *Page {
	return &Page{
		menuentry: menuentry,
		Title:     title,
		Path:      path,
		BulmaURL:  bulmaURL,
		Children:  content,
	}
}

func (p *Page) MenuEntry(activePath string) *b.Element {
	var active any
	if activePath == p.Path {
		active = b.Active
	}

	return b.MenuEntry(
		b.AHref(
			p.URL(),
			p.menuentry,
			active,
			gomponents.Attr("hx-get", p.URL()),
			gomponents.Attr("hx-push-url", p.URL()),
			gomponents.Attr("hx-select-oob", "#menu,#content"),
		),
	)
}

func (p *Page) Section(title, bulmaURL string, content ...any) *Page {
	titleSlug := slug(title)

	p.Children = append(
		p.Children,
		gomponents.Group([]gomponents.Node{
			b.Title4(
				html.H2,
				title,
				" ",
				b.AHref(
					"#"+titleSlug,
					b.FontSize5,
					html.Name(titleSlug),
					fa.Icon(fa.Solid, "link"),
				),
			),
			b.Content(b.AHref(bulmaURL, html.Target("_blank"), "Bulma documentation")),
		}),
	)
	p.Children = append(p.Children, content...)

	p.internalMenu = append(
		p.internalMenu,
		b.MenuEntry(b.AHref("#"+titleSlug, title)),
	)
	return p
}

func (p *Page) InternalMenu() []any {
	if len(p.internalMenu) == 0 {
		return nil
	}

	return []any{
		b.Style("margin-right", "9rem"),
		b.Box(
			b.Style(
				"position", "fixed",
				"height", "100vh",
				"overflow-y", "auto",
				"width", "9rem",
				"top", "0",
				"right", "0",
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
