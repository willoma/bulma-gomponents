package fa

import (
	"strconv"

	e "github.com/willoma/gomplements"
)

const (
	SwapOpacity = Class("fa-swap-opacity")
)

func PrimaryOpacity(value float64) e.Styles {
	return e.Styles{"--fa-primary-opacity": strconv.FormatFloat(value, 'f', 2, 64)}
}

func SecondaryOpacity(value float64) e.Styles {
	return e.Styles{"--fa-secondary-opacity": strconv.FormatFloat(value, 'f', 2, 64)}
}

func DuoOpacities(primary, secondary float64) e.Styles {
	return e.Styles{
		"--fa-primary-opacity":   strconv.FormatFloat(primary, 'f', 2, 64),
		"--fa-secondary-opacity": strconv.FormatFloat(secondary, 'f', 2, 64),
	}
}

func PrimaryColor(value string) e.Styles {
	return e.Styles{"--fa-primary-color": value}
}

func SecondaryColor(value string) e.Styles {
	return e.Styles{"--fa-secondary-color": value}
}

func DuoColors(primary, secondary string) e.Styles {
	return e.Styles{
		"--fa-primary-color":   primary,
		"--fa-secondary-color": secondary,
	}
}
