package main

import (
	"context"
	"embed"
	"errors"
	"net/http"
	"os"

	"dreampicai/config"
	"dreampicai/handler"
	"dreampicai/pkg/logger"
	"dreampicai/pkg/sb"
	"dreampicai/types"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//go:embed public
var FS embed.FS

func main() {
	log := logger.NewInfoJsonSlogLogger()

	ctx := context.Background()
	cfg, err := config.LoadEnvVars()
	if err != nil {
		log.Error(ctx, "Loading .env file fail", "err", err)
		os.Exit(1)
	}

	s := types.Server{
		Config: cfg,
		Logger: log,
	}

	if len(s.Config.Environment) == 0 {
		os.Exit(1)
	}

	if s.Config.Environment == "development" {
		s.Logger = logger.NewDebugJsonSlogLogger()
	}

	if err := sb.InitSupabaseClient(s); err != nil {
		s.Logger.Error(ctx, err.Error())
	}

	runDbMigration(ctx, s.Logger, s.Config.MigrationURL, s.Config.DbSource)

	router := chi.NewMux()
	router.Use(handler.WithLogger(s), handler.WithUser(s))
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(s, handler.HandleHomeIndex))
	router.Get("/login", handler.Make(s, handler.HandleLoginIndex))
	router.Get("/login/provider/google", handler.Make(s, handler.HandleLoginWithGoogle))
	router.Post("/login", handler.Make(s, func(w http.ResponseWriter, r *http.Request) error { return handler.HandleLoginCreate(s, w, r) }))
	router.Post("/logout", handler.Make(s, func(w http.ResponseWriter, r *http.Request) error { return handler.HandleLogoutCreate(s, w, r) }))
	router.Get("/signup", handler.Make(s, handler.HandleSignupIndex))
	router.Post("/signup", handler.Make(s, handler.HandleSignupCreate))
	router.Get("/auth/callback", handler.Make(s, func(w http.ResponseWriter, r *http.Request) error { return handler.HandleAuthCallback(s, w, r) }))

	router.Group(func(auth chi.Router) {
		auth.Use(handler.WithAuth)
		auth.Get("/settings", handler.Make(s, handler.HandleSettingsIndex))
	})

	s.Logger.Info(ctx, "application running on address", "address", s.Config.HttpListenAddr)
	err = http.ListenAndServe(s.Config.HttpListenAddr, router)
	if err != nil {
		s.Logger.Error(ctx, "fatal server error", "err", err)
		os.Exit(1)
	}
}

func runDbMigration(ctx context.Context, log types.Logger, migrationUrl string, dbSource string) {
	migration, err := migrate.New(migrationUrl, dbSource)
	if err != nil {
		log.Error(ctx, "cannot create new migrate instance", "err", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error(ctx, "failed to run migrate up", "err", err)
	} else if errors.Is(err, migrate.ErrNoChange) {
		log.Info(ctx, "no change in db migration")
	} else {
		log.Info(ctx, "db migrated successfully")
	}
}
