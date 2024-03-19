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
	"https://bulma.io/documentation/components/message/",
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
).Section(
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
).Section(
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
).Section(
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
