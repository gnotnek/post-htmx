package renderer

import (
	"html/template"
	"net/http"
	"path/filepath"
	"post-htmx/internal/api/resp"

	"github.com/rs/zerolog/log"
)

type Renderer struct {
	templates *template.Template
}

func NewRenderer(templateDir string) *Renderer {
	tmpl := template.Must(template.ParseFiles(
		filepath.Join(templateDir, "layout.html"),
		filepath.Join(templateDir, "index.html"),
	))
	return &Renderer{
		templates: tmpl,
	}
}

func (R *Renderer) Render(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := R.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Error().Err(err).Msgf("failed to render template: %s", name)
		resp.WriteError(w, err)
	} else {
		log.Info().Msgf("successfully rendered template: %s", name)
	}
}

func (R *Renderer) TestRender(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("internal/templates/layout.html", "internal/templates/index.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Title": "Home",
	})
	if err != nil {
		log.Error().Err(err).Msgf("failed to render template: %s", "layout.html")
		resp.WriteError(w, err)
	} else {
		log.Info().Msgf("successfully rendered template: %s", "layout.html")
	}
}
