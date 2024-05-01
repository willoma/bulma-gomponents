package fa

import e "github.com/willoma/gomplements"

const (
	Border    = Class("fa-border")
	PullLeft  = Class("fa-pull-left")
	PullRight = Class("fa-pull-right")
)

func BorderColor(value string) e.Styles {
	return e.Styles{"--fa-border-color": value}
}

func BorderPadding(value string) e.Styles {
	return e.Styles{"--fa-border-padding": value}
}

func BorderRadius(value string) e.Styles {
	return e.Styles{"--fa-border-radius": value}
}

func BorderStyle(value string) e.Styles {
	return e.Styles{"--fa-border-style": value}
}

func BorderWidth(value string) e.Styles {
	return e.Styles{"--fa-border-width": value}
}

func PullMargin(value string) e.Styles {
	return e.Styles{"--fa-pull-margin": value}
}
