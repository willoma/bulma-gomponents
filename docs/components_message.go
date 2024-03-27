package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/easy"
	"github.com/willoma/bulma-gomponents/el"
)

var message = c.NewPage(
	"Message", "Message", "/message",
	"",

	b.Content(
		el.P(
			"The ", el.Code("b.Message"), " constructor creates a message. It accepts the following values additionally to the standard set of children types:",
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
		"The ", el.Code("easy.Message"), " constructor creates a message. It accepts the same colors definitions as the ", el.Code("b.Message"), " constructor as well as the following values additionally to the standard set of children types:",
	),
	b.DList(
		el.Code("b.Inner(any)"),
		[]any{"forcibly apply the child to the body element"},

		el.Code("b.Outer(any)"),
		[]any{"forcibly apply the child to the message element"},

		el.Code("easy.MessageTitle"),
		"Include a header with the provided title",

		el.Code("easy.MessageDeleteOnClick"),
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
	b.MessageHeader(
		el.P("Hello world"),
		b.Delete(html.Aria("label", "delete")),
	),
	b.MessageBody(
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
	),
)`,
		b.Message(
			b.MessageHeader(
				el.P("Hello world"),
				b.Delete(html.Aria("label", "delete")),
			),
			b.MessageBody(
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
			),
		),
	),
	b.Content(el.P(el.Em("Bulma-Gomponents"), " provides the ", el.Code("easy.Message"), " helper:")),
	c.Example(
		`easy.Message(
	easy.MessageTitle("Hello world"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			easy.MessageTitle("Hello world"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/components/message/#colors",
	c.Example(
		`easy.Message(
	b.Dark,
	easy.MessageTitle("Dark"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Dark,
			easy.MessageTitle("Dark"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Primary,
	easy.MessageTitle("Primary"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Primary,
			easy.MessageTitle("Primary"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Link,
	easy.MessageTitle("Link"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Link,
			easy.MessageTitle("Link"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Info,
	easy.MessageTitle("Info"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Info,
			easy.MessageTitle("Info"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Success,
	easy.MessageTitle("Success"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Success,
			easy.MessageTitle("Success"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Primary,
	easy.MessageTitle("Warning"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Warning,
			easy.MessageTitle("Warning"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Danger,
	easy.MessageTitle("Danger"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Danger,
			easy.MessageTitle("Danger"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Message body only",
	"https://bulma.io/documentation/components/message/#message-body-only",
	c.Example(
		`easy.Message(
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Dark,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Dark,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Primary,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Primary,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Link,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Link,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Info,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Info,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Success,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Success,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Warning,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Warning,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
	c.Example(
		`easy.Message(
	b.Danger,
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
)`,
		easy.Message(
			b.Danger,
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus. Donec sodales, arcu et sollicitudin porttitor, tortor urna tempor ligula, id porttitor mi magna a neque. Donec dui urna, vehicula et sem eget, facilisis sodales sem.",
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/components/message/#sizes",
	c.Example(
		`easy.Message(
	b.Small,
	easy.MessageTitle("Small message"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		easy.Message(
			b.Small,
			easy.MessageTitle("Small message"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
		),
	),
	c.Example(
		`easy.Message(
	easy.MessageTitle("Normal message"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		easy.Message(
			easy.MessageTitle("Normal message"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus."),
	),
	c.Example(
		`easy.Message(
	b.Medium,
	easy.MessageTitle("Medium message"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		easy.Message(
			b.Medium,
			easy.MessageTitle("Medium message"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus."),
	),
	c.Example(
		`easy.Message(
	b.Large,
	easy.MessageTitle("Large message"),
	easy.MessageDeleteOnClick(′alert("click")′),
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus.",
)`,
		easy.Message(
			b.Large,
			easy.MessageTitle("Large message"),
			easy.MessageDeleteOnClick(`alert("click")`),
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. ", el.Strong("Pellentesque risus mi"), " tempus quis placerat ut, porta nec nulla. Nullam gravida purus diam, et dictum ", el.A("felis venenatis"), " efficitur. Aenean ac ", el.Em("eleifend lacus"), " in mollis lectus."),
	),
)
