package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"dreampicai/handler"

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

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	addr := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "addr", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func initEverything() error {
	return godotenv.Load()
}
