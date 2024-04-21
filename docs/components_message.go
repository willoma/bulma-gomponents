package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var message = c.NewPage(
	"Message", "Message", "/message",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Message"), " constructor creates a message. The following children have a special meaning:",
		),
		b.DList(
			el.Code("b.Dark"),
			"Set message color to dark",

			el.Code("b.Primary"),
			"Set message color to primary",

			el.Code("b.Link"),
			"Set message color to link",

			el.Code("b.Info"),
			"Set message color to info",

			el.Code("b.Success"),
			"Set message color to success",

			el.Code("b.Warning"),
			"Set message color to warning",

			el.Code("b.Danger"),
			"Set message color to danger",
		),
		el.P("A ", el.Code("b.MessageBody(...)"), " child is needed to define the message body, and an optional ", el.Code("b.MessageHeader(...)"), ` child may be provided to define the header. With no header, the message is displayed with the "message body only" style.`),
	),
).Section(
	"Easy helper", "",

	el.P(
		"The ", el.Code("b.Message"), " constructor creates a message. It accepts the same colors definitions as the ", el.Code("b.Message"), " constructor as well as the following values additionally to the standard set of children types:",
	),
	b.DList(
		el.Code("b.Inner(any)"),
		[]any{"forcibly apply the child to the body element"},

		el.Code("b.Outer(any)"),
		[]any{"forcibly apply the child to the message element"},

		el.Code("b.MessageTitle"),
		"Include a header with the provided title",

		el.Code("b.MessageDeleteOnClick"),
		"Include a delete button in the header, with the provided script as an onclick event",

		"A class or style",
		"Apply to the message element",

		[]any{el.Code("gomponents.Node"), " of type ", el.Code("gomponents.AttributeType")},
		"Apply the attribute to the message element",

		[]any{"Other ", el.Code("gomponents.Node")},
		"Add this element to the body element",

		"Any other type",
		"Add the child to the body element",
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/message/",
	c.Example(
		`b.Message(
	b.MessageTitle("Hello world"),
	b.Delete(html.Aria("label", "delete")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.MessageTitle("Hello world"),
			b.Delete(html.Aria("label", "delete")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/components/message/#colors",
	c.Example(
		`b.Message(
	b.Dark,
	b.MessageTitle("Dark"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
),`,
		b.Message(
			b.Dark,
			b.MessageTitle("Dark"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Primary,
	b.MessageTitle("Primary"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Primary,
			b.MessageTitle("Primary"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Link,
	b.MessageTitle("Link"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Link,
			b.MessageTitle("Link"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Info,
	b.MessageTitle("Info"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Info,
			b.MessageTitle("Info"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Success,
	b.MessageTitle("Success"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Success,
			b.MessageTitle("Success"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Primary,
	b.MessageTitle("Warning"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Warning,
			b.MessageTitle("Warning"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Danger,
	b.MessageTitle("Danger"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Danger,
			b.MessageTitle("Danger"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Message body only",
	"https://bulma.io/documentation/components/message/#message-body-only",
	c.Example(
		`b.Message(
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Dark,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Dark,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Primary,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Primary,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Link,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Link,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Info,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Info,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Success,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Success,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Warning,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Warning,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`b.Message(
	b.Danger,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		b.Message(
			b.Danger,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/components/message/#sizes",
	c.Example(
		`b.Message(
	b.Small,
	b.MessageTitle("Small message"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.Small,
			b.MessageTitle("Small message"),
			b.Delete(b.Small, b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
		),
	),
	c.Example(
		`b.Message(
	b.MessageTitle("Normal message"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.MessageTitle("Normal message"),
			b.Delete(b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus."),
	),
	c.Example(
		`b.Message(
	b.Medium,
	b.MessageTitle("Medium message"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.Medium,
			b.MessageTitle("Medium message"),
			b.Delete(b.Medium, b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus."),
	),
	c.Example(
		`b.Message(
	b.Large,
	b.MessageTitle("Large message"),
	b.Delete(b.OnClick("alert('click')")),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		b.Message(
			b.Large,
			b.MessageTitle("Large message"),
			b.Delete(b.Large, b.OnClick("alert('click')")),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus."),
	),
)
