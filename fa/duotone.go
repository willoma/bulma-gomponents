package fa

import (
	"strconv"

	b "github.com/willoma/bulma-gomponents"
)

const (
	SwapOpacity = Class("fa-swap-opacity")
)

func PrimaryOpacity(value float64) b.Styles {
	return b.Style("--fa-primary-opacity", strconv.FormatFloat(value, 'f', 2, 64))
}

func SecondaryOpacity(value float64) b.Styles {
	return b.Style("--fa-secondary-opacity", strconv.FormatFloat(value, 'f', 2, 64))
}

func PrimaryColor(value string) b.Styles {
	return b.Style("--fa-primary-color", value)
}

func SecondaryColor(value string) b.Styles {
	return b.Style("--fa-secondary-color", value)
}
