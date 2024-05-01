package docs

import (
	"github.com/maragudk/gomponents/html"
	e "github.com/willoma/gomplements"

	c "bulma-gomponents.docs/components"
	b "github.com/willoma/bulma-gomponents"
)

var tableHeadFoot = []any{
	b.Abbr("Position", "Pos"),
	"Team",
	b.Abbr("Played", "Pld"),
	b.Abbr("Won", "W"),
	b.Abbr("Drawn", "D"),
	b.Abbr("Lost", "L"),
	b.Abbr("Goals for", "GF"),
	b.Abbr("Goals against", "GA"),
	b.Abbr("Goal difference", "GD"),
	b.Abbr("Points", "Pts"),
	"Qualification or relegation",
}

var table = c.NewPage(
	"Table", "Table", "/table",
	"",

	b.Content(
		e.P(
			"The ", e.Code("b.Table"), " constructor creates a table. The ", e.Code("b.ScrollableTable"), " constructor creates a table in a table-container e.Element, making the table scrollable. The following children have a special meaning:",
		),
		b.DList(
			e.Code("b.OnHead(any)"),
			[]any{"Forcibly apply the child to the ", e.Code("<thead>"), " e.Element"},

			e.Code("b.OnBody(any)"),
			[]any{"Forcibly apply the child to the ", e.Code("<tbody>"), " e.Element"},

			e.Code("b.OnFoot(any)"),
			[]any{"Forcibly apply the child to the ", e.Code("<tfoot>"), " e.Element"},

			e.Code("b.Bordered"),
			"Add borders to the table",

			e.Code("b.Striped"),
			"Add stripes to the body rows",

			e.Code("b.Narrow"),
			"Make the cells narrower",

			e.Code("b.Hoverable"),
			"Add a hover effect on the body rows",

			e.Code("b.FullWidth"),
			"Take the whole width",

			e.Code("b.Row(...)"),
			[]any{"Add a row to the table body, its cells are defined as ", e.Code("<td>"), " elements"},

			e.Code("b.HeadRow(...)"),
			[]any{"Add a row to the table header, its cells are defined as ", e.Code("<th>"), " elements"},

			e.Code("b.FootRow(...)"),
			[]any{"Add a row to the table footer, its cells are defined as ", e.Code("<th>"), " elements"},
		),
		e.P("Other children are added to the ", e.Code("<table>"), " e.Element."),
		e.P("The rows constructors accept either children embedded in ", e.Code("b.TCell(...)"), " constructors, or embed their children automatically."),
	),
).Section(
	"Bulma examples", "https://bulma.io/documentation/elements/table/",

	b.Content(e.P("Bulma-styled table are implemented with the ", e.Code("b.Table"), " function.")),

	c.Example(
		`headFoot := []any{
	b.Abbr("Position", "Pos"),
	"Team",
	b.Abbr("Played", "Pld"),
	b.Abbr("Won", "W"),
	b.Abbr("Drawn", "D"),
	b.Abbr("Lost", "L"),
	b.Abbr("Goals for", "GF"),
	b.Abbr("Goals against", "GA"),
	b.Abbr("Goal difference", "GD"),
	b.Abbr("Points", "Pts"),
	"Qualification or relegation",
}
b.Table(
	b.HeadRow(headFoot...),
	b.FootRow(headFoot...),
	b.Row(
		b.TCell(html.Th, "1"),
		b.TCell(
			e.A(
				html.Href("https://en.wikipedia.org/wiki/Leicester_City_F.C."),
				html.TitleAttr("Leicester City F.C."),
				"Leicester City",
			),
			" ",
			e.Strong("(C)"),
		),
		"38",
		"23",
		"12",
		"3",
		"68",
		"36",
		"+32",
		"81",
		b.TCell(
			"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Group_stage"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League group stage"),
		),
	),
	b.Row(
		b.TCell(html.Th, "2"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Arsenal_F.C."), html.TitleAttr("Arsenal F.C."), "Arsenal"),
		"38",
		"20",
		"11",
		"7",
		"65",
		"36",
		"+29",
		"71",
		b.TCell(
			"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Group_stage"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League group stage"),
		),
	),
	b.Row(
		b.TCell(html.Th, "3"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Tottenham_Hotspur_F.C."), html.TitleAttr("Tottenham Hotspur F.C."), "Tottenham Hotspur"),
		"38",
		"19",
		"13",
		"6",
		"69",
		"35",
		"+34",
		"70",
		b.TCell(
			"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Group_stage"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League group stage"),
		),
	),
	b.Row(
		b.Selected,
		b.TCell(html.Th, "4"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Manchester_City_F.C."), html.TitleAttr("Manchester City F.C."), "Manchester City"),
		"38",
		"19",
		"9",
		"10",
		"71",
		"41",
		"+30",
		"66",
		b.TCell(
			"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Play-off_round"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League play-off round"),
		),
	),
	b.Row(
		b.TCell(html.Th, "5"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Manchester_United_F.C."), html.TitleAttr("Manchester United F.C."), "Manchester United"),
		"38",
		"19",
		"9",
		"10",
		"49",
		"35",
		"+14",
		"66",
		b.TCell(
			"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Europa_League#Group_stage"), html.TitleAttr("2016–17 UEFA Europa League"), "Europa League group stage"),
		),
	),
	b.Row(
		b.TCell(html.Th, "6"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Southampton_F.C."), html.TitleAttr("Southampton F.C."), "Southampton"),
		"38",
		"18",
		"9",
		"11",
		"59",
		"41",
		"+18",
		"63",
		b.TCell(
			"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Europa_League#Group_stage"), html.TitleAttr("2016–17 UEFA Europa League"), "Europa League group stage"),
		),
	),
	b.Row(
		b.TCell(html.Th, "7"),
		e.A(html.Href("https://en.wikipedia.org/wiki/West_Ham_United_F.C."), html.TitleAttr("West Ham United F.C."), "West Ham United"),
		"38",
		"16",
		"14",
		"8",
		"65",
		"51",
		"+14",
		"62",
		b.TCell(
			"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Europa_League#Third_qualifying_round"), html.TitleAttr("2016–17 UEFA Europa League"), "Europa League third qualifying round"),
		),
	),
	b.Row(
		b.TCell(html.Th, "8"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Liverpool_F.C."), html.TitleAttr("Liverpool F.C."), "Liverpool"),
		"38",
		"16",
		"12",
		"10",
		"63",
		"50",
		"+13",
		"60",
		"",
	),
	b.Row(
		b.TCell(html.Th, "9"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Stoke_City_F.C."), html.TitleAttr("Stoke City F.C."), "Stoke City"),
		"38",
		"14",
		"9",
		"15",
		"41",
		"55",
		b.TCell("-14"),
		"51",
		"",
	),
	b.Row(
		b.TCell(html.Th, "10"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Chelsea_F.C."), html.TitleAttr("Chelsea F.C."), "Chelsea"),
		"38",
		"12",
		"14",
		"12",
		"59",
		"53",
		"+6",
		"50",
		"",
	),
	b.Row(
		b.TCell(html.Th, "11"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Everton_F.C."), html.TitleAttr("Everton F.C."), "Everton"),
		"38",
		"11",
		"14",
		"13",
		"59",
		"55",
		"+4",
		"47",
		"",
	),
	b.Row(
		b.TCell(html.Th, "12"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Swansea_City_A.F.C."), html.TitleAttr("Swansea City A.F.C."), "Swansea City"),
		"38",
		"12",
		"11",
		"15",
		"42",
		"52",
		"-10",
		"47",
		"",
	),
	b.Row(
		b.TCell(html.Th, "13"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Watford_F.C."), html.TitleAttr("Watford F.C."), "Watford"),
		"38",
		"12",
		"9",
		"17",
		"40",
		"50",
		"-10",
		"45",
		"",
	),
	b.Row(
		b.TCell(html.Th, "14"),
		e.A(html.Href("https://en.wikipedia.org/wiki/West_Bromwich_Albion_F.C."), html.TitleAttr("West Bromwich Albion F.C."), "West Bromwich Albion"),
		"38",
		"10",
		"13",
		"15",
		"34",
		"48",
		"-14",
		"43",
		"",
	),
	b.Row(
		b.TCell(html.Th, "15"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Crystal_Palace_F.C."), html.TitleAttr("Crystal Palace F.C."), "Crystal Palace"),
		"38",
		"11",
		"9",
		"18",
		"39",
		"51",
		"-12",
		"42",
		"",
	),
	b.Row(
		b.TCell(html.Th, "16"),
		e.A(html.Href("https://en.wikipedia.org/wiki/A.F.C._Bournemouth"), html.TitleAttr("A.F.C. Bournemouth"), "AFC Bournemouth"),
		"38",
		"11",
		"9",
		"18",
		"45",
		"67",
		"-22",
		"42",
		"",
	),
	b.Row(
		b.TCell(html.Th, "17"),
		e.A(html.Href("https://en.wikipedia.org/wiki/Sunderland_A.F.C."), html.TitleAttr("Sunderland A.F.C."), "Sunderland"),
		"38",
		"9",
		"12",
		"17",
		"48",
		"62",
		"-14",
		"39",
		"",
	),
	b.Row(
		b.TCell(html.Th, "18"),
		b.TCell(
			e.A(
				html.Href("https://en.wikipedia.org/wiki/Newcastle_United_F.C."),
				html.TitleAttr("Newcastle United F.C."), "Newcastle United",
			),
			" ",
			e.Strong("(R)"),
		),
		"38",
		"9",
		"10",
		"19",
		"44",
		"65",
		"-21",
		"37",
		b.TCell(
			"Relegation to the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_Football_League_Championship"), html.TitleAttr("2016–17 Football League Championship"), "Football League Championship"),
		),
	),
	b.Row(
		b.TCell(html.Th, "19"),
		b.TCell(
			e.A(
				html.Href("https://en.wikipedia.org/wiki/Norwich_City_F.C."),
				html.TitleAttr("Norwich City F.C."),
				"Norwich City",
			),
			" ",
			e.Strong("(R)"),
		),
		"38",
		"9",
		"7",
		"22",
		"39",
		"67",
		"-28",
		"34",
		b.TCell(
			"Relegation to the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_Football_League_Championship"), html.TitleAttr("2016–17 Football League Championship"), "Football League Championship"),
		),
	),
	b.Row(
		b.TCell(html.Th, "20"),
		b.TCell(
			e.A(
				html.Href("https://en.wikipedia.org/wiki/Aston_Villa_F.C."),
				html.TitleAttr("Aston Villa F.C."),
				"Aston Villa",
			),
			" ",
			e.Strong("(R)"),
		),
		"38",
		"3",
		"8",
		"27",
		"27",
		"76",
		"-49",
		"17",
		b.TCell(
			"Relegation to the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_Football_League_Championship"), html.TitleAttr("2016–17 Football League Championship"), "Football League Championship"),
		),
	),
)`,
		b.Table(
			b.HeadRow(tableHeadFoot...),
			b.FootRow(tableHeadFoot...),
			b.Row(
				b.TCell(html.Th, "1"),
				b.TCell(
					e.A(
						html.Href("https://en.wikipedia.org/wiki/Leicester_City_F.C."),
						html.TitleAttr("Leicester City F.C."),
						"Leicester City",
					),
					" ",
					e.Strong("(C)"),
				),
				"38",
				"23",
				"12",
				"3",
				"68",
				"36",
				"+32",
				"81",
				b.TCell(
					"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Group_stage"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League group stage"),
				),
			),
			b.Row(
				b.TCell(html.Th, "2"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Arsenal_F.C."), html.TitleAttr("Arsenal F.C."), "Arsenal"),
				"38",
				"20",
				"11",
				"7",
				"65",
				"36",
				"+29",
				"71",
				b.TCell(
					"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Group_stage"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League group stage"),
				),
			),
			b.Row(
				b.TCell(html.Th, "3"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Tottenham_Hotspur_F.C."), html.TitleAttr("Tottenham Hotspur F.C."), "Tottenham Hotspur"),
				"38",
				"19",
				"13",
				"6",
				"69",
				"35",
				"+34",
				"70",
				b.TCell(
					"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Group_stage"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League group stage"),
				),
			),
			b.Row(
				b.Selected,
				b.TCell(html.Th, "4"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Manchester_City_F.C."), html.TitleAttr("Manchester City F.C."), "Manchester City"),
				"38",
				"19",
				"9",
				"10",
				"71",
				"41",
				"+30",
				"66",
				b.TCell(
					"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Champions_League#Play-off_round"), html.TitleAttr("2016–17 UEFA Champions League"), "Champions League play-off round"),
				),
			),
			b.Row(
				b.TCell(html.Th, "5"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Manchester_United_F.C."), html.TitleAttr("Manchester United F.C."), "Manchester United"),
				"38",
				"19",
				"9",
				"10",
				"49",
				"35",
				"+14",
				"66",
				b.TCell(
					"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Europa_League#Group_stage"), html.TitleAttr("2016–17 UEFA Europa League"), "Europa League group stage"),
				),
			),
			b.Row(
				b.TCell(html.Th, "6"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Southampton_F.C."), html.TitleAttr("Southampton F.C."), "Southampton"),
				"38",
				"18",
				"9",
				"11",
				"59",
				"41",
				"+18",
				"63",
				b.TCell(
					"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Europa_League#Group_stage"), html.TitleAttr("2016–17 UEFA Europa League"), "Europa League group stage"),
				),
			),
			b.Row(
				b.TCell(html.Th, "7"),
				e.A(html.Href("https://en.wikipedia.org/wiki/West_Ham_United_F.C."), html.TitleAttr("West Ham United F.C."), "West Ham United"),
				"38",
				"16",
				"14",
				"8",
				"65",
				"51",
				"+14",
				"62",
				b.TCell(
					"Qualification for the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_UEFA_Europa_League#Third_qualifying_round"), html.TitleAttr("2016–17 UEFA Europa League"), "Europa League third qualifying round"),
				),
			),
			b.Row(
				b.TCell(html.Th, "8"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Liverpool_F.C."), html.TitleAttr("Liverpool F.C."), "Liverpool"),
				"38",
				"16",
				"12",
				"10",
				"63",
				"50",
				"+13",
				"60",
				"",
			),
			b.Row(
				b.TCell(html.Th, "9"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Stoke_City_F.C."), html.TitleAttr("Stoke City F.C."), "Stoke City"),
				"38",
				"14",
				"9",
				"15",
				"41",
				"55",
				b.TCell("-14"),
				"51",
				"",
			),
			b.Row(
				b.TCell(html.Th, "10"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Chelsea_F.C."), html.TitleAttr("Chelsea F.C."), "Chelsea"),
				"38",
				"12",
				"14",
				"12",
				"59",
				"53",
				"+6",
				"50",
				"",
			),
			b.Row(
				b.TCell(html.Th, "11"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Everton_F.C."), html.TitleAttr("Everton F.C."), "Everton"),
				"38",
				"11",
				"14",
				"13",
				"59",
				"55",
				"+4",
				"47",
				"",
			),
			b.Row(
				b.TCell(html.Th, "12"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Swansea_City_A.F.C."), html.TitleAttr("Swansea City A.F.C."), "Swansea City"),
				"38",
				"12",
				"11",
				"15",
				"42",
				"52",
				"-10",
				"47",
				"",
			),
			b.Row(
				b.TCell(html.Th, "13"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Watford_F.C."), html.TitleAttr("Watford F.C."), "Watford"),
				"38",
				"12",
				"9",
				"17",
				"40",
				"50",
				"-10",
				"45",
				"",
			),
			b.Row(
				b.TCell(html.Th, "14"),
				e.A(html.Href("https://en.wikipedia.org/wiki/West_Bromwich_Albion_F.C."), html.TitleAttr("West Bromwich Albion F.C."), "West Bromwich Albion"),
				"38",
				"10",
				"13",
				"15",
				"34",
				"48",
				"-14",
				"43",
				"",
			),
			b.Row(
				b.TCell(html.Th, "15"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Crystal_Palace_F.C."), html.TitleAttr("Crystal Palace F.C."), "Crystal Palace"),
				"38",
				"11",
				"9",
				"18",
				"39",
				"51",
				"-12",
				"42",
				"",
			),
			b.Row(
				b.TCell(html.Th, "16"),
				e.A(html.Href("https://en.wikipedia.org/wiki/A.F.C._Bournemouth"), html.TitleAttr("A.F.C. Bournemouth"), "AFC Bournemouth"),
				"38",
				"11",
				"9",
				"18",
				"45",
				"67",
				"-22",
				"42",
				"",
			),
			b.Row(
				b.TCell(html.Th, "17"),
				e.A(html.Href("https://en.wikipedia.org/wiki/Sunderland_A.F.C."), html.TitleAttr("Sunderland A.F.C."), "Sunderland"),
				"38",
				"9",
				"12",
				"17",
				"48",
				"62",
				"-14",
				"39",
				"",
			),
			b.Row(
				b.TCell(html.Th, "18"),
				b.TCell(
					e.A(
						html.Href("https://en.wikipedia.org/wiki/Newcastle_United_F.C."),
						html.TitleAttr("Newcastle United F.C."),
						"Newcastle United",
					),
					" ",
					e.Strong("(R)"),
				),
				"38",
				"9",
				"10",
				"19",
				"44",
				"65",
				"-21",
				"37",
				b.TCell(
					"Relegation to the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_Football_League_Championship"), html.TitleAttr("2016–17 Football League Championship"), "Football League Championship"),
				),
			),
			b.Row(
				b.TCell(html.Th, "19"),
				b.TCell(
					e.A(
						html.Href("https://en.wikipedia.org/wiki/Norwich_City_F.C."),
						html.TitleAttr("Norwich City F.C."),
						"Norwich City",
					),
					" ",
					e.Strong("(R)"),
				),
				"38",
				"9",
				"7",
				"22",
				"39",
				"67",
				"-28",
				"34",
				b.TCell(
					"Relegation to the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_Football_League_Championship"), html.TitleAttr("2016–17 Football League Championship"), "Football League Championship"),
				),
			),
			b.Row(
				b.TCell(html.Th, "20"),
				b.TCell(
					e.A(
						html.Href("https://en.wikipedia.org/wiki/Aston_Villa_F.C."),
						html.TitleAttr("Aston Villa F.C."),
						"Aston Villa",
					),
					" ",
					e.Strong("(R)"),
				),
				"38",
				"3",
				"8",
				"27",
				"27",
				"76",
				"-49",
				"17",
				b.TCell(
					"Relegation to the ", e.A(html.Href("https://en.wikipedia.org/wiki/2016%E2%80%9317_Football_League_Championship"), html.TitleAttr("2016–17 Football League Championship"), "Football League Championship"),
				),
			),
		),
	),
).Subsection(
	"Colors",
	"https://bulma.io/documentation/elements/table/#colors",
	c.Example(
		`b.Table(
	b.Bordered,
	b.Row(b.TCell(html.Th, b.Link, "Link th cell"), "Two", b.TCell(b.Link, "Link td cell"), "Four", "Five"),
	b.Row(b.Link, b.TCell(html.Th, "Link row"), "Two", "Three", "Four", "Five"),
	b.Row(b.TCell(html.Th, b.Primary, "Primary th cell"), "Two", b.TCell(b.Primary, "Primary td cell"), "Four", "Five"),
	b.Row(b.Primary, b.TCell(html.Th, "Primary row"), "Two", "Three", "Four", "Five"),
	b.Row(b.TCell(html.Th, b.Info, "Info th cell"), "Two", b.TCell(b.Info, "Info td cell"), "Four", "Five"),
	b.Row(b.Info, b.TCell(html.Th, "Info row"), "Two", "Three", "Four", "Five"),
	b.Row(b.TCell(html.Th, b.Success, "Success th cell"), "Two", b.TCell(b.Success, "Success td cell"), "Four", "Five"),
	b.Row(b.Success, b.TCell(html.Th, "Success row"), "Two", "Three", "Four", "Five"),
	b.Row(b.TCell(html.Th, b.Warning, "Warning th cell"), "Two", b.TCell(b.Warning, "Warning td cell"), "Four", "Five"),
	b.Row(b.Warning, b.TCell(html.Th, "Warning row"), "Two", "Three", "Four", "Five"),
	b.Row(b.TCell(html.Th, b.Danger, "Danger th cell"), "Two", b.TCell(b.Danger, "Danger td cell"), "Four", "Five"),
	b.Row(b.Danger, b.TCell(html.Th, "Danger row"), "Two", "Three", "Four", "Five"),
)`,
		b.Table(
			b.Bordered,
			b.Row(b.TCell(html.Th, b.Link, "Link th cell"), "Two", b.TCell(b.Link, "Link td cell"), "Four", "Five"),
			b.Row(b.Link, b.TCell(html.Th, "Link row"), "Two", "Three", "Four", "Five"),
			b.Row(b.TCell(html.Th, b.Primary, "Primary th cell"), "Two", b.TCell(b.Primary, "Primary td cell"), "Four", "Five"),
			b.Row(b.Primary, b.TCell(html.Th, "Primary row"), "Two", "Three", "Four", "Five"),
			b.Row(b.TCell(html.Th, b.Info, "Info th cell"), "Two", b.TCell(b.Info, "Info td cell"), "Four", "Five"),
			b.Row(b.Info, b.TCell(html.Th, "Info row"), "Two", "Three", "Four", "Five"),
			b.Row(b.TCell(html.Th, b.Success, "Success th cell"), "Two", b.TCell(b.Success, "Success td cell"), "Four", "Five"),
			b.Row(b.Success, b.TCell(html.Th, "Success row"), "Two", "Three", "Four", "Five"),
			b.Row(b.TCell(html.Th, b.Warning, "Warning th cell"), "Two", b.TCell(b.Warning, "Warning td cell"), "Four", "Five"),
			b.Row(b.Warning, b.TCell(html.Th, "Warning row"), "Two", "Three", "Four", "Five"),
			b.Row(b.TCell(html.Th, b.Danger, "Danger th cell"), "Two", b.TCell(b.Danger, "Danger td cell"), "Four", "Five"),
			b.Row(b.Danger, b.TCell(html.Th, "Danger row"), "Two", "Three", "Four", "Five"),
		),
	),
).Subsection(
	"Modifiers",
	"https://bulma.io/documentation/elements/table/#modifiers",
	c.Example(
		`b.Table(
	b.Bordered,
	b.HeadRow("One", "Two"),
	b.Row("Three", "Four"),
)`,
		b.Table(
			b.Bordered,
			b.HeadRow("One", "Two"),
			b.Row("Three", "Four"),
		),
	),
	c.Example(
		`b.Table(
	b.Striped,
	b.HeadRow("One", "Two"),
	b.Row("Three", "Four"),
	b.Row("Five", "Six"),
	b.Row("Seven", "Eight"),
	b.Row("Nine", "Ten"),
	b.Row("Eleven", "Twelve"),
)`,
		b.Table(
			b.Striped,
			b.HeadRow("One", "Two"),
			b.Row("Three", "Four"),
			b.Row("Five", "Six"),
			b.Row("Seven", "Eight"),
			b.Row("Nine", "Ten"),
			b.Row("Eleven", "Twelve"),
		),
	),
	c.Example(
		`b.Table(
	b.Narrow,
	b.HeadRow("One", "Two"),
	b.Row("Three", "Four"),
	b.Row("Five", "Six"),
	b.Row("Seven", "Eight"),
	b.Row("Nine", "Ten"),
	b.Row("Eleven", "Twelve"),
)`,
		b.Table(
			b.Narrow,
			b.HeadRow("One", "Two"),
			b.Row("Three", "Four"),
			b.Row("Five", "Six"),
			b.Row("Seven", "Eight"),
			b.Row("Nine", "Ten"),
			b.Row("Eleven", "Twelve"),
		),
	),
	c.Example(
		`b.Table(
	b.Hoverable,
	b.HeadRow("One", "Two"),
	b.Row("Three", "Four"),
	b.Row("Five", "Six"),
	b.Row("Seven", "Eight"),
	b.Row("Nine", "Ten"),
	b.Row("Eleven", "Twelve"),
)`,
		b.Table(
			b.Hoverable,
			b.HeadRow("One", "Two"),
			b.Row("Three", "Four"),
			b.Row("Five", "Six"),
			b.Row("Seven", "Eight"),
			b.Row("Nine", "Ten"),
			b.Row("Eleven", "Twelve"),
		),
	),
	c.Example(
		`b.Table(
	b.FullWidth,
	b.HeadRow("One", "Two"),
	b.Row("Three", "Four"),
	b.Row("Five", "Six"),
	b.Row("Seven", "Eight"),
	b.Row("Nine", "Ten"),
	b.Row("Eleven", "Twelve"),
)`,
		b.Table(
			b.FullWidth,
			b.HeadRow("One", "Two"),
			b.Row("Three", "Four"),
			b.Row("Five", "Six"),
			b.Row("Seven", "Eight"),
			b.Row("Nine", "Ten"),
			b.Row("Eleven", "Twelve"),
		),
	),
	c.Example(
		`b.Table(
	b.Bordered,
	b.Striped,
	b.Narrow,
	b.Hoverable,
	b.FullWidth,
	b.HeadRow("One", "Two"),
	b.Row("Three", "Four"),
	b.Row("Five", "Six"),
	b.Row("Seven", "Eight"),
	b.Row("Nine", "Ten"),
	b.Row("Eleven", "Twelve"),
)`,
		b.Table(
			b.Bordered,
			b.Striped,
			b.Narrow,
			b.Hoverable,
			b.FullWidth,
			b.HeadRow("One", "Two"),
			b.Row("Three", "Four"),
			b.Row("Five", "Six"),
			b.Row("Seven", "Eight"),
			b.Row("Nine", "Ten"),
			b.Row("Eleven", "Twelve"),
		),
	),
).Subsection(
	"Table container",
	"https://bulma.io/documentation/elements/table/#table-container",
	c.HorizontalExample(
		`b.ScrollableTable(
	// Your table content
)`,
		b.ScrollableTable(
			b.Bordered,
			b.Striped,
			b.Row("1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "100"),
			b.Row("2", "4", "6", "8", "10", "12", "14", "16", "18", "20", "22", "24", "26", "28", "30", "32", "34", "36", "38", "40", "42", "44", "46", "48", "50", "52", "54", "56", "58", "60", "62", "64", "66", "68", "70", "72", "74", "76", "78", "80", "82", "84", "86", "88", "90", "92", "94", "96", "98", "100", "102", "104", "106", "108", "110", "112", "114", "116", "118", "120", "122", "124", "126", "128", "130", "132", "134", "136", "138", "140", "142", "144", "146", "148", "150", "152", "154", "156", "158", "160", "162", "164", "166", "168", "170", "172", "174", "176", "178", "180", "182", "184", "186", "188", "190", "192", "194", "196", "198", "200"),
			b.Row("3", "6", "9", "12", "15", "18", "21", "24", "27", "30", "33", "36", "39", "42", "45", "48", "51", "54", "57", "60", "63", "66", "69", "72", "75", "78", "81", "84", "87", "90", "93", "96", "99", "102", "105", "108", "111", "114", "117", "120", "123", "126", "129", "132", "135", "138", "141", "144", "147", "150", "153", "156", "159", "162", "165", "168", "171", "174", "177", "180", "183", "186", "189", "192", "195", "198", "201", "204", "207", "210", "213", "216", "219", "222", "225", "228", "231", "234", "237", "240", "243", "246", "249", "252", "255", "258", "261", "264", "267", "270", "273", "276", "279", "282", "285", "288", "291", "294", "297", "300"),
			b.Row("4", "8", "12", "16", "20", "24", "28", "32", "36", "40", "44", "48", "52", "56", "60", "64", "68", "72", "76", "80", "84", "88", "92", "96", "100", "104", "108", "112", "116", "120", "124", "128", "132", "136", "140", "144", "148", "152", "156", "160", "164", "168", "172", "176", "180", "184", "188", "192", "196", "200", "204", "208", "212", "216", "220", "224", "228", "232", "236", "240", "244", "248", "252", "256", "260", "264", "268", "272", "276", "280", "284", "288", "292", "296", "300", "304", "308", "312", "316", "320", "324", "328", "332", "336", "340", "344", "348", "352", "356", "360", "364", "368", "372", "376", "380", "384", "388", "392", "396", "400"),
			b.Row("5", "10", "15", "20", "25", "30", "35", "40", "45", "50", "55", "60", "65", "70", "75", "80", "85", "90", "95", "100", "105", "110", "115", "120", "125", "130", "135", "140", "145", "150", "155", "160", "165", "170", "175", "180", "185", "190", "195", "200", "205", "210", "215", "220", "225", "230", "235", "240", "245", "250", "255", "260", "265", "270", "275", "280", "285", "290", "295", "300", "305", "310", "315", "320", "325", "330", "335", "340", "345", "350", "355", "360", "365", "370", "375", "380", "385", "390", "395", "400", "405", "410", "415", "420", "425", "430", "435", "440", "445", "450", "455", "460", "465", "470", "475", "480", "485", "490", "495", "500"),
		),
	),
)
