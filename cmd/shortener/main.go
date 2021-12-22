package main

import (
	"github.com/SSlavskii/go_url_shortener/internal/app/config"
	"github.com/SSlavskii/go_url_shortener/internal/app/handlers"
	"github.com/SSlavskii/go_url_shortener/internal/app/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, _ := config.New()
	e := echo.New()
	s, _ := storage.New(cfg.FileStoragePath)
	h := handlers.New(s, *cfg)

	e.GET("/:url_id", h.GetHandler)
	e.POST("/", h.PostHandler)
	e.POST("/api/shorten", h.PostAPIShortenHandler)
	e.Logger.Fatal(e.Start(cfg.ServerAdress))
}
