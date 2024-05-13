package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"dreampicai/config"
	"dreampicai/handler"
	"dreampicai/pkg/logger"
	"dreampicai/pkg/sb"

	"github.com/go-chi/chi/v5"
)

//go:embed public
var FS embed.FS

func main() {
	ctx := context.Background()
	log := logger.NewInfoJsonLogger()

	cfg, err := config.LoadEnvVars()
	if err != nil {
		log.Error(ctx, "Loading .env file fail", "err", err)
		os.Exit(1)
	}

	if len(cfg.Environment) == 0 {
		os.Exit(1)
	}

	if cfg.Environment == "development" {
		log = logger.NewDebugJsonLogger()
	}

	if err := sb.InitSupabaseClient(); err != nil {
		log.Error(ctx, err.Error())
	}

	router := chi.NewMux()
	router.Use(handler.WithUser)
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(log, handler.HandleHomeIndex))
	router.Get("/login", handler.Make(log, handler.HandleLoginIndex))
	router.Get("/login/provider/google", handler.Make(log, handler.HandleLoginWithGoogle))
	router.Post("/login", handler.Make(log, handler.HandleLoginCreate))
	router.Post("/logout", handler.Make(log, handler.HandleLogoutCreate))
	router.Get("/signup", handler.Make(log, handler.HandleSignupIndex))
	router.Post("/signup", handler.Make(log, handler.HandleSignupCreate))
	router.Get("/auth/callback", handler.Make(log, handler.HandleAuthCallback))

	router.Group(func(auth chi.Router) {
		auth.Use(handler.WithAuth)
		auth.Get("/settings", handler.Make(log, handler.HandleSettingsIndex))
	})

	log.Info(ctx, "application running on address", "address", cfg.HttpListenAddr)
	err = http.ListenAndServe(cfg.HttpListenAddr, router)
	if err != nil {
		log.Error(ctx, "fatal server error", "err", err)
		os.Exit(1)
	}
}
