// A minimal Buffalo app shaped for the fleet.
package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gobuffalo/buffalo"
)

func basePath() string {
	raw := strings.Trim(strings.TrimSpace(os.Getenv("BASE_PATH")), "/")
	if raw == "" {
		return ""
	}
	return "/" + raw
}

func port() string {
	if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
		return p
	}
	return "8080"
}

func newApp() *buffalo.App {
	app := buffalo.New(buffalo.Options{
		Env:  "production",
		Addr: ":" + port(),
	})

	// Buffalo mounts a group at a prefix; with BASE_PATH empty the group is the
	// app root, so the same code serves standalone.
	g := app.Group("/")
	if bp := basePath(); bp != "" {
		g = app.Group(bp)
	}
	g.GET("/health", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, buffaloJSON(map[string]string{"status": "ok"}))
	})
	return app
}

func main() {
	log.Printf("buffalo-template listening on :%s (base_path=%q)", port(), basePath())
	if err := newApp().Serve(); err != nil {
		log.Fatal(err)
	}
}
