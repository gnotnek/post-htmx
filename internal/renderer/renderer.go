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

func (r *Renderer) Render(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := r.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Error().Err(err).Msgf("failed to render template: %s", name)
		resp.WriteError(w, err)
	} else {
		log.Info().Msgf("successfully rendered template: %s", name)
		w.WriteHeader(http.StatusOK)
	}
}
