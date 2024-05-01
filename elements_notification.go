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
			"The ", e.Code("b.Notification"), " constructor creates a simple color block meant to draw the attention to the user about something. As such, it an be used as a pinned notification in the corner of the viewport. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.White"),
			"Set the notification color to white",

			e.Code("b.Black"),
			"Set the notification color to black",

			e.Code("b.Light"),
			"Set the notification color to light",

			e.Code("b.Dark"),
			"Set the notification color to dark",

			e.Code("b.Primary"),
			"Set the notification color to primary",

			e.Code("b.Link"),
			"Set the notification color to link",

			e.Code("b.Info"),
			"Set the notification color to info",

			e.Code("b.Success"),
			"Set the notification color to success",

			e.Code("b.Warning"),
			"Set the notification color to warning",

			e.Code("b.Danger"),
			"Set the notification color to danger",

			e.Code("b.PrimaryLight"),
			"Set the notification color to primary light",

			e.Code("b.LinkLight"),
			"Set the notification color to link light",

			e.Code("b.InfoLight"),
			"Set the notification color to info light",

			e.Code("b.SuccessLight"),
			"Set the notification color to success light",

			e.Code("b.WarningLight"),
			"Set the notification color to warning light",

			e.Code("b.DangerLight"),
			"Set the notification color to danger light",
		),
		e.P(
			"When you provide ", e.Code("b.Delete(...)"), " as a child, you may add ", e.Code("e.OnClick(b.JSRemoveThisNotification)"), " to its children in order to close the notification when the user clicks on it.",
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
