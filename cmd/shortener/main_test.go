package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SSlavskii/go_url_shortener/internal/app/handlers"
	"github.com/SSlavskii/go_url_shortener/internal/app/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostHandler(t *testing.T) {
	type want struct {
		// contentType string
		statusCode int
		shortURL   string
	}
	tests := []struct {
		name    string
		request string
		want    want
		storage *storage.SimpleStorage
	}{
		{
			name: "simple test #1",
			storage: &storage.SimpleStorage{
				URLToInt: map[string]int{"ya.ru": 0},
				IntToURL: []string{"ya.ru"},
			},

			request: "/",
			want: want{
				statusCode: 201,
				shortURL:   "http://localhost:8080/0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			request := httptest.NewRequest(http.MethodPost, tt.request, strings.NewReader("ya.ru"))
			rec := httptest.NewRecorder()
			c := e.NewContext(request, rec)
			h := handlers.New(tt.storage)
			h.PostHandler(c)
			result := rec.Result()
			defer result.Body.Close()
			shortURL, _ := io.ReadAll(result.Body)

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.shortURL, string(shortURL))
		})
	}
}

func TestGetHandler(t *testing.T) {
	type want struct {
		// contentType string
		statusCode int
		location   string
	}
	tests := []struct {
		name    string
		request string
		want    want
		storage *storage.SimpleStorage
	}{
		{
			name:    "simple test #1",
			request: "/0",
			storage: &storage.SimpleStorage{
				URLToInt: map[string]int{"ya.ru": 0},
				IntToURL: []string{"ya.ru"},
			},
			want: want{
				statusCode: 200,
				location:   "ya.ru",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(request, rec)
			h := handlers.New(tt.storage)
			h.GetHandler(c)
			result := rec.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			//assert.Equal(t, tt.want.location, result.Header["Location"])
		})
	}
}
