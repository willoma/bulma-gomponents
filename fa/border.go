package fa

import b "github.com/willoma/bulma-gomponents"

const (
	Border    = Class("fa-border")
	PullLeft  = Class("fa-pull-left")
	PullRight = Class("fa-pull-right")
)

func BorderColor(value string) b.Styles {
	return b.Style("--fa-border-color", value)
}

func BorderPadding(value string) b.Styles {
	return b.Style("--fa-border-padding", value)
}

func BorderRadius(value string) b.Styles {
	return b.Style("--fa-border-radius", value)
}

func BorderStyle(value string) b.Styles {
	return b.Style("--fa-border-style", value)
}

func BorderWidth(value string) b.Styles {
	return b.Style("--fa-border-width", value)
}

func PullMargin(value string) b.Styles {
	return b.Style("--fa-pull-margin", value)
}
