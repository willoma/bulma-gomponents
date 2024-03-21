package bulma

type Styles map[string]string

// Style applies styles to the element.
func Style(propertiesValues ...string) Styles {
	stylesObj := Styles{}
	var limit int
	if len(propertiesValues)%2 == 0 {
		limit = len(propertiesValues)
	} else {
		limit = len(propertiesValues) - 1
	}
	for i := 0; i < limit; i += 2 {
		stylesObj[propertiesValues[i]] = propertiesValues[i+1]
	}
	return stylesObj
}
