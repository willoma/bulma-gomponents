package easy

import b "github.com/willoma/bulma-gomponents"

// Menu creates a menu.
//   - when a child is a string, it is wrapped in a menu label
//   - when a child is anything else, it is added as an entry in a menu list
func Menu(children ...any) b.Element {
	menu := b.Menu()

	var currentList []any

	for _, c := range children {
		switch c := c.(type) {
		case string:
			if len(currentList) > 0 {
				menu.With(b.MenuList(currentList...))
			}
			menu.With(b.MenuLabel(c))
			currentList = currentList[:0]
		case []any:
			currentList = append(currentList, b.MenuEntry(c...))
		default:
			currentList = append(currentList, b.MenuEntry(c))
		}
	}

	if len(currentList) > 0 {
		menu.With(b.MenuList(currentList...))
	}

	return menu
}

// MenuEntryWithSublist creates a menu sublist, associated with a menu entry.
//
// Each of its children is included as a submenu entry.
func MenuEntryWithSublist(entry any, children ...any) []any {
	list := b.MenuSublist()
	for _, c := range children {
		switch c := c.(type) {
		case []any:
			list.With(b.MenuEntry(c...))
		default:
			list.With(b.MenuEntry(c))
		}
	}

	return []any{entry, list}
}
