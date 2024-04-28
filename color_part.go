package bulma

import "fmt"

type partColor struct {
	name    string
	text    bool
	variant string
	invert  bool
}

func (c partColor) Class() Class {
	class := "has-"

	if c.text {
		class += "text-"
	} else {
		class += "background-"
	}

	class += c.name

	if c.variant != "" {
		class += "-" + c.variant
	}

	if c.invert {
		class += "-invert"
	}

	return Class(class)
}

func (c partColor) Base() Color {
	return baseColor{
		name:    c.name,
		variant: c.variant,
	}
}

func (c partColor) Text() Color {
	return partColor{
		name:    c.name,
		text:    true,
		variant: c.variant,
		invert:  c.invert,
	}
}

func (c partColor) Background() Color {
	return partColor{
		name:    c.name,
		text:    false,
		variant: c.variant,
		invert:  c.invert,
	}
}

func (c partColor) Normal() Color {
	return partColor{
		name:   c.name,
		text:   c.text,
		invert: c.invert,
	}
}

func (c partColor) Light() Color {
	return partColor{
		name:    c.name,
		text:    c.text,
		variant: colorLight,
		invert:  c.invert,
	}
}

func (c partColor) Dark() Color {
	return partColor{
		name:    c.name,
		text:    c.text,
		variant: colorDark,
		invert:  c.invert,
	}
}

func (c partColor) Soft() Color {
	return partColor{
		name:    c.name,
		text:    c.text,
		variant: colorSoft,
		invert:  c.invert,
	}
}

func (c partColor) Bold() Color {
	return partColor{
		name:    c.name,
		text:    c.text,
		variant: colorBold,
		invert:  c.invert,
	}
}

func (c partColor) OnScheme() Color {
	return partColor{
		name:    c.name,
		text:    c.text,
		variant: colorOnScheme,
		invert:  c.invert,
	}
}

func (c partColor) Shade(shade int) Color {
	return partColor{
		name:    c.name,
		text:    c.text,
		variant: fmt.Sprintf("%02d", shade),
		invert:  c.invert,
	}
}

func (c partColor) Invert() Color {
	return partColor{
		name:    c.name,
		text:    c.text,
		variant: c.variant,
		invert:  !c.invert,
	}
}

