package easy

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
)

// ModalCardHeadTitle creates a card head with a text title and a close button.
func ModalCardHeadTitle(title string) any {
	return b.ModalCardHead(
		b.ModalCardTitle(title),
		b.Delete(
			b.OnClick(b.JSCloseThisModal),
			html.Aria("label", "close"),
		),
	)
}
