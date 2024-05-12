package main

import (
	"embed"
	"log/slog"
	"net/http"
	"os"

	"dreampicai/config"
	"dreampicai/handler"
	"dreampicai/pkg/sb"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

var log = *slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	config, err := config.LoadConfig(log)
	if err != nil {
		log.Error("Loading config failed.", err)
		os.Exit(1)
	}

	if err := initEverything(); err != nil {
		log.Error(err.Error())
	}

	router := chi.NewMux()
	router.Use(handler.WithUser)
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))
	router.Get("/login/provider/google", handler.Make(handler.HandleLoginWithGoogle))
	router.Post("/login", handler.Make(handler.HandleLoginCreate))
	router.Post("/logout", handler.Make(handler.HandleLogoutCreate))
	router.Get("/signup", handler.Make(handler.HandleSignupIndex))
	router.Post("/signup", handler.Make(handler.HandleSignupCreate))
	router.Get("/auth/callback", handler.Make(handler.HandleAuthCallback))

	router.Group(func(auth chi.Router) {
		auth.Use(handler.WithAuth)
		auth.Get("/settings", handler.Make(handler.HandleSettingsIndex))
	})

	log.Info("application running", "addr", config.HttpListenAddr)
	err = http.ListenAndServe(config.HttpListenAddr, router)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return sb.Init()
}
