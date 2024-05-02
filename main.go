package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"dreampicai/handler"
	"dreampicai/pkg/sb"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Use(handler.WithUser)
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))
	router.Post("/login", handler.Make(handler.HandleLoginCreate))

	addr := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "addr", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return sb.Init()
}