var (
	TextWhite       Color = partColor{name: "white", text: true}
	TextWhiteInvert Color = partColor{name: "white", text: true, invert: true}
	TextBlack       Color = partColor{name: "black", text: true}
	TextBlackInvert Color = partColor{name: "black", text: true, invert: true}
	TextLight       Color = partColor{name: "light", text: true}
	TextLightInvert Color = partColor{name: "light", text: true, invert: true}
	TextDark        Color = partColor{name: "dark", text: true}
	TextDarkInvert  Color = partColor{name: "dark", text: true, invert: true}

	TextText               Color = partColor{name: "text", text: true}
	TextTextInvert         Color = partColor{name: "text", text: true, invert: true}
	TextText00             Color = partColor{name: "text", variant: "00", text: true}
	TextText00Invert       Color = partColor{name: "text", variant: "00", text: true, invert: true}
	TextText05             Color = partColor{name: "text", variant: "05", text: true}
	TextText05Invert       Color = partColor{name: "text", variant: "05", text: true, invert: true}
	TextText10             Color = partColor{name: "text", variant: "10", text: true}
	TextText10Invert       Color = partColor{name: "text", variant: "10", text: true, invert: true}
	TextText15             Color = partColor{name: "text", variant: "15", text: true}
	TextText15Invert       Color = partColor{name: "text", variant: "15", text: true, invert: true}
	TextText20             Color = partColor{name: "text", variant: "20", text: true}
	TextText20Invert       Color = partColor{name: "text", variant: "20", text: true, invert: true}
	TextText25             Color = partColor{name: "text", variant: "25", text: true}
	TextText25Invert       Color = partColor{name: "text", variant: "25", text: true, invert: true}
	TextText30             Color = partColor{name: "text", variant: "30", text: true}
	TextText30Invert       Color = partColor{name: "text", variant: "30", text: true, invert: true}
	TextText35             Color = partColor{name: "text", variant: "35", text: true}
	TextText35Invert       Color = partColor{name: "text", variant: "35", text: true, invert: true}
	TextText40             Color = partColor{name: "text", variant: "40", text: true}
	TextText40Invert       Color = partColor{name: "text", variant: "40", text: true, invert: true}
	TextText45             Color = partColor{name: "text", variant: "45", text: true}
	TextText45Invert       Color = partColor{name: "text", variant: "45", text: true, invert: true}
	TextText50             Color = partColor{name: "text", variant: "50", text: true}
	TextText50Invert       Color = partColor{name: "text", variant: "50", text: true, invert: true}
	TextText55             Color = partColor{name: "text", variant: "55", text: true}
	TextText55Invert       Color = partColor{name: "text", variant: "55", text: true, invert: true}
	TextText60             Color = partColor{name: "text", variant: "60", text: true}
	TextText60Invert       Color = partColor{name: "text", variant: "60", text: true, invert: true}
	TextText65             Color = partColor{name: "text", variant: "65", text: true}
	TextText65Invert       Color = partColor{name: "text", variant: "65", text: true, invert: true}
	TextText70             Color = partColor{name: "text", variant: "70", text: true}
	TextText70Invert       Color = partColor{name: "text", variant: "70", text: true, invert: true}
	TextText75             Color = partColor{name: "text", variant: "75", text: true}
	TextText75Invert       Color = partColor{name: "text", variant: "75", text: true, invert: true}
	TextText80             Color = partColor{name: "text", variant: "80", text: true}
	TextText80Invert       Color = partColor{name: "text", variant: "80", text: true, invert: true}
	TextText85             Color = partColor{name: "text", variant: "85", text: true}
	TextText85Invert       Color = partColor{name: "text", variant: "85", text: true, invert: true}
	TextText90             Color = partColor{name: "text", variant: "90", text: true}
	TextText90Invert       Color = partColor{name: "text", variant: "90", text: true, invert: true}
	TextText95             Color = partColor{name: "text", variant: "95", text: true}
	TextText95Invert       Color = partColor{name: "text", variant: "95", text: true, invert: true}
	TextText100            Color = partColor{name: "text", variant: "100", text: true}
	TextText100Invert      Color = partColor{name: "text", variant: "100", text: true, invert: true}
	TextTextLight          Color = partColor{name: "text", text: true, variant: colorLight}
	TextTextLightInvert    Color = partColor{name: "text", text: true, variant: colorLight, invert: true}
	TextTextDark           Color = partColor{name: "text", text: true, variant: colorDark}
	TextTextDarkInvert     Color = partColor{name: "text", text: true, variant: colorDark, invert: true}
	TextTextSoft           Color = partColor{name: "text", text: true, variant: colorSoft}
	TextTextSoftInvert     Color = partColor{name: "text", text: true, variant: colorSoft, invert: true}
	TextTextBold           Color = partColor{name: "text", text: true, variant: colorBold}
	TextTextBoldInvert     Color = partColor{name: "text", text: true, variant: colorBold, invert: true}
	TextTextOnScheme       Color = partColor{name: "text", text: true, variant: colorOnScheme}
	TextTextOnSchemeInvert Color = partColor{name: "text", text: true, variant: colorOnScheme, invert: true}

	TextPrimary               Color = partColor{name: "primary", text: true}
	TextPrimaryInvert         Color = partColor{name: "primary", text: true, invert: true}
	TextPrimary00             Color = partColor{name: "primary", variant: "00", text: true}
	TextPrimary00Invert       Color = partColor{name: "primary", variant: "00", text: true, invert: true}
	TextPrimary05             Color = partColor{name: "primary", variant: "05", text: true}
	TextPrimary05Invert       Color = partColor{name: "primary", variant: "05", text: true, invert: true}
	TextPrimary10             Color = partColor{name: "primary", variant: "10", text: true}
	TextPrimary10Invert       Color = partColor{name: "primary", variant: "10", text: true, invert: true}
	TextPrimary15             Color = partColor{name: "primary", variant: "15", text: true}
	TextPrimary15Invert       Color = partColor{name: "primary", variant: "15", text: true, invert: true}
	TextPrimary20             Color = partColor{name: "primary", variant: "20", text: true}
	TextPrimary20Invert       Color = partColor{name: "primary", variant: "20", text: true, invert: true}
	TextPrimary25             Color = partColor{name: "primary", variant: "25", text: true}
	TextPrimary25Invert       Color = partColor{name: "primary", variant: "25", text: true, invert: true}
	TextPrimary30             Color = partColor{name: "primary", variant: "30", text: true}
	TextPrimary30Invert       Color = partColor{name: "primary", variant: "30", text: true, invert: true}
	TextPrimary35             Color = partColor{name: "primary", variant: "35", text: true}
	TextPrimary35Invert       Color = partColor{name: "primary", variant: "35", text: true, invert: true}
	TextPrimary40             Color = partColor{name: "primary", variant: "40", text: true}
	TextPrimary40Invert       Color = partColor{name: "primary", variant: "40", text: true, invert: true}
	TextPrimary45             Color = partColor{name: "primary", variant: "45", text: true}
	TextPrimary45Invert       Color = partColor{name: "primary", variant: "45", text: true, invert: true}
	TextPrimary50             Color = partColor{name: "primary", variant: "50", text: true}
	TextPrimary50Invert       Color = partColor{name: "primary", variant: "50", text: true, invert: true}
	TextPrimary55             Color = partColor{name: "primary", variant: "55", text: true}
	TextPrimary55Invert       Color = partColor{name: "primary", variant: "55", text: true, invert: true}
	TextPrimary60             Color = partColor{name: "primary", variant: "60", text: true}
	TextPrimary60Invert       Color = partColor{name: "primary", variant: "60", text: true, invert: true}
	TextPrimary65             Color = partColor{name: "primary", variant: "65", text: true}
	TextPrimary65Invert       Color = partColor{name: "primary", variant: "65", text: true, invert: true}
	TextPrimary70             Color = partColor{name: "primary", variant: "70", text: true}
	TextPrimary70Invert       Color = partColor{name: "primary", variant: "70", text: true, invert: true}
	TextPrimary75             Color = partColor{name: "primary", variant: "75", text: true}
	TextPrimary75Invert       Color = partColor{name: "primary", variant: "75", text: true, invert: true}
	TextPrimary80             Color = partColor{name: "primary", variant: "80", text: true}
	TextPrimary80Invert       Color = partColor{name: "primary", variant: "80", text: true, invert: true}
	TextPrimary85             Color = partColor{name: "primary", variant: "85", text: true}
	TextPrimary85Invert       Color = partColor{name: "primary", variant: "85", text: true, invert: true}
	TextPrimary90             Color = partColor{name: "primary", variant: "90", text: true}
	TextPrimary90Invert       Color = partColor{name: "primary", variant: "90", text: true, invert: true}
	TextPrimary95             Color = partColor{name: "primary", variant: "95", text: true}
	TextPrimary95Invert       Color = partColor{name: "primary", variant: "95", text: true, invert: true}
	TextPrimary100            Color = partColor{name: "primary", variant: "100", text: true}
	TextPrimary100Invert      Color = partColor{name: "primary", variant: "100", text: true, invert: true}
	TextPrimaryLight          Color = partColor{name: "primary", text: true, variant: colorLight}
	TextPrimaryLightInvert    Color = partColor{name: "primary", text: true, variant: colorLight, invert: true}
	TextPrimaryDark           Color = partColor{name: "primary", text: true, variant: colorDark}
	TextPrimaryDarkInvert     Color = partColor{name: "primary", text: true, variant: colorDark, invert: true}
	TextPrimarySoft           Color = partColor{name: "primary", text: true, variant: colorSoft}
	TextPrimarySoftInvert     Color = partColor{name: "primary", text: true, variant: colorSoft, invert: true}
	TextPrimaryBold           Color = partColor{name: "primary", text: true, variant: colorBold}
	TextPrimaryBoldInvert     Color = partColor{name: "primary", text: true, variant: colorBold, invert: true}
	TextPrimaryOnScheme       Color = partColor{name: "primary", text: true, variant: colorOnScheme}
	TextPrimaryOnSchemeInvert Color = partColor{name: "primary", text: true, variant: colorOnScheme, invert: true}

	TextLink               Color = partColor{name: "link", text: true}
	TextLinkInvert         Color = partColor{name: "link", text: true, invert: true}
	TextLink00             Color = partColor{name: "link", variant: "00", text: true}
	TextLink00Invert       Color = partColor{name: "link", variant: "00", text: true, invert: true}
	TextLink05             Color = partColor{name: "link", variant: "05", text: true}
	TextLink05Invert       Color = partColor{name: "link", variant: "05", text: true, invert: true}
	TextLink10             Color = partColor{name: "link", variant: "10", text: true}
	TextLink10Invert       Color = partColor{name: "link", variant: "10", text: true, invert: true}
	TextLink15             Color = partColor{name: "link", variant: "15", text: true}
	TextLink15Invert       Color = partColor{name: "link", variant: "15", text: true, invert: true}
	TextLink20             Color = partColor{name: "link", variant: "20", text: true}
	TextLink20Invert       Color = partColor{name: "link", variant: "20", text: true, invert: true}
	TextLink25             Color = partColor{name: "link", variant: "25", text: true}
	TextLink25Invert       Color = partColor{name: "link", variant: "25", text: true, invert: true}
	TextLink30             Color = partColor{name: "link", variant: "30", text: true}
	TextLink30Invert       Color = partColor{name: "link", variant: "30", text: true, invert: true}
	TextLink35             Color = partColor{name: "link", variant: "35", text: true}
	TextLink35Invert       Color = partColor{name: "link", variant: "35", text: true, invert: true}
	TextLink40             Color = partColor{name: "link", variant: "40", text: true}
	TextLink40Invert       Color = partColor{name: "link", variant: "40", text: true, invert: true}
	TextLink45             Color = partColor{name: "link", variant: "45", text: true}
	TextLink45Invert       Color = partColor{name: "link", variant: "45", text: true, invert: true}
	TextLink50             Color = partColor{name: "link", variant: "50", text: true}
	TextLink50Invert       Color = partColor{name: "link", variant: "50", text: true, invert: true}
	TextLink55             Color = partColor{name: "link", variant: "55", text: true}
	TextLink55Invert       Color = partColor{name: "link", variant: "55", text: true, invert: true}
	TextLink60             Color = partColor{name: "link", variant: "60", text: true}
	TextLink60Invert       Color = partColor{name: "link", variant: "60", text: true, invert: true}
	TextLink65             Color = partColor{name: "link", variant: "65", text: true}
	TextLink65Invert       Color = partColor{name: "link", variant: "65", text: true, invert: true}
	TextLink70             Color = partColor{name: "link", variant: "70", text: true}
	TextLink70Invert       Color = partColor{name: "link", variant: "70", text: true, invert: true}
	TextLink75             Color = partColor{name: "link", variant: "75", text: true}
	TextLink75Invert       Color = partColor{name: "link", variant: "75", text: true, invert: true}
	TextLink80             Color = partColor{name: "link", variant: "80", text: true}
	TextLink80Invert       Color = partColor{name: "link", variant: "80", text: true, invert: true}
	TextLink85             Color = partColor{name: "link", variant: "85", text: true}
	TextLink85Invert       Color = partColor{name: "link", variant: "85", text: true, invert: true}
	TextLink90             Color = partColor{name: "link", variant: "90", text: true}
	TextLink90Invert       Color = partColor{name: "link", variant: "90", text: true, invert: true}
	TextLink95             Color = partColor{name: "link", variant: "95", text: true}
	TextLink95Invert       Color = partColor{name: "link", variant: "95", text: true, invert: true}
	TextLink100            Color = partColor{name: "link", variant: "100", text: true}
	TextLink100Invert      Color = partColor{name: "link", variant: "100", text: true, invert: true}
	TextLinkLight          Color = partColor{name: "link", text: true, variant: colorLight}
	TextLinkLightInvert    Color = partColor{name: "link", text: true, variant: colorLight, invert: true}
	TextLinkDark           Color = partColor{name: "link", text: true, variant: colorDark}
	TextLinkDarkInvert     Color = partColor{name: "link", text: true, variant: colorDark, invert: true}
	TextLinkSoft           Color = partColor{name: "link", text: true, variant: colorSoft}
	TextLinkSoftInvert     Color = partColor{name: "link", text: true, variant: colorSoft, invert: true}
	TextLinkBold           Color = partColor{name: "link", text: true, variant: colorBold}
	TextLinkBoldInvert     Color = partColor{name: "link", text: true, variant: colorBold, invert: true}
	TextLinkOnScheme       Color = partColor{name: "link", text: true, variant: colorOnScheme}
	TextLinkOnSchemeInvert Color = partColor{name: "link", text: true, variant: colorOnScheme, invert: true}

	TextInfo               Color = partColor{name: "info", text: true}
	TextInfoInvert         Color = partColor{name: "info", text: true, invert: true}
	TextInfo00             Color = partColor{name: "info", variant: "00", text: true}
	TextInfo00Invert       Color = partColor{name: "info", variant: "00", text: true, invert: true}
	TextInfo05             Color = partColor{name: "info", variant: "05", text: true}
	TextInfo05Invert       Color = partColor{name: "info", variant: "05", text: true, invert: true}
	TextInfo10             Color = partColor{name: "info", variant: "10", text: true}
	TextInfo10Invert       Color = partColor{name: "info", variant: "10", text: true, invert: true}
	TextInfo15             Color = partColor{name: "info", variant: "15", text: true}
	TextInfo15Invert       Color = partColor{name: "info", variant: "15", text: true, invert: true}
	TextInfo20             Color = partColor{name: "info", variant: "20", text: true}
	TextInfo20Invert       Color = partColor{name: "info", variant: "20", text: true, invert: true}
	TextInfo25             Color = partColor{name: "info", variant: "25", text: true}
	TextInfo25Invert       Color = partColor{name: "info", variant: "25", text: true, invert: true}
	TextInfo30             Color = partColor{name: "info", variant: "30", text: true}
	TextInfo30Invert       Color = partColor{name: "info", variant: "30", text: true, invert: true}
	TextInfo35             Color = partColor{name: "info", variant: "35", text: true}
	TextInfo35Invert       Color = partColor{name: "info", variant: "35", text: true, invert: true}
	TextInfo40             Color = partColor{name: "info", variant: "40", text: true}
	TextInfo40Invert       Color = partColor{name: "info", variant: "40", text: true, invert: true}
	TextInfo45             Color = partColor{name: "info", variant: "45", text: true}
	TextInfo45Invert       Color = partColor{name: "info", variant: "45", text: true, invert: true}
	TextInfo50             Color = partColor{name: "info", variant: "50", text: true}
	TextInfo50Invert       Color = partColor{name: "info", variant: "50", text: true, invert: true}
	TextInfo55             Color = partColor{name: "info", variant: "55", text: true}
	TextInfo55Invert       Color = partColor{name: "info", variant: "55", text: true, invert: true}
	TextInfo60             Color = partColor{name: "info", variant: "60", text: true}
	TextInfo60Invert       Color = partColor{name: "info", variant: "60", text: true, invert: true}
	TextInfo65             Color = partColor{name: "info", variant: "65", text: true}
	TextInfo65Invert       Color = partColor{name: "info", variant: "65", text: true, invert: true}
	TextInfo70             Color = partColor{name: "info", variant: "70", text: true}
	TextInfo70Invert       Color = partColor{name: "info", variant: "70", text: true, invert: true}
	TextInfo75             Color = partColor{name: "info", variant: "75", text: true}
	TextInfo75Invert       Color = partColor{name: "info", variant: "75", text: true, invert: true}
	TextInfo80             Color = partColor{name: "info", variant: "80", text: true}
	TextInfo80Invert       Color = partColor{name: "info", variant: "80", text: true, invert: true}
	TextInfo85             Color = partColor{name: "info", variant: "85", text: true}
	TextInfo85Invert       Color = partColor{name: "info", variant: "85", text: true, invert: true}
	TextInfo90             Color = partColor{name: "info", variant: "90", text: true}
	TextInfo90Invert       Color = partColor{name: "info", variant: "90", text: true, invert: true}
	TextInfo95             Color = partColor{name: "info", variant: "95", text: true}
	TextInfo95Invert       Color = partColor{name: "info", variant: "95", text: true, invert: true}
	TextInfo100            Color = partColor{name: "info", variant: "100", text: true}
	TextInfo100Invert      Color = partColor{name: "info", variant: "100", text: true, invert: true}
	TextInfoLight          Color = partColor{name: "info", text: true, variant: colorLight}
	TextInfoLightInvert    Color = partColor{name: "info", text: true, variant: colorLight, invert: true}
	TextInfoDark           Color = partColor{name: "info", text: true, variant: colorDark}
	TextInfoDarkInvert     Color = partColor{name: "info", text: true, variant: colorDark, invert: true}
	TextInfoSoft           Color = partColor{name: "info", text: true, variant: colorSoft}
	TextInfoSoftInvert     Color = partColor{name: "info", text: true, variant: colorSoft, invert: true}
	TextInfoBold           Color = partColor{name: "info", text: true, variant: colorBold}
	TextInfoBoldInvert     Color = partColor{name: "info", text: true, variant: colorBold, invert: true}
	TextInfoOnScheme       Color = partColor{name: "info", text: true, variant: colorOnScheme}
	TextInfoOnSchemeInvert Color = partColor{name: "info", text: true, variant: colorOnScheme, invert: true}

	TextSuccess               Color = partColor{name: "success", text: true}
	TextSuccessInvert         Color = partColor{name: "success", text: true, invert: true}
	TextSuccess00             Color = partColor{name: "success", variant: "00", text: true}
	TextSuccess00Invert       Color = partColor{name: "success", variant: "00", text: true, invert: true}
	TextSuccess05             Color = partColor{name: "success", variant: "05", text: true}
	TextSuccess05Invert       Color = partColor{name: "success", variant: "05", text: true, invert: true}
	TextSuccess10             Color = partColor{name: "success", variant: "10", text: true}
	TextSuccess10Invert       Color = partColor{name: "success", variant: "10", text: true, invert: true}
	TextSuccess15             Color = partColor{name: "success", variant: "15", text: true}
	TextSuccess15Invert       Color = partColor{name: "success", variant: "15", text: true, invert: true}
	TextSuccess20             Color = partColor{name: "success", variant: "20", text: true}
	TextSuccess20Invert       Color = partColor{name: "success", variant: "20", text: true, invert: true}
	TextSuccess25             Color = partColor{name: "success", variant: "25", text: true}
	TextSuccess25Invert       Color = partColor{name: "success", variant: "25", text: true, invert: true}
	TextSuccess30             Color = partColor{name: "success", variant: "30", text: true}
	TextSuccess30Invert       Color = partColor{name: "success", variant: "30", text: true, invert: true}
	TextSuccess35             Color = partColor{name: "success", variant: "35", text: true}
	TextSuccess35Invert       Color = partColor{name: "success", variant: "35", text: true, invert: true}
	TextSuccess40             Color = partColor{name: "success", variant: "40", text: true}
	TextSuccess40Invert       Color = partColor{name: "success", variant: "40", text: true, invert: true}
	TextSuccess45             Color = partColor{name: "success", variant: "45", text: true}
	TextSuccess45Invert       Color = partColor{name: "success", variant: "45", text: true, invert: true}
	TextSuccess50             Color = partColor{name: "success", variant: "50", text: true}
	TextSuccess50Invert       Color = partColor{name: "success", variant: "50", text: true, invert: true}
	TextSuccess55             Color = partColor{name: "success", variant: "55", text: true}
	TextSuccess55Invert       Color = partColor{name: "success", variant: "55", text: true, invert: true}
	TextSuccess60             Color = partColor{name: "success", variant: "60", text: true}
	TextSuccess60Invert       Color = partColor{name: "success", variant: "60", text: true, invert: true}
	TextSuccess65             Color = partColor{name: "success", variant: "65", text: true}
	TextSuccess65Invert       Color = partColor{name: "success", variant: "65", text: true, invert: true}
	TextSuccess70             Color = partColor{name: "success", variant: "70", text: true}
	TextSuccess70Invert       Color = partColor{name: "success", variant: "70", text: true, invert: true}
	TextSuccess75             Color = partColor{name: "success", variant: "75", text: true}
	TextSuccess75Invert       Color = partColor{name: "success", variant: "75", text: true, invert: true}
	TextSuccess80             Color = partColor{name: "success", variant: "80", text: true}
	TextSuccess80Invert       Color = partColor{name: "success", variant: "80", text: true, invert: true}
	TextSuccess85             Color = partColor{name: "success", variant: "85", text: true}
	TextSuccess85Invert       Color = partColor{name: "success", variant: "85", text: true, invert: true}
	TextSuccess90             Color = partColor{name: "success", variant: "90", text: true}
	TextSuccess90Invert       Color = partColor{name: "success", variant: "90", text: true, invert: true}
	TextSuccess95             Color = partColor{name: "success", variant: "95", text: true}
	TextSuccess95Invert       Color = partColor{name: "success", variant: "95", text: true, invert: true}
	TextSuccess100            Color = partColor{name: "success", variant: "100", text: true}
	TextSuccess100Invert      Color = partColor{name: "success", variant: "100", text: true, invert: true}
	TextSuccessLight          Color = partColor{name: "success", text: true, variant: colorLight}
	TextSuccessLightInvert    Color = partColor{name: "success", text: true, variant: colorLight, invert: true}
	TextSuccessDark           Color = partColor{name: "success", text: true, variant: colorDark}
	TextSuccessDarkInvert     Color = partColor{name: "success", text: true, variant: colorDark, invert: true}
	TextSuccessSoft           Color = partColor{name: "success", text: true, variant: colorSoft}
	TextSuccessSoftInvert     Color = partColor{name: "success", text: true, variant: colorSoft, invert: true}
	TextSuccessBold           Color = partColor{name: "success", text: true, variant: colorBold}
	TextSuccessBoldInvert     Color = partColor{name: "success", text: true, variant: colorBold, invert: true}
	TextSuccessOnScheme       Color = partColor{name: "success", text: true, variant: colorOnScheme}
	TextSuccessOnSchemeInvert Color = partColor{name: "success", text: true, variant: colorOnScheme, invert: true}

	TextWarning               Color = partColor{name: "warning", text: true}
	TextWarningInvert         Color = partColor{name: "warning", text: true, invert: true}
	TextWarning00             Color = partColor{name: "warning", variant: "00", text: true}
	TextWarning00Invert       Color = partColor{name: "warning", variant: "00", text: true, invert: true}
	TextWarning05             Color = partColor{name: "warning", variant: "05", text: true}
	TextWarning05Invert       Color = partColor{name: "warning", variant: "05", text: true, invert: true}
	TextWarning10             Color = partColor{name: "warning", variant: "10", text: true}
	TextWarning10Invert       Color = partColor{name: "warning", variant: "10", text: true, invert: true}
	TextWarning15             Color = partColor{name: "warning", variant: "15", text: true}
	TextWarning15Invert       Color = partColor{name: "warning", variant: "15", text: true, invert: true}
	TextWarning20             Color = partColor{name: "warning", variant: "20", text: true}
	TextWarning20Invert       Color = partColor{name: "warning", variant: "20", text: true, invert: true}
	TextWarning25             Color = partColor{name: "warning", variant: "25", text: true}
	TextWarning25Invert       Color = partColor{name: "warning", variant: "25", text: true, invert: true}
	TextWarning30             Color = partColor{name: "warning", variant: "30", text: true}
	TextWarning30Invert       Color = partColor{name: "warning", variant: "30", text: true, invert: true}
	TextWarning35             Color = partColor{name: "warning", variant: "35", text: true}
	TextWarning35Invert       Color = partColor{name: "warning", variant: "35", text: true, invert: true}
	TextWarning40             Color = partColor{name: "warning", variant: "40", text: true}
	TextWarning40Invert       Color = partColor{name: "warning", variant: "40", text: true, invert: true}
	TextWarning45             Color = partColor{name: "warning", variant: "45", text: true}
	TextWarning45Invert       Color = partColor{name: "warning", variant: "45", text: true, invert: true}
	TextWarning50             Color = partColor{name: "warning", variant: "50", text: true}
	TextWarning50Invert       Color = partColor{name: "warning", variant: "50", text: true, invert: true}
	TextWarning55             Color = partColor{name: "warning", variant: "55", text: true}
	TextWarning55Invert       Color = partColor{name: "warning", variant: "55", text: true, invert: true}
	TextWarning60             Color = partColor{name: "warning", variant: "60", text: true}
	TextWarning60Invert       Color = partColor{name: "warning", variant: "60", text: true, invert: true}
	TextWarning65             Color = partColor{name: "warning", variant: "65", text: true}
	TextWarning65Invert       Color = partColor{name: "warning", variant: "65", text: true, invert: true}
	TextWarning70             Color = partColor{name: "warning", variant: "70", text: true}
	TextWarning70Invert       Color = partColor{name: "warning", variant: "70", text: true, invert: true}
	TextWarning75             Color = partColor{name: "warning", variant: "75", text: true}
	TextWarning75Invert       Color = partColor{name: "warning", variant: "75", text: true, invert: true}
	TextWarning80             Color = partColor{name: "warning", variant: "80", text: true}
	TextWarning80Invert       Color = partColor{name: "warning", variant: "80", text: true, invert: true}
	TextWarning85             Color = partColor{name: "warning", variant: "85", text: true}
	TextWarning85Invert       Color = partColor{name: "warning", variant: "85", text: true, invert: true}
	TextWarning90             Color = partColor{name: "warning", variant: "90", text: true}
	TextWarning90Invert       Color = partColor{name: "warning", variant: "90", text: true, invert: true}
	TextWarning95             Color = partColor{name: "warning", variant: "95", text: true}
	TextWarning95Invert       Color = partColor{name: "warning", variant: "95", text: true, invert: true}
	TextWarning100            Color = partColor{name: "warning", variant: "100", text: true}
	TextWarning100Invert      Color = partColor{name: "warning", variant: "100", text: true, invert: true}
	TextWarningLight          Color = partColor{name: "warning", text: true, variant: colorLight}
	TextWarningLightInvert    Color = partColor{name: "warning", text: true, variant: colorLight, invert: true}
	TextWarningDark           Color = partColor{name: "warning", text: true, variant: colorDark}
	TextWarningDarkInvert     Color = partColor{name: "warning", text: true, variant: colorDark, invert: true}
	TextWarningSoft           Color = partColor{name: "warning", text: true, variant: colorSoft}
	TextWarningSoftInvert     Color = partColor{name: "warning", text: true, variant: colorSoft, invert: true}
	TextWarningBold           Color = partColor{name: "warning", text: true, variant: colorBold}
	TextWarningBoldInvert     Color = partColor{name: "warning", text: true, variant: colorBold, invert: true}
	TextWarningOnScheme       Color = partColor{name: "warning", text: true, variant: colorOnScheme}
	TextWarningOnSchemeInvert Color = partColor{name: "warning", text: true, variant: colorOnScheme, invert: true}

	TextDanger               Color = partColor{name: "danger", text: true}
	TextDangerInvert         Color = partColor{name: "danger", text: true, invert: true}
	TextDanger00             Color = partColor{name: "danger", variant: "00", text: true}
	TextDanger00Invert       Color = partColor{name: "danger", variant: "00", text: true, invert: true}
	TextDanger05             Color = partColor{name: "danger", variant: "05", text: true}
	TextDanger05Invert       Color = partColor{name: "danger", variant: "05", text: true, invert: true}
	TextDanger10             Color = partColor{name: "danger", variant: "10", text: true}
	TextDanger10Invert       Color = partColor{name: "danger", variant: "10", text: true, invert: true}
	TextDanger15             Color = partColor{name: "danger", variant: "15", text: true}
	TextDanger15Invert       Color = partColor{name: "danger", variant: "15", text: true, invert: true}
	TextDanger20             Color = partColor{name: "danger", variant: "20", text: true}
	TextDanger20Invert       Color = partColor{name: "danger", variant: "20", text: true, invert: true}
	TextDanger25             Color = partColor{name: "danger", variant: "25", text: true}
	TextDanger25Invert       Color = partColor{name: "danger", variant: "25", text: true, invert: true}
	TextDanger30             Color = partColor{name: "danger", variant: "30", text: true}
	TextDanger30Invert       Color = partColor{name: "danger", variant: "30", text: true, invert: true}
	TextDanger35             Color = partColor{name: "danger", variant: "35", text: true}
	TextDanger35Invert       Color = partColor{name: "danger", variant: "35", text: true, invert: true}
	TextDanger40             Color = partColor{name: "danger", variant: "40", text: true}
	TextDanger40Invert       Color = partColor{name: "danger", variant: "40", text: true, invert: true}
	TextDanger45             Color = partColor{name: "danger", variant: "45", text: true}
	TextDanger45Invert       Color = partColor{name: "danger", variant: "45", text: true, invert: true}
	TextDanger50             Color = partColor{name: "danger", variant: "50", text: true}
	TextDanger50Invert       Color = partColor{name: "danger", variant: "50", text: true, invert: true}
	TextDanger55             Color = partColor{name: "danger", variant: "55", text: true}
	TextDanger55Invert       Color = partColor{name: "danger", variant: "55", text: true, invert: true}
	TextDanger60             Color = partColor{name: "danger", variant: "60", text: true}
	TextDanger60Invert       Color = partColor{name: "danger", variant: "60", text: true, invert: true}
	TextDanger65             Color = partColor{name: "danger", variant: "65", text: true}
	TextDanger65Invert       Color = partColor{name: "danger", variant: "65", text: true, invert: true}
	TextDanger70             Color = partColor{name: "danger", variant: "70", text: true}
	TextDanger70Invert       Color = partColor{name: "danger", variant: "70", text: true, invert: true}
	TextDanger75             Color = partColor{name: "danger", variant: "75", text: true}
	TextDanger75Invert       Color = partColor{name: "danger", variant: "75", text: true, invert: true}
	TextDanger80             Color = partColor{name: "danger", variant: "80", text: true}
	TextDanger80Invert       Color = partColor{name: "danger", variant: "80", text: true, invert: true}
	TextDanger85             Color = partColor{name: "danger", variant: "85", text: true}
	TextDanger85Invert       Color = partColor{name: "danger", variant: "85", text: true, invert: true}
	TextDanger90             Color = partColor{name: "danger", variant: "90", text: true}
	TextDanger90Invert       Color = partColor{name: "danger", variant: "90", text: true, invert: true}
	TextDanger95             Color = partColor{name: "danger", variant: "95", text: true}
	TextDanger95Invert       Color = partColor{name: "danger", variant: "95", text: true, invert: true}
	TextDanger100            Color = partColor{name: "danger", variant: "100", text: true}
	TextDanger100Invert      Color = partColor{name: "danger", variant: "100", text: true, invert: true}
	TextDangerLight          Color = partColor{name: "danger", text: true, variant: colorLight}
	TextDangerLightInvert    Color = partColor{name: "danger", text: true, variant: colorLight, invert: true}
	TextDangerDark           Color = partColor{name: "danger", text: true, variant: colorDark}
	TextDangerDarkInvert     Color = partColor{name: "danger", text: true, variant: colorDark, invert: true}
	TextDangerSoft           Color = partColor{name: "danger", text: true, variant: colorSoft}
	TextDangerSoftInvert     Color = partColor{name: "danger", text: true, variant: colorSoft, invert: true}
	TextDangerBold           Color = partColor{name: "danger", text: true, variant: colorBold}
	TextDangerBoldInvert     Color = partColor{name: "danger", text: true, variant: colorBold, invert: true}
	TextDangerOnScheme       Color = partColor{name: "danger", text: true, variant: colorOnScheme}
	TextDangerOnSchemeInvert Color = partColor{name: "danger", text: true, variant: colorOnScheme, invert: true}

	TextBlackBis    Color = partColor{name: "black-bis", text: true}
	TextBlackTer    Color = partColor{name: "black-ter", text: true}
	TextGreyDarker  Color = partColor{name: "grey-darker", text: true}
	TextGreyDark    Color = partColor{name: "grey-dark", text: true}
	TextGrey        Color = partColor{name: "grey", text: true}
	TextGreyLight   Color = partColor{name: "grey-light", text: true}
	TextGreyLighter Color = partColor{name: "grey-lighter", text: true}
	TextWhiteTer    Color = partColor{name: "white-ter", text: true}
	TextWhiteBis    Color = partColor{name: "white-bis", text: true}

	BackgroundWhite       Color = partColor{name: "white"}
	BackgroundWhiteInvert Color = partColor{name: "white", invert: true}
	BackgroundBlack       Color = partColor{name: "black"}
	BackgroundBlackInvert Color = partColor{name: "black", invert: true}
	BackgroundLight       Color = partColor{name: "light"}
	BackgroundLightInvert Color = partColor{name: "light", invert: true}
	BackgroundDark        Color = partColor{name: "dark"}
	BackgroundDarkInvert  Color = partColor{name: "dark", invert: true}

	BackgroundText               Color = partColor{name: "text"}
	BackgroundTextInvert         Color = partColor{name: "text", invert: true}
	BackgroundText00             Color = partColor{name: "text", variant: "00"}
	BackgroundText00Invert       Color = partColor{name: "text", variant: "00", invert: true}
	BackgroundText05             Color = partColor{name: "text", variant: "05"}
	BackgroundText05Invert       Color = partColor{name: "text", variant: "05", invert: true}
	BackgroundText10             Color = partColor{name: "text", variant: "10"}
	BackgroundText10Invert       Color = partColor{name: "text", variant: "10", invert: true}
	BackgroundText15             Color = partColor{name: "text", variant: "15"}
	BackgroundText15Invert       Color = partColor{name: "text", variant: "15", invert: true}
	BackgroundText20             Color = partColor{name: "text", variant: "20"}
	BackgroundText20Invert       Color = partColor{name: "text", variant: "20", invert: true}
	BackgroundText25             Color = partColor{name: "text", variant: "25"}
	BackgroundText25Invert       Color = partColor{name: "text", variant: "25", invert: true}
	BackgroundText30             Color = partColor{name: "text", variant: "30"}
	BackgroundText30Invert       Color = partColor{name: "text", variant: "30", invert: true}
	BackgroundText35             Color = partColor{name: "text", variant: "35"}
	BackgroundText35Invert       Color = partColor{name: "text", variant: "35", invert: true}
	BackgroundText40             Color = partColor{name: "text", variant: "40"}
	BackgroundText40Invert       Color = partColor{name: "text", variant: "40", invert: true}
	BackgroundText45             Color = partColor{name: "text", variant: "45"}
	BackgroundText45Invert       Color = partColor{name: "text", variant: "45", invert: true}
	BackgroundText50             Color = partColor{name: "text", variant: "50"}
	BackgroundText50Invert       Color = partColor{name: "text", variant: "50", invert: true}
	BackgroundText55             Color = partColor{name: "text", variant: "55"}
	BackgroundText55Invert       Color = partColor{name: "text", variant: "55", invert: true}
	BackgroundText60             Color = partColor{name: "text", variant: "60"}
	BackgroundText60Invert       Color = partColor{name: "text", variant: "60", invert: true}
	BackgroundText65             Color = partColor{name: "text", variant: "65"}
	BackgroundText65Invert       Color = partColor{name: "text", variant: "65", invert: true}
	BackgroundText70             Color = partColor{name: "text", variant: "70"}
	BackgroundText70Invert       Color = partColor{name: "text", variant: "70", invert: true}
	BackgroundText75             Color = partColor{name: "text", variant: "75"}
	BackgroundText75Invert       Color = partColor{name: "text", variant: "75", invert: true}
	BackgroundText80             Color = partColor{name: "text", variant: "80"}
	BackgroundText80Invert       Color = partColor{name: "text", variant: "80", invert: true}
	BackgroundText85             Color = partColor{name: "text", variant: "85"}
	BackgroundText85Invert       Color = partColor{name: "text", variant: "85", invert: true}
	BackgroundText90             Color = partColor{name: "text", variant: "90"}
	BackgroundText90Invert       Color = partColor{name: "text", variant: "90", invert: true}
	BackgroundText95             Color = partColor{name: "text", variant: "95"}
	BackgroundText95Invert       Color = partColor{name: "text", variant: "95", invert: true}
	BackgroundText100            Color = partColor{name: "text", variant: "100"}
	BackgroundText100Invert      Color = partColor{name: "text", variant: "100", invert: true}
	BackgroundTextLight          Color = partColor{name: "text", text: true, variant: colorLight}
	BackgroundTextLightInvert    Color = partColor{name: "text", text: true, variant: colorLight, invert: true}
	BackgroundTextDark           Color = partColor{name: "text", text: true, variant: colorDark}
	BackgroundTextDarkInvert     Color = partColor{name: "text", text: true, variant: colorDark, invert: true}
	BackgroundTextSoft           Color = partColor{name: "text", text: true, variant: colorSoft}
	BackgroundTextSoftInvert     Color = partColor{name: "text", text: true, variant: colorSoft, invert: true}
	BackgroundTextBold           Color = partColor{name: "text", text: true, variant: colorBold}
	BackgroundTextBoldInvert     Color = partColor{name: "text", text: true, variant: colorBold, invert: true}
	BackgroundTextOnScheme       Color = partColor{name: "text", text: true, variant: colorOnScheme}
	BackgroundTextOnSchemeInvert Color = partColor{name: "text", text: true, variant: colorOnScheme, invert: true}

	BackgroundPrimary               Color = partColor{name: "primary"}
	BackgroundPrimaryInvert         Color = partColor{name: "primary", invert: true}
	BackgroundPrimary00             Color = partColor{name: "primary", variant: "00"}
	BackgroundPrimary00Invert       Color = partColor{name: "primary", variant: "00", invert: true}
	BackgroundPrimary05             Color = partColor{name: "primary", variant: "05"}
	BackgroundPrimary05Invert       Color = partColor{name: "primary", variant: "05", invert: true}
	BackgroundPrimary10             Color = partColor{name: "primary", variant: "10"}
	BackgroundPrimary10Invert       Color = partColor{name: "primary", variant: "10", invert: true}
	BackgroundPrimary15             Color = partColor{name: "primary", variant: "15"}
	BackgroundPrimary15Invert       Color = partColor{name: "primary", variant: "15", invert: true}
	BackgroundPrimary20             Color = partColor{name: "primary", variant: "20"}
	BackgroundPrimary20Invert       Color = partColor{name: "primary", variant: "20", invert: true}
	BackgroundPrimary25             Color = partColor{name: "primary", variant: "25"}
	BackgroundPrimary25Invert       Color = partColor{name: "primary", variant: "25", invert: true}
	BackgroundPrimary30             Color = partColor{name: "primary", variant: "30"}
	BackgroundPrimary30Invert       Color = partColor{name: "primary", variant: "30", invert: true}
	BackgroundPrimary35             Color = partColor{name: "primary", variant: "35"}
	BackgroundPrimary35Invert       Color = partColor{name: "primary", variant: "35", invert: true}
	BackgroundPrimary40             Color = partColor{name: "primary", variant: "40"}
	BackgroundPrimary40Invert       Color = partColor{name: "primary", variant: "40", invert: true}
	BackgroundPrimary45             Color = partColor{name: "primary", variant: "45"}
	BackgroundPrimary45Invert       Color = partColor{name: "primary", variant: "45", invert: true}
	BackgroundPrimary50             Color = partColor{name: "primary", variant: "50"}
	BackgroundPrimary50Invert       Color = partColor{name: "primary", variant: "50", invert: true}
	BackgroundPrimary55             Color = partColor{name: "primary", variant: "55"}
	BackgroundPrimary55Invert       Color = partColor{name: "primary", variant: "55", invert: true}
	BackgroundPrimary60             Color = partColor{name: "primary", variant: "60"}
	BackgroundPrimary60Invert       Color = partColor{name: "primary", variant: "60", invert: true}
	BackgroundPrimary65             Color = partColor{name: "primary", variant: "65"}
	BackgroundPrimary65Invert       Color = partColor{name: "primary", variant: "65", invert: true}
	BackgroundPrimary70             Color = partColor{name: "primary", variant: "70"}
	BackgroundPrimary70Invert       Color = partColor{name: "primary", variant: "70", invert: true}
	BackgroundPrimary75             Color = partColor{name: "primary", variant: "75"}
	BackgroundPrimary75Invert       Color = partColor{name: "primary", variant: "75", invert: true}
	BackgroundPrimary80             Color = partColor{name: "primary", variant: "80"}
	BackgroundPrimary80Invert       Color = partColor{name: "primary", variant: "80", invert: true}
	BackgroundPrimary85             Color = partColor{name: "primary", variant: "85"}
	BackgroundPrimary85Invert       Color = partColor{name: "primary", variant: "85", invert: true}
	BackgroundPrimary90             Color = partColor{name: "primary", variant: "90"}
	BackgroundPrimary90Invert       Color = partColor{name: "primary", variant: "90", invert: true}
	BackgroundPrimary95             Color = partColor{name: "primary", variant: "95"}
	BackgroundPrimary95Invert       Color = partColor{name: "primary", variant: "95", invert: true}
	BackgroundPrimary100            Color = partColor{name: "primary", variant: "100"}
	BackgroundPrimary100Invert      Color = partColor{name: "primary", variant: "100", invert: true}
	BackgroundPrimaryLight          Color = partColor{name: "primary", variant: colorLight}
	BackgroundPrimaryLightInvert    Color = partColor{name: "primary", variant: colorLight, invert: true}
	BackgroundPrimaryDark           Color = partColor{name: "primary", variant: colorDark}
	BackgroundPrimaryDarkInvert     Color = partColor{name: "primary", variant: colorDark, invert: true}
	BackgroundPrimarySoft           Color = partColor{name: "primary", variant: colorSoft}
	BackgroundPrimarySoftInvert     Color = partColor{name: "primary", variant: colorSoft, invert: true}
	BackgroundPrimaryBold           Color = partColor{name: "primary", variant: colorBold}
	BackgroundPrimaryBoldInvert     Color = partColor{name: "primary", variant: colorBold, invert: true}
	BackgroundPrimaryOnScheme       Color = partColor{name: "primary", variant: colorOnScheme}
	BackgroundPrimaryOnSchemeInvert Color = partColor{name: "primary", variant: colorOnScheme, invert: true}

	BackgroundLink               Color = partColor{name: "link"}
	BackgroundLinkInvert         Color = partColor{name: "link", invert: true}
	BackgroundLink00             Color = partColor{name: "link", variant: "00"}
	BackgroundLink00Invert       Color = partColor{name: "link", variant: "00", invert: true}
	BackgroundLink05             Color = partColor{name: "link", variant: "05"}
	BackgroundLink05Invert       Color = partColor{name: "link", variant: "05", invert: true}
	BackgroundLink10             Color = partColor{name: "link", variant: "10"}
	BackgroundLink10Invert       Color = partColor{name: "link", variant: "10", invert: true}
	BackgroundLink15             Color = partColor{name: "link", variant: "15"}
	BackgroundLink15Invert       Color = partColor{name: "link", variant: "15", invert: true}
	BackgroundLink20             Color = partColor{name: "link", variant: "20"}
	BackgroundLink20Invert       Color = partColor{name: "link", variant: "20", invert: true}
	BackgroundLink25             Color = partColor{name: "link", variant: "25"}
	BackgroundLink25Invert       Color = partColor{name: "link", variant: "25", invert: true}
	BackgroundLink30             Color = partColor{name: "link", variant: "30"}
	BackgroundLink30Invert       Color = partColor{name: "link", variant: "30", invert: true}
	BackgroundLink35             Color = partColor{name: "link", variant: "35"}
	BackgroundLink35Invert       Color = partColor{name: "link", variant: "35", invert: true}
	BackgroundLink40             Color = partColor{name: "link", variant: "40"}
	BackgroundLink40Invert       Color = partColor{name: "link", variant: "40", invert: true}
	BackgroundLink45             Color = partColor{name: "link", variant: "45"}
	BackgroundLink45Invert       Color = partColor{name: "link", variant: "45", invert: true}
	BackgroundLink50             Color = partColor{name: "link", variant: "50"}
	BackgroundLink50Invert       Color = partColor{name: "link", variant: "50", invert: true}
	BackgroundLink55             Color = partColor{name: "link", variant: "55"}
	BackgroundLink55Invert       Color = partColor{name: "link", variant: "55", invert: true}
	BackgroundLink60             Color = partColor{name: "link", variant: "60"}
	BackgroundLink60Invert       Color = partColor{name: "link", variant: "60", invert: true}
	BackgroundLink65             Color = partColor{name: "link", variant: "65"}
	BackgroundLink65Invert       Color = partColor{name: "link", variant: "65", invert: true}
	BackgroundLink70             Color = partColor{name: "link", variant: "70"}
	BackgroundLink70Invert       Color = partColor{name: "link", variant: "70", invert: true}
	BackgroundLink75             Color = partColor{name: "link", variant: "75"}
	BackgroundLink75Invert       Color = partColor{name: "link", variant: "75", invert: true}
	BackgroundLink80             Color = partColor{name: "link", variant: "80"}
	BackgroundLink80Invert       Color = partColor{name: "link", variant: "80", invert: true}
	BackgroundLink85             Color = partColor{name: "link", variant: "85"}
	BackgroundLink85Invert       Color = partColor{name: "link", variant: "85", invert: true}
	BackgroundLink90             Color = partColor{name: "link", variant: "90"}
	BackgroundLink90Invert       Color = partColor{name: "link", variant: "90", invert: true}
	BackgroundLink95             Color = partColor{name: "link", variant: "95"}
	BackgroundLink95Invert       Color = partColor{name: "link", variant: "95", invert: true}
	BackgroundLink100            Color = partColor{name: "link", variant: "100"}
	BackgroundLink100Invert      Color = partColor{name: "link", variant: "100", invert: true}
	BackgroundLinkLight          Color = partColor{name: "link", variant: colorLight}
	BackgroundLinkLightInvert    Color = partColor{name: "link", variant: colorLight, invert: true}
	BackgroundLinkDark           Color = partColor{name: "link", variant: colorDark}
	BackgroundLinkDarkInvert     Color = partColor{name: "link", variant: colorDark, invert: true}
	BackgroundLinkSoft           Color = partColor{name: "link", variant: colorSoft}
	BackgroundLinkSoftInvert     Color = partColor{name: "link", variant: colorSoft, invert: true}
	BackgroundLinkBold           Color = partColor{name: "link", variant: colorBold}
	BackgroundLinkBoldInvert     Color = partColor{name: "link", variant: colorBold, invert: true}
	BackgroundLinkOnScheme       Color = partColor{name: "link", variant: colorOnScheme}
	BackgroundLinkOnSchemeInvert Color = partColor{name: "link", variant: colorOnScheme, invert: true}

	BackgroundInfo               Color = partColor{name: "info"}
	BackgroundInfoInvert         Color = partColor{name: "info", invert: true}
	BackgroundInfo00             Color = partColor{name: "info", variant: "00"}
	BackgroundInfo00Invert       Color = partColor{name: "info", variant: "00", invert: true}
	BackgroundInfo05             Color = partColor{name: "info", variant: "05"}
	BackgroundInfo05Invert       Color = partColor{name: "info", variant: "05", invert: true}
	BackgroundInfo10             Color = partColor{name: "info", variant: "10"}
	BackgroundInfo10Invert       Color = partColor{name: "info", variant: "10", invert: true}
	BackgroundInfo15             Color = partColor{name: "info", variant: "15"}
	BackgroundInfo15Invert       Color = partColor{name: "info", variant: "15", invert: true}
	BackgroundInfo20             Color = partColor{name: "info", variant: "20"}
	BackgroundInfo20Invert       Color = partColor{name: "info", variant: "20", invert: true}
	BackgroundInfo25             Color = partColor{name: "info", variant: "25"}
	BackgroundInfo25Invert       Color = partColor{name: "info", variant: "25", invert: true}
	BackgroundInfo30             Color = partColor{name: "info", variant: "30"}
	BackgroundInfo30Invert       Color = partColor{name: "info", variant: "30", invert: true}
	BackgroundInfo35             Color = partColor{name: "info", variant: "35"}
	BackgroundInfo35Invert       Color = partColor{name: "info", variant: "35", invert: true}
	BackgroundInfo40             Color = partColor{name: "info", variant: "40"}
	BackgroundInfo40Invert       Color = partColor{name: "info", variant: "40", invert: true}
	BackgroundInfo45             Color = partColor{name: "info", variant: "45"}
	BackgroundInfo45Invert       Color = partColor{name: "info", variant: "45", invert: true}
	BackgroundInfo50             Color = partColor{name: "info", variant: "50"}
	BackgroundInfo50Invert       Color = partColor{name: "info", variant: "50", invert: true}
	BackgroundInfo55             Color = partColor{name: "info", variant: "55"}
	BackgroundInfo55Invert       Color = partColor{name: "info", variant: "55", invert: true}
	BackgroundInfo60             Color = partColor{name: "info", variant: "60"}
	BackgroundInfo60Invert       Color = partColor{name: "info", variant: "60", invert: true}
	BackgroundInfo65             Color = partColor{name: "info", variant: "65"}
	BackgroundInfo65Invert       Color = partColor{name: "info", variant: "65", invert: true}
	BackgroundInfo70             Color = partColor{name: "info", variant: "70"}
	BackgroundInfo70Invert       Color = partColor{name: "info", variant: "70", invert: true}
	BackgroundInfo75             Color = partColor{name: "info", variant: "75"}
	BackgroundInfo75Invert       Color = partColor{name: "info", variant: "75", invert: true}
	BackgroundInfo80             Color = partColor{name: "info", variant: "80"}
	BackgroundInfo80Invert       Color = partColor{name: "info", variant: "80", invert: true}
	BackgroundInfo85             Color = partColor{name: "info", variant: "85"}
	BackgroundInfo85Invert       Color = partColor{name: "info", variant: "85", invert: true}
	BackgroundInfo90             Color = partColor{name: "info", variant: "90"}
	BackgroundInfo90Invert       Color = partColor{name: "info", variant: "90", invert: true}
	BackgroundInfo95             Color = partColor{name: "info", variant: "95"}
	BackgroundInfo95Invert       Color = partColor{name: "info", variant: "95", invert: true}
	BackgroundInfo100            Color = partColor{name: "info", variant: "100"}
	BackgroundInfo100Invert      Color = partColor{name: "info", variant: "100", invert: true}
	BackgroundInfoLight          Color = partColor{name: "info", variant: colorLight}
	BackgroundInfoLightInvert    Color = partColor{name: "info", variant: colorLight, invert: true}
	BackgroundInfoDark           Color = partColor{name: "info", variant: colorDark}
	BackgroundInfoDarkInvert     Color = partColor{name: "info", variant: colorDark, invert: true}
	BackgroundInfoSoft           Color = partColor{name: "info", variant: colorSoft}
	BackgroundInfoSoftInvert     Color = partColor{name: "info", variant: colorSoft, invert: true}
	BackgroundInfoBold           Color = partColor{name: "info", variant: colorBold}
	BackgroundInfoBoldInvert     Color = partColor{name: "info", variant: colorBold, invert: true}
	BackgroundInfoOnScheme       Color = partColor{name: "info", variant: colorOnScheme}
	BackgroundInfoOnSchemeInvert Color = partColor{name: "info", variant: colorOnScheme, invert: true}

	BackgroundSuccess               Color = partColor{name: "success"}
	BackgroundSuccessInvert         Color = partColor{name: "success", invert: true}
	BackgroundSuccess00             Color = partColor{name: "success", variant: "00"}
	BackgroundSuccess00Invert       Color = partColor{name: "success", variant: "00", invert: true}
	BackgroundSuccess05             Color = partColor{name: "success", variant: "05"}
	BackgroundSuccess05Invert       Color = partColor{name: "success", variant: "05", invert: true}
	BackgroundSuccess10             Color = partColor{name: "success", variant: "10"}
	BackgroundSuccess10Invert       Color = partColor{name: "success", variant: "10", invert: true}
	BackgroundSuccess15             Color = partColor{name: "success", variant: "15"}
	BackgroundSuccess15Invert       Color = partColor{name: "success", variant: "15", invert: true}
	BackgroundSuccess20             Color = partColor{name: "success", variant: "20"}
	BackgroundSuccess20Invert       Color = partColor{name: "success", variant: "20", invert: true}
	BackgroundSuccess25             Color = partColor{name: "success", variant: "25"}
	BackgroundSuccess25Invert       Color = partColor{name: "success", variant: "25", invert: true}
	BackgroundSuccess30             Color = partColor{name: "success", variant: "30"}
	BackgroundSuccess30Invert       Color = partColor{name: "success", variant: "30", invert: true}
	BackgroundSuccess35             Color = partColor{name: "success", variant: "35"}
	BackgroundSuccess35Invert       Color = partColor{name: "success", variant: "35", invert: true}
	BackgroundSuccess40             Color = partColor{name: "success", variant: "40"}
	BackgroundSuccess40Invert       Color = partColor{name: "success", variant: "40", invert: true}
	BackgroundSuccess45             Color = partColor{name: "success", variant: "45"}
	BackgroundSuccess45Invert       Color = partColor{name: "success", variant: "45", invert: true}
	BackgroundSuccess50             Color = partColor{name: "success", variant: "50"}
	BackgroundSuccess50Invert       Color = partColor{name: "success", variant: "50", invert: true}
	BackgroundSuccess55             Color = partColor{name: "success", variant: "55"}
	BackgroundSuccess55Invert       Color = partColor{name: "success", variant: "55", invert: true}
	BackgroundSuccess60             Color = partColor{name: "success", variant: "60"}
	BackgroundSuccess60Invert       Color = partColor{name: "success", variant: "60", invert: true}
	BackgroundSuccess65             Color = partColor{name: "success", variant: "65"}
	BackgroundSuccess65Invert       Color = partColor{name: "success", variant: "65", invert: true}
	BackgroundSuccess70             Color = partColor{name: "success", variant: "70"}
	BackgroundSuccess70Invert       Color = partColor{name: "success", variant: "70", invert: true}
	BackgroundSuccess75             Color = partColor{name: "success", variant: "75"}
	BackgroundSuccess75Invert       Color = partColor{name: "success", variant: "75", invert: true}
	BackgroundSuccess80             Color = partColor{name: "success", variant: "80"}
	BackgroundSuccess80Invert       Color = partColor{name: "success", variant: "80", invert: true}
	BackgroundSuccess85             Color = partColor{name: "success", variant: "85"}
	BackgroundSuccess85Invert       Color = partColor{name: "success", variant: "85", invert: true}
	BackgroundSuccess90             Color = partColor{name: "success", variant: "90"}
	BackgroundSuccess90Invert       Color = partColor{name: "success", variant: "90", invert: true}
	BackgroundSuccess95             Color = partColor{name: "success", variant: "95"}
	BackgroundSuccess95Invert       Color = partColor{name: "success", variant: "95", invert: true}
	BackgroundSuccess100            Color = partColor{name: "success", variant: "100"}
	BackgroundSuccess100Invert      Color = partColor{name: "success", variant: "100", invert: true}
	BackgroundSuccessLight          Color = partColor{name: "success", variant: colorLight}
	BackgroundSuccessLightInvert    Color = partColor{name: "success", variant: colorLight, invert: true}
	BackgroundSuccessDark           Color = partColor{name: "success", variant: colorDark}
	BackgroundSuccessDarkInvert     Color = partColor{name: "success", variant: colorDark, invert: true}
	BackgroundSuccessSoft           Color = partColor{name: "success", variant: colorSoft}
	BackgroundSuccessSoftInvert     Color = partColor{name: "success", variant: colorSoft, invert: true}
	BackgroundSuccessBold           Color = partColor{name: "success", variant: colorBold}
	BackgroundSuccessBoldInvert     Color = partColor{name: "success", variant: colorBold, invert: true}
	BackgroundSuccessOnScheme       Color = partColor{name: "success", variant: colorOnScheme}
	BackgroundSuccessOnSchemeInvert Color = partColor{name: "success", variant: colorOnScheme, invert: true}

	BackgroundWarning               Color = partColor{name: "warning"}
	BackgroundWarningInvert         Color = partColor{name: "warning", invert: true}
	BackgroundWarning00             Color = partColor{name: "warning", variant: "00"}
	BackgroundWarning00Invert       Color = partColor{name: "warning", variant: "00", invert: true}
	BackgroundWarning05             Color = partColor{name: "warning", variant: "05"}
	BackgroundWarning05Invert       Color = partColor{name: "warning", variant: "05", invert: true}
	BackgroundWarning10             Color = partColor{name: "warning", variant: "10"}
	BackgroundWarning10Invert       Color = partColor{name: "warning", variant: "10", invert: true}
	BackgroundWarning15             Color = partColor{name: "warning", variant: "15"}
	BackgroundWarning15Invert       Color = partColor{name: "warning", variant: "15", invert: true}
	BackgroundWarning20             Color = partColor{name: "warning", variant: "20"}
	BackgroundWarning20Invert       Color = partColor{name: "warning", variant: "20", invert: true}
	BackgroundWarning25             Color = partColor{name: "warning", variant: "25"}
	BackgroundWarning25Invert       Color = partColor{name: "warning", variant: "25", invert: true}
	BackgroundWarning30             Color = partColor{name: "warning", variant: "30"}
	BackgroundWarning30Invert       Color = partColor{name: "warning", variant: "30", invert: true}
	BackgroundWarning35             Color = partColor{name: "warning", variant: "35"}
	BackgroundWarning35Invert       Color = partColor{name: "warning", variant: "35", invert: true}
	BackgroundWarning40             Color = partColor{name: "warning", variant: "40"}
	BackgroundWarning40Invert       Color = partColor{name: "warning", variant: "40", invert: true}
	BackgroundWarning45             Color = partColor{name: "warning", variant: "45"}
	BackgroundWarning45Invert       Color = partColor{name: "warning", variant: "45", invert: true}
	BackgroundWarning50             Color = partColor{name: "warning", variant: "50"}
	BackgroundWarning50Invert       Color = partColor{name: "warning", variant: "50", invert: true}
	BackgroundWarning55             Color = partColor{name: "warning", variant: "55"}
	BackgroundWarning55Invert       Color = partColor{name: "warning", variant: "55", invert: true}
	BackgroundWarning60             Color = partColor{name: "warning", variant: "60"}
	BackgroundWarning60Invert       Color = partColor{name: "warning", variant: "60", invert: true}
	BackgroundWarning65             Color = partColor{name: "warning", variant: "65"}
	BackgroundWarning65Invert       Color = partColor{name: "warning", variant: "65", invert: true}
	BackgroundWarning70             Color = partColor{name: "warning", variant: "70"}
	BackgroundWarning70Invert       Color = partColor{name: "warning", variant: "70", invert: true}
	BackgroundWarning75             Color = partColor{name: "warning", variant: "75"}
	BackgroundWarning75Invert       Color = partColor{name: "warning", variant: "75", invert: true}
	BackgroundWarning80             Color = partColor{name: "warning", variant: "80"}
	BackgroundWarning80Invert       Color = partColor{name: "warning", variant: "80", invert: true}
	BackgroundWarning85             Color = partColor{name: "warning", variant: "85"}
	BackgroundWarning85Invert       Color = partColor{name: "warning", variant: "85", invert: true}
	BackgroundWarning90             Color = partColor{name: "warning", variant: "90"}
	BackgroundWarning90Invert       Color = partColor{name: "warning", variant: "90", invert: true}
	BackgroundWarning95             Color = partColor{name: "warning", variant: "95"}
	BackgroundWarning95Invert       Color = partColor{name: "warning", variant: "95", invert: true}
	BackgroundWarning100            Color = partColor{name: "warning", variant: "100"}
	BackgroundWarning100Invert      Color = partColor{name: "warning", variant: "100", invert: true}
	BackgroundWarningLight          Color = partColor{name: "warning", variant: colorLight}
	BackgroundWarningLightInvert    Color = partColor{name: "warning", variant: colorLight, invert: true}
	BackgroundWarningDark           Color = partColor{name: "warning", variant: colorDark}
	BackgroundWarningDarkInvert     Color = partColor{name: "warning", variant: colorDark, invert: true}
	BackgroundWarningSoft           Color = partColor{name: "warning", variant: colorSoft}
	BackgroundWarningSoftInvert     Color = partColor{name: "warning", variant: colorSoft, invert: true}
	BackgroundWarningBold           Color = partColor{name: "warning", variant: colorBold}
	BackgroundWarningBoldInvert     Color = partColor{name: "warning", variant: colorBold, invert: true}
	BackgroundWarningOnScheme       Color = partColor{name: "warning", variant: colorOnScheme}
	BackgroundWarningOnSchemeInvert Color = partColor{name: "warning", variant: colorOnScheme, invert: true}

	BackgroundDanger               Color = partColor{name: "danger"}
	BackgroundDangerInvert         Color = partColor{name: "danger", invert: true}
	BackgroundDanger00             Color = partColor{name: "danger", variant: "00"}
	BackgroundDanger00Invert       Color = partColor{name: "danger", variant: "00", invert: true}
	BackgroundDanger05             Color = partColor{name: "danger", variant: "05"}
	BackgroundDanger05Invert       Color = partColor{name: "danger", variant: "05", invert: true}
	BackgroundDanger10             Color = partColor{name: "danger", variant: "10"}
	BackgroundDanger10Invert       Color = partColor{name: "danger", variant: "10", invert: true}
	BackgroundDanger15             Color = partColor{name: "danger", variant: "15"}
	BackgroundDanger15Invert       Color = partColor{name: "danger", variant: "15", invert: true}
	BackgroundDanger20             Color = partColor{name: "danger", variant: "20"}
	BackgroundDanger20Invert       Color = partColor{name: "danger", variant: "20", invert: true}
	BackgroundDanger25             Color = partColor{name: "danger", variant: "25"}
	BackgroundDanger25Invert       Color = partColor{name: "danger", variant: "25", invert: true}
	BackgroundDanger30             Color = partColor{name: "danger", variant: "30"}
	BackgroundDanger30Invert       Color = partColor{name: "danger", variant: "30", invert: true}
	BackgroundDanger35             Color = partColor{name: "danger", variant: "35"}
	BackgroundDanger35Invert       Color = partColor{name: "danger", variant: "35", invert: true}
	BackgroundDanger40             Color = partColor{name: "danger", variant: "40"}
	BackgroundDanger40Invert       Color = partColor{name: "danger", variant: "40", invert: true}
	BackgroundDanger45             Color = partColor{name: "danger", variant: "45"}
	BackgroundDanger45Invert       Color = partColor{name: "danger", variant: "45", invert: true}
	BackgroundDanger50             Color = partColor{name: "danger", variant: "50"}
	BackgroundDanger50Invert       Color = partColor{name: "danger", variant: "50", invert: true}
	BackgroundDanger55             Color = partColor{name: "danger", variant: "55"}
	BackgroundDanger55Invert       Color = partColor{name: "danger", variant: "55", invert: true}
	BackgroundDanger60             Color = partColor{name: "danger", variant: "60"}
	BackgroundDanger60Invert       Color = partColor{name: "danger", variant: "60", invert: true}
	BackgroundDanger65             Color = partColor{name: "danger", variant: "65"}
	BackgroundDanger65Invert       Color = partColor{name: "danger", variant: "65", invert: true}
	BackgroundDanger70             Color = partColor{name: "danger", variant: "70"}
	BackgroundDanger70Invert       Color = partColor{name: "danger", variant: "70", invert: true}
	BackgroundDanger75             Color = partColor{name: "danger", variant: "75"}
	BackgroundDanger75Invert       Color = partColor{name: "danger", variant: "75", invert: true}
	BackgroundDanger80             Color = partColor{name: "danger", variant: "80"}
	BackgroundDanger80Invert       Color = partColor{name: "danger", variant: "80", invert: true}
	BackgroundDanger85             Color = partColor{name: "danger", variant: "85"}
	BackgroundDanger85Invert       Color = partColor{name: "danger", variant: "85", invert: true}
	BackgroundDanger90             Color = partColor{name: "danger", variant: "90"}
	BackgroundDanger90Invert       Color = partColor{name: "danger", variant: "90", invert: true}
	BackgroundDanger95             Color = partColor{name: "danger", variant: "95"}
	BackgroundDanger95Invert       Color = partColor{name: "danger", variant: "95", invert: true}
	BackgroundDanger100            Color = partColor{name: "danger", variant: "100"}
	BackgroundDanger100Invert      Color = partColor{name: "danger", variant: "100", invert: true}
	BackgroundDangerLight          Color = partColor{name: "danger", variant: colorLight}
	BackgroundDangerLightInvert    Color = partColor{name: "danger", variant: colorLight, invert: true}
	BackgroundDangerDark           Color = partColor{name: "danger", variant: colorDark}
	BackgroundDangerDarkInvert     Color = partColor{name: "danger", variant: colorDark, invert: true}
	BackgroundDangerSoft           Color = partColor{name: "danger", variant: colorSoft}
	BackgroundDangerSoftInvert     Color = partColor{name: "danger", variant: colorSoft, invert: true}
	BackgroundDangerBold           Color = partColor{name: "danger", variant: colorBold}
	BackgroundDangerBoldInvert     Color = partColor{name: "danger", variant: colorBold, invert: true}
	BackgroundDangerOnScheme       Color = partColor{name: "danger", variant: colorOnScheme}
	BackgroundDangerOnSchemeInvert Color = partColor{name: "danger", variant: colorOnScheme, invert: true}

	BackgroundBlackBis    Color = partColor{name: "black-bis"}
	BackgroundBlackTer    Color = partColor{name: "black-ter"}
	BackgroundGreyDarker  Color = partColor{name: "grey-darker"}
	BackgroundGreyDark    Color = partColor{name: "grey-dark"}
	BackgroundGrey        Color = partColor{name: "grey"}
	BackgroundGreyLight   Color = partColor{name: "grey-light"}
	BackgroundGreyLighter Color = partColor{name: "grey-lighter"}
	BackgroundWhiteTer    Color = partColor{name: "white-ter"}
	BackgroundWhiteBis    Color = partColor{name: "white-bis"}
)
