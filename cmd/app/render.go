package main

import (
	"encoding/json"
	"io"

	"github.com/gobuffalo/buffalo/render"
)

// buffaloJSON renders a value as JSON without pulling in a template engine.
func buffaloJSON(v any) render.Renderer {
	return render.Func("application/json", func(w io.Writer, d render.Data) error {
		return json.NewEncoder(w).Encode(v)
	})
}
