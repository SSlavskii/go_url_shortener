package handlers

import (
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/SSlavskii/go_url_shortener/internal/app/storage"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	storage storage.Storage
}

func New(s storage.Storage) Handler {
	return Handler{storage: s}
}

func (h *Handler) GetHandler(e echo.Context) error {
	index, err := strconv.Atoi(path.Base(e.Param("url_id")))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	url, err := h.storage.GetFullURLFromID(index)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	} else {
		e.Response().Header().Set(echo.HeaderLocation, url)
		return e.String(http.StatusTemporaryRedirect, "")
	}
}

func (h *Handler) PostHandler(e echo.Context) error {
	defer e.Request().Body.Close()
	rawURL, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	shortURL, err := h.storage.GetIDFromFullURL(string(rawURL))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	e.Response().Header().Add("Content-Type", "application/json")
	e.Response().Header().Add("Accept-Charset", "utf-8")
	e.Response().WriteHeader(201)
	return e.String(201, "http://localhost:8080/"+shortURL)
}
