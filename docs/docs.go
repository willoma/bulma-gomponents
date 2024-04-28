package docs

import (
	_ "embed"

	c "github.com/willoma/bulma-gomponents/docs/components"
)

//go:embed htmx.min.js
var HtmxJS []byte

var Sections = []c.DocSection{
	{
		Title: "",
		Pages: []*c.Page{
			intro,
			elements,
			classes,
			extending,
		},
	},
	{
		Title: "Elements",
		Pages: []*c.Page{
			block,
			box,
			button,
			content,
			delete,
			icon,
			image,
			notification,
			progress,
			table,
			tag,
			title,
		},
	},
	{
		Title: "Components",
		Pages: []*c.Page{
			breadcrumb,
			card,
			dropdown,
			menu,
			message,
			modal,
			navbar,
			pagination,
			panel,
			tabs,
		},
	},
	{
		Title: "Form",
		Pages: []*c.Page{
			formGeneral,
			formInput,
			formTextarea,
			formSelect,
			formCheckbox,
			formRadio,
			formFile,
		},
	},
	{
		Title: "Columns and grid",
		Pages: []*c.Page{
			columns,
			grid,
		},
	},
	{
		Title: "Layout",
		Pages: []*c.Page{
			container,
			hero,
			section,
			level,
			mediaObject,
			footer,
		},
	},
	{
		Title: "Helpers and features",
		Pages: []*c.Page{
			skeletons,
			color,
			positioning,
			spacing,
			typography,
			visibility,
			flexbox,
			other,
		},
	},
}
