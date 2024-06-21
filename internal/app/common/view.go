package common

import "github.com/a-h/templ"

type ViewContext struct {
	AppContext *AppContext

	Title   string
	Favicon string
	Metas   []templ.Attributes

	Links   []templ.Attributes
	Scripts []templ.Attributes
}

func DefaultViewContext(appContext *AppContext) *ViewContext {
	viewContext := &ViewContext{
		AppContext: appContext,

		Title:   "Lisfun -",
		Favicon: "",
		Metas: []templ.Attributes{
			{
				"charset": "UTF-8",
			}, {
				"name":    "viewport",
				"content": "width=device-width, initial-scale=1.0",
			}, {
				"name":    "description",
				"content": "hello",
			}, {
				"name":    "google",
				"content": "notranslate",
			},
		},

		Links: []templ.Attributes{
			{
				"rel":  "stylesheet",
				"type": "text/css",
				"href": "https://css.gg/css",
			}, {
				"ref":  "stylesheet",
				"type": "text/css",
				"href": "https://cdn.jsdelivr.net/npm/daisyui@4.4.10/dist/full.min.css",
			}, {
				"ref":  "stylesheet",
				"type": "text/css",
				"href": "/assets/css/main.css",
			},
		},
		Scripts: []templ.Attributes{
			{
				"type":        "text/javascript",
				"src":         "https://unpkg.com/htmx.org@1.9.9",
				"crossorigin": "anonymous",
				"integrity":   "sha384-QFjmbokDn2DjBjq+fM+8LUIVrAgqcNW2s0PjAxHETgRn9l4fvX31ZxDxvwQnyMOX",
			}, {
				"type": "text/javascript",
				"src":  "https://unpkg.com/htmx.org/dist/ext/json-enc.js",
			}, {
				"type": "text/javascript",
				"src":  "https://unpkg.com/htmx.org@1.9.12/dist/ext/response-targets.js",
			}, {
				"type": "text/javascript",
				"src":  "https://unpkg.com/hyperscript.org@0.9.12",
			},
		},
	}

	return viewContext
}
