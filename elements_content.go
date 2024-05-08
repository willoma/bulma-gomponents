package docs

import (
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var content = c.NewPage(
	"Content", "Content", "/content",
	"",
	b.Content(
		e.P(
			"The ", e.Code("b.Content"), " constructor creates a block make to contain ", e.Strong("WYSIWYG"), " text.",
		),
		c.Modifiers(
			c.Row("b.Small", "Set size to small"),
			c.Row("b.Normal", "Set size to normal"),
			c.Row("b.Medium", "Set size to medium"),
			c.Row("b.Large", "Set size to large"),
		),
		c.Children(
			c.Row("string", "Wrap text in a ", e.Code("<p>"), " element"),
			c.Row("b.NoP(...any)", "Remove the ", e.Code("<p>"), " wrapper for the provided children"),
			c.RowDefault("Apply child to the content, as-is"),
		),

		e.P("You may apply the following modifiers on ", e.Code("e.Ol"), " children in order to change their style:"),
		c.Modifiers(
			c.Row("b.OlLowerAlpha", "Lowercase letters"),
			c.Row("b.OlLowerRoman", "Lowercase roman numbers"),
			c.Row("b.OlUpperAlpha", "Uppercase letters"),
			c.Row("b.OlUpperRoman", "Uppercase roman numbers"),
		),
	),
).Section(
	"Bulma examples",
	"",
).Subsection(
	"Full example",
	"https://bulma.io/documentation/elements/content/#full-example",
	c.Example(
		`b.Content(
	e.H1("Hello World"),
	e.P(
		"Lorem ipsum", e.Sup(e.A("[1]")), " dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque. Sub", e.Sub("script"), " works as well!",
	),
	e.H2("Second level"),
	e.P(
		"Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl.",
	),
	e.Ul(
		e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		e.Li("Ut non enim metus."),
	),
	e.H3("Third level"),
	e.P(
		"Quisque ante lacus, malesuada ac auctor vitae, congue ", e.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus.",
	),
	e.Ol(
		e.Li("Donec blandit a lorem id convallis."),
		e.Li("Cras gravida arcu at diam gravida gravida."),
		e.Li("Integer in volutpat libero."),
		e.Li("Donec a diam tellus."),
		e.Li("Aenean nec tortor orci."),
		e.Li("Quisque aliquam cursus urna, non bibendum massa viverra eget."),
		e.Li("Vivamus maximus ultricies pulvinar."),
	),
	e.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis."),
	e.P(
		"Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", e.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie.",
	),
	"Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo.",
	e.Dl(
		e.Dt("Web"),
		e.Dd("The part of the Internet that contains websites and web pages"),
		e.Dt("HTML"),
		e.Dd("A markup language for creating web pages"),
		e.Dt("CSS"),
		e.Dd("A technology to make HTML look better"),
	),
	"Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante.",
	e.H4("Fourth level"),
	"Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur.",
	e.Pre(
		b.Padding(2),
		b.FontSize7,
		e.Styles{"tab-size": "4"},
		`+"`"+`<!DOCTYPE html>
<html>
	<head>
		<title>Hello World</title>
	</head>
	<body>
		<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.
		Donec viverra nec nulla vitae mollis.</p>
	</body>
</html>`+"`"+`),
	"Maecenas eleifend sollicitudin dui, faucibus sollicitudin augue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien.",
	e.Table(
		e.THead(
			e.Tr(
				e.Th("One"),
				e.Th("Two"),
			),
		),
		e.TBody(
			e.Tr(
				e.Td("Three"),
				e.Td("Four"),
			),
			e.Tr(
				e.Td("Five"),
				e.Td("Six"),
			),
			e.Tr(
				e.Td("Seven"),
				e.Td("Eight"),
			),
			e.Tr(
				e.Td("Nine"),
				e.Td("Ten"),
			),
			e.Tr(
				e.Td("Eleven"),
				e.Td("Twelve"),
			),
		),
	),
	"Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque.",
	e.H5("Fifth level"),
	"Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio.",
	e.Figure(
		e.ImgSrc("https://bulma.io/assets/images/placeholders/256x256.png"),
		e.ImgSrc("https://bulma.io/assets/images/placeholders/256x256.png"),
		e.FigCaption("Figure 1: Some beautiful placeholders"),
	),
	e.H6("Sixth level"),
	"Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus.",
)`,
		b.Content(
			e.H1("Hello World"),
			e.P(
				"Lorem ipsum", e.Sup(e.A("[1]")), " dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque. Sub", e.Sub("script"), " works as well!",
			),
			e.H2("Second level"),
			e.P(
				"Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl.",
			),
			e.Ul(
				e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				e.Li("Ut non enim metus."),
			),
			e.H3("Third level"),
			e.P(
				"Quisque ante lacus, malesuada ac auctor vitae, congue ", e.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus.",
			),
			e.Ol(
				e.Li("Donec blandit a lorem id convallis."),
				e.Li("Cras gravida arcu at diam gravida gravida."),
				e.Li("Integer in volutpat libero."),
				e.Li("Donec a diam tellus."),
				e.Li("Aenean nec tortor orci."),
				e.Li("Quisque aliquam cursus urna, non bibendum massa viverra eget."),
				e.Li("Vivamus maximus ultricies pulvinar."),
			),
			e.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis."),
			e.P(
				"Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", e.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie.",
			),
			"Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo.",
			e.Dl(
				e.Dt("Web"),
				e.Dd("The part of the Internet that contains websites and web pages"),
				e.Dt("HTML"),
				e.Dd("A markup language for creating web pages"),
				e.Dt("CSS"),
				e.Dd("A technology to make HTML look better"),
			),
			"Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante.",
			e.H4("Fourth level"),
			"Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur.",
			e.Pre(
				b.Padding(2),
				b.FontSize(7),
				e.Styles{"tab-size": "4"},
				`<!DOCTYPE html>
<html>
	<head>
		<title>Hello World</title>
	</head>
	<body>
		<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.
		Donec viverra nec nulla vitae mollis.</p>
	</body>
</html>`,
			),
			"Maecenas eleifend sollicitudin dui, faucibus sollicitudin augue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien.",
			e.Table(
				e.THead(
					e.Tr(
						e.Th("One"),
						e.Th("Two"),
					),
				),
				e.TBody(
					e.Tr(
						e.Td("Three"),
						e.Td("Four"),
					),
					e.Tr(
						e.Td("Five"),
						e.Td("Six"),
					),
					e.Tr(
						e.Td("Seven"),
						e.Td("Eight"),
					),
					e.Tr(
						e.Td("Nine"),
						e.Td("Ten"),
					),
					e.Tr(
						e.Td("Eleven"),
						e.Td("Twelve"),
					),
				),
			),
			"Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque.",
			e.H5("Fifth level"),
			"Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio.",
			e.Figure(
				e.ImgSrc("https://bulma.io/assets/images/placeholders/256x256.png"),
				e.ImgSrc("https://bulma.io/assets/images/placeholders/256x256.png"),
				e.FigCaption("Figure 1: Some beautiful placeholders"),
			),
			e.H6("Sixth level"),
			"Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus.",
		),
	),
).Subsection(
	"Ordered lists alternatives",
	"https://bulma.io/documentation/elements/content/#ordered-lists-alternatives",
	c.Example(
		`b.Content(
	e.Ol(
		e.Type("1"),
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
	e.Ol(
		e.Type("A"),
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
	e.Ol(
		e.Type("a"),
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
	e.Ol(
		e.Type("I"),
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
	e.Ol(
		e.Type("i"),
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
)`,
		b.Content(
			e.Ol(
				e.Type("1"),
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
			e.Ol(
				e.Type("A"),
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
			e.Ol(
				e.Type("a"),
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
			e.Ol(
				e.Type("I"),
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
			e.Ol(
				e.Type("i"),
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
		),
	),
	c.Example(
		`b.Content(
	e.Ol(
		b.OlLowerAlpha,
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
	e.Ol(
		b.OlLowerRoman,
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
	e.Ol(
		b.OlUpperAlpha,
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
	e.Ol(
		b.OlUpperRoman,
		e.Li("Coffee"),
		e.Li("Tea"),
		e.Li("Milk"),
	),
)`,
		b.Content(
			e.Ol(
				b.OlLowerAlpha,
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
			e.Ol(
				b.OlLowerRoman,
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
			e.Ol(
				b.OlUpperAlpha,
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
			e.Ol(
				b.OlUpperRoman,
				e.Li("Coffee"),
				e.Li("Tea"),
				e.Li("Milk"),
			),
		),
	),
).Subsection(
	"Sizes",
	"https://bulma.io/documentation/elements/content/#sizes",
	c.Example(
		`b.Content(
	b.Small,
	e.H1("Hello World"),
	e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	e.H2("Second level"),
	e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	e.Ul(
		e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		e.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Small,
			e.H1("Hello World"),
			e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			e.H2("Second level"),
			e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			e.Ul(
				e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				e.Li("Ut non enim metus."),
			),
		),
	),
	c.Example(
		`b.Content(
	b.Normal,
	e.H1("Hello World"),
	e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	e.H2("Second level"),
	e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	e.Ul(
		e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		e.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Normal,
			e.H1("Hello World"),
			e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			e.H2("Second level"),
			e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			e.Ul(
				e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				e.Li("Ut non enim metus."),
			),
		),
	),
	c.Example(
		`b.Content(
	b.Medium,
	e.H1("Hello World"),
	e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	e.H2("Second level"),
	e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	e.Ul(
		e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		e.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Medium,
			e.H1("Hello World"),
			e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			e.H2("Second level"),
			e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			e.Ul(
				e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				e.Li("Ut non enim metus."),
			),
		),
	),
	c.Example(
		`b.Content(
	b.Large,
	e.H1("Hello World"),
	e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	e.H2("Second level"),
	e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	e.Ul(
		e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		e.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Large,
			e.H1("Hello World"),
			e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			e.H2("Second level"),
			e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			e.Ul(
				e.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				e.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				e.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				e.Li("Ut non enim metus."),
			),
		),
	),
)
