package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/SSlavskii/go_url_shortener/internal/app/config"
	"github.com/SSlavskii/go_url_shortener/internal/app/storage"

	"github.com/labstack/echo/v4"
)

type APIShortenPayload struct {
	RawURL string `json:"url"`
}

type APIShortenResult struct {
	ShortURL string `json:"result"`
}

type Handler struct {
	storage storage.Storager
	config  config.Config
}

func New(s storage.Storager, cfg config.Config) *Handler {
	return &Handler{storage: s, config: cfg}
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
	e.Response().Header().Add("Content-Type", "text/plain")
	e.Response().Header().Add("Accept-Charset", "utf-8")
	return e.String(201, h.config.BaseURL+"/"+shortURL)
}

func (h *Handler) PostAPIShortenHandler(e echo.Context) error {
	defer e.Request().Body.Close()
	url := APIShortenPayload{}
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	if err := json.Unmarshal(body, &url); err != nil {
		return echo.NewHTTPError(400, "Error binding payload")
	}
	if url.RawURL == "" {
		return echo.NewHTTPError(400, "No url payload")
	}

	rawURL := url.RawURL
	shortURL, err := h.storage.GetIDFromFullURL(string(rawURL))

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	result := APIShortenResult{
		ShortURL: h.config.BaseURL + "/" + shortURL,
	}

	e.Response().Header().Add("Content-Type", echo.MIMEApplicationJSON)
	e.Response().Header().Add("Accept-Charset", "utf-8")
	return e.JSON(201, result)
}
