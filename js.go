package bulma

// JSClose returns javascript code that removed the 'is-active' class from the
// element with the provided ID.
func JSClose(id string) string {
	return "getElementById('" + id + "').classList.remove('is-active')"
}

// JSOpen returns javascript code that adds the 'is-active' class to the
// element with the provided ID.
func JSOpen(id string) string {
	return "getElementById('" + id + "').classList.add('is-active')"
}

// JSToggle returns javascript code that toggles the 'is-active' class on the
// element with the provided ID.
func JSToggle(id string) string {
	return "getElementById('" + id + "').classList.toggle('is-active')"
}

// Javascript code which applied to the current element, from any of its children.
const (
	JSCloseThisDropdown  = "this.closest('.dropdown').classList.remove('is-active')"
	JSOpenThisDropdown   = "this.closest('.dropdown').classList.add('is-active')"
	JSToggleThisDropdown = "this.closest('.dropdown').classList.toggle('is-active')"

	JSCloseThisModal  = "this.closest('.modal').classList.remove('is-active')"
	JSOpenThisModal   = "this.closest('.modal').classList.add('is-active')"
	JSToggleThisModal = "this.closest('.modal').classList.toggle('is-active')"

	JSCloseThisNavbarItem  = "this.closest('.navbar-item').classList.remove('is-active')"
	JSOpenThisNavbarItem   = "this.closest('.navbar-item').classList.add('is-active')"
	JSToggleThisNavbarItem = "this.closest('.navbar-item').classList.toggle('is-active')"

	JSRemoveThisNotification = "this.closest('.notification').parentNode.removeChild(this.closest('.notification'))"

	JSCloseMe  = "this.classList.remove('is-active')"
	JSOpenMe   = "this.classList.add('is-active')"
	JSToggleMe = "this.classList.toggle('is-active')"
)
