package main

import (
	"github.com/SSlavskii/go_url_shortener/internal/app/handlers"
	"github.com/SSlavskii/go_url_shortener/internal/app/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	s := storage.New()
	h := handlers.New(s)

	e.GET("/:url_id", h.GetHandler)
	e.POST("/", h.PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
