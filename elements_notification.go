package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var notification = c.NewPage(
	"Notification", "Notification", "/notification",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Notification"), " constructor creates a simple color block meant to draw the attention to the user about something. As such, it an be used as a pinned notification in the corner of the viewport.",
		),
		c.Modifiers(
			c.Row("b.White", "Set color to white"),
			c.Row("b.Black", "Set color to black"),
			c.Row("b.Light", "Set color to light"),
			c.Row("b.Dark", "Set color to dark"),
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
			c.Row("b.PrimaryLight", "Set color to primary light"),
			c.Row("b.LinkLight", "Set color to link light"),
			c.Row("b.InfoLight", "Set color to info light"),
			c.Row("b.SuccessLight", "Set color to success light"),
			c.Row("b.WarningLight", "Set color to warning light"),
			c.Row("b.DangerLight", "Set color to danger light"),
		),
		e.P(
			"When you provide ", e.Code("b.Delete(...any)"), " as a child, you may add ", e.Code("e.OnClick(b.JSRemoveThisNotification)"), " to its children in order to close the notification when the user clicks on it.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/notification/",

	c.Example(
		`b.Notification(
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		)`,
		b.Notification(
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/notification/#colors",
	c.Example(
		`b.Notification(
	b.Primary,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Primary,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),

	c.Example(
		`b.Notification(
	b.Link,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Link,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Info,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Info,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Success,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Success,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Warning,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Warning,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Danger,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Danger,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
).Subsection(
	"Light colors",
	"https://bulma.io/documentation/elements/notification/#light-colors",
	c.Example(
		`b.Notification(
	b.PrimaryLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.PrimaryLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),

	c.Example(
		`b.Notification(
	b.LinkLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.LinkLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.InfoLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.InfoLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.SuccessLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.SuccessLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.WarningLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.WarningLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.DangerLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.DangerLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", e.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur.",
		),
	),
)
