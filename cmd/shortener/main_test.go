package main

import (
	"net/http"
	"net/http/httptest"
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
	}
	simpleStorage := &storage.SimpleStorage{
		URLToInt: map[string]int{"ya.ru": 0},
		IntToURL: []string{"ya.ru"},
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "simple test #1",
			request: "/",
			want: want{
				statusCode: 201,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			request := httptest.NewRequest(http.MethodPost, tt.request, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(request, rec)
			h := handlers.New(simpleStorage)
			h.PostHandler(c)
			result := rec.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}

func TestGetHandler(t *testing.T) {
	type want struct {
		// contentType string
		statusCode int
	}
	simpleStorage := &storage.SimpleStorage{
		URLToInt: map[string]int{"ya.ru": 0},
		IntToURL: []string{"ya.ru"},
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "simple test #1",
			request: "/0",
			want: want{
				statusCode: 200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(request, rec)
			h := handlers.New(simpleStorage)
			h.GetHandler(c)
			result := rec.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}
