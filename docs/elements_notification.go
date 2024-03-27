package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var notification = c.NewPage(
	"Notification", "Notification", "/notification",
	"",

	b.Content(
		b.Style("column-count", "2"),
		el.P(
			b.Style("column-span", "all"),
			"The ", el.Code("b.Notification"), " constructor returns a simple color block meant to draw the attention to the user about something. As such, it an be used as a pinned notification in the corner of the viewport. It accepts the following values additionally to the standard set of children types:",
		),
		b.DList(
			el.Code("b.White"),
			"Set the notification color to white",
			el.Code("b.Black"),
			"Set the notification color to black",
			el.Code("b.Light"),
			"Set the notification color to light",
			el.Code("b.Dark"),
			"Set the notification color to dark",
			el.Code("b.Primary"),
			"Set the notification color to primary",
			el.Code("b.Link"),
			"Set the notification color to link",
			el.Code("b.Info"),
			"Set the notification color to info",
			el.Code("b.Success"),
			"Set the notification color to success",
			el.Code("b.Warning"),
			"Set the notification color to warning",
			el.Code("b.Danger"),
			"Set the notification color to danger",
			el.Code("b.PrimaryLight"),
			"Set the notification color to primary light",
			el.Code("b.LinkLight"),
			"Set the notification color to link light",
			el.Code("b.InfoLight"),
			"Set the notification color to info light",
			el.Code("b.SuccessLight"),
			"Set the notification color to success light",
			el.Code("b.WarningLight"),
			"Set the notification color to warning light",
			el.Code("b.DangerLight"),
			"Set the notification color to danger light",
		),
		el.P(
			b.Style("column-span", "all"),
			"You may add ", el.Code("b.OnClick(b.JSRemoveThisNotification)"), " to the ", el.Code("b.Delete"), " inner element in order to close the notification when the user clicks on that element.",
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/notification/",

	c.Example(
		`b.Notification(
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		)`,
		b.Notification(
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/notification/#colors",
	c.Example(
		`b.Notification(
	b.Primary,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Primary,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),

	c.Example(
		`b.Notification(
	b.Link,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Link,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Info,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Info,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Success,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Success,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Warning,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Warning,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.Danger,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.Danger,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
).Subsection(
	"Light colors",
	"https://bulma.io/documentation/elements/notification/#light-colors",
	c.Example(
		`b.Notification(
	b.PrimaryLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.PrimaryLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),

	c.Example(
		`b.Notification(
	b.LinkLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.LinkLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.InfoLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.InfoLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.SuccessLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.SuccessLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.WarningLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.WarningLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
	c.Example(
		`b.Notification(
	b.DangerLight,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
)`,
		b.Notification(
			b.DangerLight,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. ", el.Strong("Pellentesque risus mi"), ", tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur.",
		),
	),
)
