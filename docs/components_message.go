package docs

import (
	e "github.com/willoma/gomplements"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
)

var message = c.NewPage(
	"Message", "Message", "/message",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Message"), " constructor creates a message. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.Dark"),
			"Set message color to dark",

			e.Code("b.Primary"),
			"Set message color to primary",

			e.Code("b.Link"),
			"Set message color to link",

			e.Code("b.Info"),
			"Set message color to info",

			e.Code("b.Success"),
			"Set message color to success",

			e.Code("b.Warning"),
			"Set message color to warning",

			e.Code("b.Danger"),
			"Set message color to danger",
		),
		e.P("A ", e.Code("b.MessageBody(...)"), " child is needed to define the message body, and an optional ", e.Code("b.MessageHeader(...)"), ` child may be provided to define the header. With no header, the message is displayed with the "message body only" style.`),
	),
).Section(
	"Easy helper", "",

	e.P(
		"The ", e.Code("b.Message"), " constructor creates a message. It accepts the same colors definitions as the ", e.Code("b.Message"), " constructor as well as the following values additionally to the standard set of children types:",
	),
	b.DList(
		e.Code("b.Inner(any)"),
		[]any{"forcibly apply the child to the body e.Element"},

		e.Code("b.Outer(any)"),
		[]any{"forcibly apply the child to the message e.Element"},

		e.Code("b.MessageTitle"),
		"Include a header with the provided title",

		e.Code("b.MessageDeleteOnClick"),
		"Include a delete button in the header, with the provided script as an onclick event",

		"A class or style",
		"Apply to the message e.Element",

		[]any{e.Code("gomponents.Node"), " of type ", e.Code("gomponents.AttributeType")},
		"Apply the attribute to the message e.Element",

		[]any{"Other ", e.Code("gomponents.Node")},
		"Add this e.Element to the body e.Element",

		"Any other type",
		"Add the child to the body e.Element",
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/message/",
	c.Example(
		`b.Message(
	b.MessageTitle("Hello world"),
	b.Delete(e.AriaLabel("delete")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.MessageTitle("Hello world"),
			b.Delete(e.AriaLabel("delete")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/components/message/#colors",
	c.Example(
		`b.Message(
	b.Dark,
	b.MessageTitle("Dark"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
),`,
		b.Message(
			b.Dark,
			b.MessageTitle("Dark"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Primary,
	b.MessageTitle("Primary"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Primary,
			b.MessageTitle("Primary"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Link,
	b.MessageTitle("Link"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Link,
			b.MessageTitle("Link"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Info,
	b.MessageTitle("Info"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Info,
			b.MessageTitle("Info"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Success,
	b.MessageTitle("Success"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Success,
			b.MessageTitle("Success"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Primary,
	b.MessageTitle("Warning"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Warning,
			b.MessageTitle("Warning"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Danger,
	b.MessageTitle("Danger"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Danger,
			b.MessageTitle("Danger"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Message body only",
	"https://bulma.io/documentation/components/message/#message-body-only",
	c.Example(
		`b.Message(
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Dark,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Dark,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Primary,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Primary,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Link,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Link,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Info,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Info,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Success,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Success,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Warning,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Warning,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Danger,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Danger,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/components/message/#sizes",
	c.Example(
		`b.Message(
	b.Small,
	b.MessageTitle("Small message"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.Small,
			b.MessageTitle("Small message"),
			b.Delete(b.Small, e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus.",
		),
	),
	c.Example(
		`b.Message(
	b.MessageTitle("Normal message"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.MessageTitle("Normal message"),
			b.Delete(e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus."),
	),
	c.Example(
		`b.Message(
	b.Medium,
	b.MessageTitle("Medium message"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.Medium,
			b.MessageTitle("Medium message"),
			b.Delete(b.Medium, e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus."),
	),
	c.Example(
		`b.Message(
	b.Large,
	b.MessageTitle("Large message"),
	b.Delete(e.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.Large,
			b.MessageTitle("Large message"),
			b.Delete(b.Large, e.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", e.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", e.A("felis venenatis"), " efficitur. Aenean ac ", e.Em("eleifend lacus"), " in mollis lectus."),
	),
)
