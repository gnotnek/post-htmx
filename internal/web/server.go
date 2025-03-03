package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"post-htmx/internal/renderer"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	// init render
	renderer := renderer.NewRenderer("internal/web/templates")
	r := chi.NewRouter()

	// Web routes
	r.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/templates/public"))).ServeHTTP)
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, r, "layout.html", map[string]interface{}{
			"Title": "Home",
		})
	})
	return &Server{router: r}
}

func (s *Server) Run(port int) {
	addr := fmt.Sprintf(":%d", port)

	h := chainMiddleware(
		s.router,
		recoverHandler,
		loggerHandler(func(w http.ResponseWriter, r *http.Request) bool { return r.URL.Path == "/" }),
		realIPHandler,
		requestIDHandler,
		corsHandler,
	)

	httpServer := http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Info().Msg("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("Server forced to shutdown")
		}

		close(done)
	}()

	log.Info().Msgf("server serving on port %d", port)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}

	<-done
	log.Info().Msg("Server stopped")
}
