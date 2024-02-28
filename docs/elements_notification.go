package docs

import (
	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var notification = c.NewPage(
	"Notification", "Notification", "/notification",
	"https://bulma.io/documentation/elements/notification/",

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
).Section(
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
).Section(
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
