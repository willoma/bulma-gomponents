package docs

import (
	e "github.com/willoma/gomplements"
	"maragu.dev/gomponents/html"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
	"github.com/willoma/bulma-gomponents/fa"
)

var modal = c.NewPage(
	"Modal", "Modal", "/modal",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Modal"), " constructor creates a modal.",
		),
		c.Children(
			c.Row("b.OnModal(...any)", "Apply children to the ", e.Code(`<div class="modal">`), " element"),
			c.Row("b.OnContent(...any)", "Apply children to the ", e.Code(`<div class="modal-content">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<div class="modal">`), " element"),
			c.RowNodeElement("Add element to the ", e.Code(`<div class="modal-content">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="modal-content">`), " element"),
		),
		e.P(
			"The ", e.Code("b.ModalCard"), " constructor creates a card modal.",
		),
		c.Children(
			c.Row("b.OnModal(...any)", "Apply children to the ", e.Code(`<div class="modal">`), " element"),

			c.Row("b.OnCard(...any)", "Apply children to the ", e.Code(`<div class="modal-card">`), " element"),

			c.Row("b.ModalCardHead(...any)", "Apply children to the ", e.Code(`<div class="modal-card-head">`), " element"),
			c.Row("b.ModalCardTitle(...any)", "Add title to the ", e.Code(`<div class="modal-card-head">`), " element"),
			c.Row("b.ModalCardTitleWithClose(string)", "Add title with a close button to the ", e.Code(`<div class="modal-card-head">`), " element"),
			c.Row("b.ModalCardFoot(...any)", "Apply children to the ", e.Code(`<div class="modal-card-foot">`), " element"),
			c.RowNodeAttribute("Apply attribute to the ", e.Code(`<div class="modal">`), " element"),
			c.RowNodeElement("Add element to the ", e.Code(`<div class="modal-card-body">`), " element"),
			c.RowDefault("Apply child to the ", e.Code(`<div class="modal-card-body">`), " element"),
		),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/components/modal/",
	c.Example(
		`b.Button(b.Primary, b.Large, e.OnClick(b.JSOpen("modal")), "Launch example modal"),
b.Modal(
	e.ID("modal"),
	b.Box(
		b.Media(
			b.MediaLeft(b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq64, b.ImgAlt("Image"))),
			b.Content(
				e.P(
					e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"), e.Br(),
					"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.",
				),
			),
			b.Level(b.Mobile, b.LevelLeft(
				e.A(e.AriaLabel("retweet"), fa.Icon(fa.Solid, "retweet", b.Small)),
				e.A(e.AriaLabel("like"), fa.Icon(fa.Solid, "heart", b.Small)),
			)),
		),
	),
)`,
		b.Button(b.Primary, b.Large, e.OnClick(b.JSOpen("modal")), "Launch example modal"),
		b.Modal(
			e.ID("modal"),
			b.Box(
				b.Media(
					b.MediaLeft(b.ImageImg("https://bulma.io/assets/images/placeholders/128x128.png", b.ImgSq64, b.ImgAlt("Image"))),
					b.Content(
						e.P(
							e.Strong("John Smith"), " ", e.Small("@johnsmith"), " ", e.Small("31m"), e.Br(),
							"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.",
						),
					),
					b.Level(b.Mobile, b.LevelLeft(
						e.A(e.AriaLabel("retweet"), fa.Icon(fa.Solid, "retweet", b.Small)),
						e.A(e.AriaLabel("like"), fa.Icon(fa.Solid, "heart", b.Small)),
					)),
				),
			),
		),
	),
).Subsection(
	"Image modal",
	"https://bulma.io/documentation/components/modal/#image-modal",
	c.Example(
		`b.Button(b.Primary, b.Large, e.OnClick(b.JSOpen("modal-image")), "Launch image modal"),
b.Modal(
	e.ID("modal-image"),
	b.ImageImg("https://bulma.io/assets/images/placeholders/1280x960.png", html.P, b.Img4By3),
)`,
		b.Button(b.Primary, b.Large, e.OnClick(b.JSOpen("modal-image")), "Launch image modal"),
		b.Modal(
			e.ID("modal-image"),
			b.ImageImg("https://bulma.io/assets/images/placeholders/1280x960.png", html.P, b.Img4By3),
		),
	),
).Subsection(
	"Modal card",
	"https://bulma.io/documentation/components/modal/#modal-card",
	c.Example(
		`b.Button(b.Primary, b.Large, e.OnClick(b.JSOpen("modal-card-example")), "Launch card modal"),
b.ModalCard(
	e.ID("modal-card-example"),
	b.ModalCardTitleWithClose("Modal title"),
	b.Content(
		e.H1("Hello World"),
		e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
		e.H2("Second level"),
		e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
		b.UList(
			"In fermentum leo eu lectus mollis, quis dictum mi aliquet.",
			"Morbi eu nulla lobortis, lobortis est in, fringilla felis.",
			"Aliquam nec felis in sapien venenatis viverra fermentum nec lectus.",
			"Ut non enim metus.",
		),
		e.H3("Third level"),
		e.P("Quisque ante lacus, malesuada ac auctor vitae, congue ", e.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus."),
		b.OList(
			"Donec blandit a lorem id convallis.",
			"Cras gravida arcu at diam gravida gravida.",
			"Integer in volutpat libero.",
			"Donec a diam tellus.",
			"Aenean nec tortor orci.",
			"Quisque aliquam cursus urna, non bibendum massa viverra eget.",
			"Vivamus maximus ultricies pulvinar.",
		),
		e.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis"),
		e.P("Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", e.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie."),
		e.P("Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo."),
		e.P("Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante."),
		e.H4("Fourth level"),
		e.P("Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur."),
		e.P("Maecenas eleifend sollicitudin dui, faucibus sollicitudin auue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien."),
		e.P("Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque."),
		e.H5("Fifth level"),
		e.P("Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio."),
		e.H6("Sixth level"),
		e.P("Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus."),
	),
	b.ModalCardFoot(
		b.Buttons(
			b.Button(e.OnClick(b.JSCloseThisModal), b.Success, "Save changes"),
			b.Button(e.OnClick(b.JSCloseThisModal), "Cancel"),
		),
	),
)`,
		b.Button(b.Primary, b.Large, e.OnClick(b.JSOpen("modal-card-example")), "Launch card modal"),
		b.ModalCard(
			e.ID("modal-card-example"),
			b.ModalCardTitleWithClose("Modal title"),
			b.Content(
				e.H1("Hello World"),
				e.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
				e.H2("Second level"),
				e.P("Curabitur accumsan turpis pharetra ", e.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
				b.UList(
					"In fermentum leo eu lectus mollis, quis dictum mi aliquet.",
					"Morbi eu nulla lobortis, lobortis est in, fringilla felis.",
					"Aliquam nec felis in sapien venenatis viverra fermentum nec lectus.",
					"Ut non enim metus.",
				),
				e.H3("Third level"),
				e.P("Quisque ante lacus, malesuada ac auctor vitae, congue ", e.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus."),
				b.OList(
					"Donec blandit a lorem id convallis.",
					"Cras gravida arcu at diam gravida gravida.",
					"Integer in volutpat libero.",
					"Donec a diam tellus.",
					"Aenean nec tortor orci.",
					"Quisque aliquam cursus urna, non bibendum massa viverra eget.",
					"Vivamus maximus ultricies pulvinar.",
				),
				e.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis"),
				e.P("Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", e.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie."),
				e.P("Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo."),
				e.P("Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante."),
				e.H4("Fourth level"),
				e.P("Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur."),
				e.P("Maecenas eleifend sollicitudin dui, faucibus sollicitudin auue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien."),
				e.P("Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque."),
				e.H5("Fifth level"),
				e.P("Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio."),
				e.H6("Sixth level"),
				e.P("Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus."),
			),
			b.ModalCardFoot(
				b.Buttons(
					b.Button(e.OnClick(b.JSCloseThisModal), b.Success, "Save changes"),
					b.Button(e.OnClick(b.JSCloseThisModal), "Cancel"),
				),
			),
		),
	),
)
