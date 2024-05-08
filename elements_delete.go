package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var delete = c.NewPage(
	"Delete", "Delete", "/delete",
	"",
	b.Content(
		e.P(
			"The ", e.Code("b.Delete"), " constructor creates a delete cross.",
		),
		c.Modifiers(
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Normal", "Set size to normal"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/delete/",

	c.Example("b.Delete()", b.Delete()),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/elements/delete/#sizes",
	c.Example(
		`b.Delete(b.Small),
b.Delete(),
b.Delete(b.Medium),
b.Delete(b.Large)`,
		b.Delete(b.Small, b.MarginRight(b.Spacing2)),
		b.Delete(b.MarginRight(b.Spacing2)),
		b.Delete(b.Medium, b.MarginRight(b.Spacing2)),
		b.Delete(b.Large, b.MarginRight(b.Spacing2)),
	),
).Subsection(
	"Combinations",
	"https://bulma.io/documentation/elements/delete/#combinations",
	c.Example(
		`b.Block(
	b.Tag(
		b.Success,
		"Hello World",
		b.Delete(b.Small),
	),
),
b.Notification(
	b.Danger,
	b.Delete(),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor sit amet, consectetur adipiscing elit",
),
b.Message(
	b.Info,
	b.MessageHeader(
		"Info",
		b.Delete(),
	),
	b.MessageBody(
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque risus mi, tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum felis venenatis efficitur. Aenean ac eleifend lacus, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
	),
)`,
		b.Block(
			b.Tag(
				b.Success,
				"Hello World",
				b.Delete(b.Small),
			),
		),
		b.Notification(
			b.Danger,
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor sit amet, consectetur adipiscing elit",
		),
		b.Message(
			b.Info,
			b.MessageTitle("Info"),
			b.Delete(),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque risus mi, tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum felis venenatis efficitur. Aenean ac eleifend lacus, in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
)
