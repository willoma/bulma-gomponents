package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var message = c.NewPage(
	"Message", "Message", "/message",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Message"), " constructor creates a message.",
		),
		c.Modifiers(
			c.Row("b.Dark", "Set color to dark"),
			c.Row("b.Primary", "Set color to primary"),
			c.Row("b.Link", "Set color to link"),
			c.Row("b.Info", "Set color to info"),
			c.Row("b.Success", "Set color to success"),
			c.Row("b.Warning", "Set color to warning"),
			c.Row("b.Danger", "Set color to danger"),
		),
		c.Children(
			c.Row("b.OnHeader(...any)", "Apply children to the ", e.Code(`<div class="message-header">`), " element"),
			c.Row("b.OnBody(...any)", "Apply children to the ", e.Code(`<div class="message-body">`), " element"),
			c.Row("b.OnMessage(...any)", "Apply children to the ", e.Code(`<article class="message">`), " element"),
			c.Row("b.MessageTitle", "Set message title as a ", e.Code("<p>"), " element in the ", e.Code(`<div class="message-header">`), " element"),
			c.Row("b.Delete(...any)", "Apply child to the ", e.Code(`<div class="message-header">`), " element"),
			c.RowClassesStyles("Apply child to the ", e.Code(`<article class="message">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<article class="message">`), " element"),
			c.RowNodeElement("Add element to the ", e.Code(`<div class="message-body">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="message-body">`), " element"),
		),
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
