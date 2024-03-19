package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/el"
)

var content = c.NewPage(
	"Content", "Content", "/content",
	"https://bulma.io/documentation/elements/content/",
).Section(
	"Full example",
	"https://bulma.io/documentation/elements/content/#full-example",
	c.Example(
		`b.Content(
	el.H1("Hello World"),
	el.P(
		"Lorem ipsum", el.Sup(el.A("[1]")), " dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque. Sub", el.Sub("script"), " works as well!",
	),
	el.H2("Second level"),
	el.P(
		"Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl.",
	),
	el.Ul(
		el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		el.Li("Ut non enim metus."),
	),
	el.H3("Third level"),
	el.P(
		"Quisque ante lacus, malesuada ac auctor vitae, congue ", b.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus.",
	),
	el.Ol(
		el.Li("Donec blandit a lorem id convallis."),
		el.Li("Cras gravida arcu at diam gravida gravida."),
		el.Li("Integer in volutpat libero."),
		el.Li("Donec a diam tellus."),
		el.Li("Aenean nec tortor orci."),
		el.Li("Quisque aliquam cursus urna, non bibendum massa viverra eget."),
		el.Li("Vivamus maximus ultricies pulvinar."),
	),
	el.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis."),
	el.P(
		"Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", el.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie.",
	),
	"Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo.",
	el.Dl(
		el.Dt("Web"),
		el.Dd("The part of the Internet that contains websites and web pages"),
		el.Dt("HTML"),
		el.Dd("A markup language for creating web pages"),
		el.Dt("CSS"),
		el.Dd("A technology to make HTML look better"),
	),
	"Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante.",
	el.H4("Fourth level"),
	"Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur.",
	el.Pre(
		b.Padding(b.Spacing2),
		b.FontSize7,
		b.Style("tab-size", "4"),
		′<!DOCTYPE html>
<html>
	<head>
		<title>Hello World</title>
	</head>
	<body>
		<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.
		Donec viverra nec nulla vitae mollis.</p>
	</body>
</html>′),
	"Maecenas eleifend sollicitudin dui, faucibus sollicitudin augue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien.",
	el.Table(
		el.THead(
			el.Tr(
				el.Th("One"),
				el.Th("Two"),
			),
		),
		el.TBody(
			el.Tr(
				el.Td("Three"),
				el.Td("Four"),
			),
			el.Tr(
				el.Td("Five"),
				el.Td("Six"),
			),
			el.Tr(
				el.Td("Seven"),
				el.Td("Eight"),
			),
			el.Tr(
				el.Td("Nine"),
				el.Td("Ten"),
			),
			el.Tr(
				el.Td("Eleven"),
				el.Td("Twelve"),
			),
		),
	),
	"Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque.",
	el.H5("Fifth level"),
	"Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio.",
	el.Figure(
		b.ImgSrc("https://bulma.io/images/placeholders/256x256.png"),
		b.ImgSrc("https://bulma.io/images/placeholders/256x256.png"),
		el.FigCaption("Figure 1: Some beautiful placeholders"),
	),
	el.H6("Sixth level"),
	"Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus.",
)`,
		b.Content(
			el.H1("Hello World"),
			el.P(
				"Lorem ipsum", el.Sup(el.A("[1]")), " dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque. Sub", el.Sub("script"), " works as well!",
			),
			el.H2("Second level"),
			el.P(
				"Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl.",
			),
			el.Ul(
				el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				el.Li("Ut non enim metus."),
			),
			el.H3("Third level"),
			el.P(
				"Quisque ante lacus, malesuada ac auctor vitae, congue ", b.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus.",
			),
			el.Ol(
				el.Li("Donec blandit a lorem id convallis."),
				el.Li("Cras gravida arcu at diam gravida gravida."),
				el.Li("Integer in volutpat libero."),
				el.Li("Donec a diam tellus."),
				el.Li("Aenean nec tortor orci."),
				el.Li("Quisque aliquam cursus urna, non bibendum massa viverra eget."),
				el.Li("Vivamus maximus ultricies pulvinar."),
			),
			el.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis."),
			el.P(
				"Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", el.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie.",
			),
			"Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo.",
			el.Dl(
				el.Dt("Web"),
				el.Dd("The part of the Internet that contains websites and web pages"),
				el.Dt("HTML"),
				el.Dd("A markup language for creating web pages"),
				el.Dt("CSS"),
				el.Dd("A technology to make HTML look better"),
			),
			"Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante.",
			el.H4("Fourth level"),
			"Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur.",
			el.Pre(
				b.Padding(b.Spacing2),
				b.FontSize7,
				b.Style("tab-size", "4"),
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
			el.Table(
				el.THead(
					el.Tr(
						el.Th("One"),
						el.Th("Two"),
					),
				),
				el.TBody(
					el.Tr(
						el.Td("Three"),
						el.Td("Four"),
					),
					el.Tr(
						el.Td("Five"),
						el.Td("Six"),
					),
					el.Tr(
						el.Td("Seven"),
						el.Td("Eight"),
					),
					el.Tr(
						el.Td("Nine"),
						el.Td("Ten"),
					),
					el.Tr(
						el.Td("Eleven"),
						el.Td("Twelve"),
					),
				),
			),
			"Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque.",
			el.H5("Fifth level"),
			"Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio.",
			el.Figure(
				b.ImgSrc("https://bulma.io/images/placeholders/256x256.png"),
				b.ImgSrc("https://bulma.io/images/placeholders/256x256.png"),
				el.FigCaption("Figure 1: Some beautiful placeholders"),
			),
			el.H6("Sixth level"),
			"Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus.",
		),
	),
).Section(
	"Ordered lists alternatives",
	"https://bulma.io/documentation/elements/content/#ordered-lists-alternatives",
	c.Example(
		`b.Content(
	el.Ol(
		html.Type("1"),
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
	el.Ol(
		html.Type("A"),
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
	el.Ol(
		html.Type("a"),
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
	el.Ol(
		html.Type("I"),
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
	el.Ol(
		html.Type("i"),
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
)`,
		b.Content(
			el.Ol(
				html.Type("1"),
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
			el.Ol(
				html.Type("A"),
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
			el.Ol(
				html.Type("a"),
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
			el.Ol(
				html.Type("I"),
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
			el.Ol(
				html.Type("i"),
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
		),
	),
	c.Example(
		`b.Content(
	el.Ol(
		b.OlLowerAlpha,
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
	el.Ol(
		b.OlLowerRoman,
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
	el.Ol(
		b.OlUpperAlpha,
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
	el.Ol(
		b.OlUpperRoman,
		el.Li("Coffee"),
		el.Li("Tea"),
		el.Li("Milk"),
	),
)`,
		b.Content(
			el.Ol(
				b.OlLowerAlpha,
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
			el.Ol(
				b.OlLowerRoman,
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
			el.Ol(
				b.OlUpperAlpha,
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
			el.Ol(
				b.OlUpperRoman,
				el.Li("Coffee"),
				el.Li("Tea"),
				el.Li("Milk"),
			),
		),
	),
).Section(
	"Sizes",
	"https://bulma.io/documentation/elements/content/#sizes",
	c.Example(
		`b.Content(
	b.Small,
	el.H1("Hello World"),
	el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	el.H2("Second level"),
	el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	el.Ul(
		el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		el.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Small,
			el.H1("Hello World"),
			el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			el.H2("Second level"),
			el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			el.Ul(
				el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				el.Li("Ut non enim metus."),
			),
		),
	),
	c.Example(
		`b.Content(
	b.Normal,
	el.H1("Hello World"),
	el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	el.H2("Second level"),
	el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	el.Ul(
		el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		el.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Normal,
			el.H1("Hello World"),
			el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			el.H2("Second level"),
			el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			el.Ul(
				el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				el.Li("Ut non enim metus."),
			),
		),
	),
	c.Example(
		`b.Content(
	b.Medium,
	el.H1("Hello World"),
	el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	el.H2("Second level"),
	el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	el.Ul(
		el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		el.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Medium,
			el.H1("Hello World"),
			el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			el.H2("Second level"),
			el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			el.Ul(
				el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				el.Li("Ut non enim metus."),
			),
		),
	),
	c.Example(
		`b.Content(
	b.Large,
	el.H1("Hello World"),
	el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
	el.H2("Second level"),
	el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
	el.Ul(
		el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
		el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
		el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
		el.Li("Ut non enim metus."),
	),
)`,
		b.Content(
			b.Large,
			el.H1("Hello World"),
			el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
			el.H2("Second level"),
			el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
			el.Ul(
				el.Li("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				el.Li("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				el.Li("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."),
				el.Li("Ut non enim metus."),
			),
		),
	),
)
