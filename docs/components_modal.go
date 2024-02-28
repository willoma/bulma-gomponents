package docs

import (
	"github.com/maragudk/gomponents/html"

	b "github.com/willoma/bulma-gomponents"
	c "github.com/willoma/bulma-gomponents/docs/components"
	"github.com/willoma/bulma-gomponents/easy"
	"github.com/willoma/bulma-gomponents/el"
	"github.com/willoma/bulma-gomponents/fa"
)

var modal = c.NewPage(
	"Modal", "Modal", "/modal",
	"https://bulma.io/documentation/components/modal/",
	c.Example(
		`b.Modal(
	"some-id",
	// Any other Bulma elements you want
),`,
		b.Button(
			b.Primary,
			b.Large,
			b.OnClick(b.JSOpen("modal")),
			"Launch example modal",
		),
		b.Modal(
			"modal",
			b.Box(
				b.Media(
					b.MediaLeft(
						b.ImageImg(
							"https://bulma.io/images/placeholders/128x128.png",
							b.ImgSq64,
							b.ImgAlt("Image"),
						),
					),
					b.Content(
						el.P(
							el.Strong("John Smith"), " ",
							el.Small("@johnsmith"), " ", el.Small("31m"),
							el.Br(),
							"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.",
						),
					),
					b.Level(
						b.Mobile,
						b.LevelLeft(
							b.LevelItem(
								html.A,
								html.Aria("label", "retweet"),
								fa.Icon(fa.Solid, "retweet", b.Small),
							),
							b.LevelItem(
								html.A,
								html.Aria("label", "like"),
								fa.Icon(fa.Solid, "heart", b.Small),
							),
						),
					),
				),
			),
		),
	),
).Section(
	"Image modal",
	"https://bulma.io/documentation/components/modal/#image-modal",
	c.Example(
		`b.Modal(
	"some-id",
	b.ImageImg(
		"https://bulma.io/images/placeholders/1280x960.png",
		html.P,
		b.Img4By3,
	),
),`,
		b.Button(
			b.Primary,
			b.Large,
			b.OnClick(b.JSOpen("modal-image")),
			"Launch image modal",
		),
		b.Modal(
			"modal-image",
			b.ImageImg(
				"https://bulma.io/images/placeholders/1280x960.png",
				html.P,
				b.Img4By3,
			),
		),
	),
).Section(
	"Modal card",
	"https://bulma.io/documentation/components/modal/#modal-card",
	c.Example(
		`b.ModalCard(
			"modal-card",
			b.ModalCardHead(
				b.ModalCardTitle("Modal title"),
				b.Delete(
					b.OnClick(b.JSCloseThisModal),
					html.Aria("label", "close"),
				),
			),
			// Content...
			b.ModalCardFoot(
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					b.Success,
					"Save changes",
				),
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					"Cancel",
				),
			),
		)`,
		b.Button(
			b.Primary,
			b.Large,
			b.OnClick(b.JSOpen("modal-card")),
			"Launch card modal",
		),
		b.ModalCard(
			"modal-card",
			b.ModalCardHead(
				b.ModalCardTitle("Modal title"),
				b.Delete(
					b.OnClick(b.JSCloseThisModal),
					html.Aria("label", "close"),
				),
			),
			b.Content(
				el.H1("Hello World"),
				el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
				el.H2("Second level"),
				el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
				b.UList(
					"In fermentum leo eu lectus mollis, quis dictum mi aliquet.",
					"Morbi eu nulla lobortis, lobortis est in, fringilla felis.",
					"Aliquam nec felis in sapien venenatis viverra fermentum nec lectus.",
					"Ut non enim metus.",
				),
				el.H3("Third level"),
				el.P("<p>Quisque ante lacus, malesuada ac auctor vitae, congue ", b.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus."),
				b.OList(
					"Donec blandit a lorem id convallis.",
					"Cras gravida arcu at diam gravida gravida.",
					"Integer in volutpat libero.",
					"Donec a diam tellus.",
					"Aenean nec tortor orci.",
					"Quisque aliquam cursus urna, non bibendum massa viverra eget.",
					"Vivamus maximus ultricies pulvinar.",
				),
				el.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis"),
				el.P("Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", el.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie."),
				el.P("Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo."),
				el.P("Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante."),
				el.H4("Fourth level"),
				el.P("Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur."),
				el.P("Maecenas eleifend sollicitudin dui, faucibus sollicitudin auue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien."),
				el.P("Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque."),
				el.H5("Fifth level"),
				el.P("Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio."),
				el.H6("Sixth level"),
				el.P("Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus."),
			),
			b.ModalCardFoot(
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					b.Success,
					"Save changes",
				),
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					"Cancel",
				),
			),
		),
	),
	b.Content(el.P("or, with the ", el.Code("easy.ModalCardTitleHead"), " helper:")),
	c.Example(
		`b.ModalCard(
			"modal-card2",
			easy.ModalCardHeadTitle("Modal title"),
			// Content...
			b.ModalCardFoot(
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					b.Success,
					"Save changes",
				),
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					"Cancel",
				),
			),
		)`,
		b.Button(
			b.Primary,
			b.Large,
			b.OnClick(b.JSOpen("modal-card2")),
			"Launch card modal",
		),
		b.ModalCard(
			"modal-card2",
			easy.ModalCardHeadTitle("Modal title"),
			b.Content(
				el.H1("Hello World"),
				el.P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque."),
				el.H2("Second level"),
				el.P("Curabitur accumsan turpis pharetra ", el.Strong("augue tincidunt"), " blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl."),
				b.UList(
					"In fermentum leo eu lectus mollis, quis dictum mi aliquet.",
					"Morbi eu nulla lobortis, lobortis est in, fringilla felis.",
					"Aliquam nec felis in sapien venenatis viverra fermentum nec lectus.",
					"Ut non enim metus.",
				),
				el.H3("Third level"),
				el.P("<p>Quisque ante lacus, malesuada ac auctor vitae, congue ", b.AHref("#", "non ante"), ". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus."),
				b.OList(
					"Donec blandit a lorem id convallis.",
					"Cras gravida arcu at diam gravida gravida.",
					"Integer in volutpat libero.",
					"Donec a diam tellus.",
					"Aenean nec tortor orci.",
					"Quisque aliquam cursus urna, non bibendum massa viverra eget.",
					"Vivamus maximus ultricies pulvinar.",
				),
				el.BlockQuote("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis"),
				el.P("Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et ", el.Em("justo sodales"), " elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie."),
				el.P("Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum commodo."),
				el.P("Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla. Nulla facilisi. Nullam ac erat ante."),
				el.H4("Fourth level"),
				el.P("Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum ex efficitur."),
				el.P("Maecenas eleifend sollicitudin dui, faucibus sollicitudin auue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien."),
				el.P("Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus, mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum. Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque."),
				el.H5("Fifth level"),
				el.P("Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend justo. Nam et sollicitudin odio."),
				el.H6("Sixth level"),
				el.P("Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan. Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis. Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui, sed varius sapien odio vitae est. Etiam at cursus metus."),
			),
			b.ModalCardFoot(
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					b.Success,
					"Save changes",
				),
				b.Button(
					b.OnClick(b.JSCloseThisModal),
					"Cancel",
				),
			),
		),
	),
)
