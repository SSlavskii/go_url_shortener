package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SSlavskii/go_url_shortener/internal/app/config"

	"github.com/SSlavskii/go_url_shortener/internal/app/handlers"
	"github.com/SSlavskii/go_url_shortener/internal/app/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var testConfig = config.Config{
	ServerAdress: ":8080",
	BaseURL:      "http://localhost:8080",
}

func TestPostHandler(t *testing.T) {
	type want struct {
		// contentType string
		statusCode int
		shortURL   string
	}
	tests := []struct {
		name    string
		want    want
		storage *storage.SimpleStorage
	}{
		{
			name: "simple test #1",
			storage: &storage.SimpleStorage{
				URLToInt: map[string]int{"ya.ru": 0},
				IntToURL: []string{"ya.ru"},
			},
			want: want{
				statusCode: 201,
				shortURL:   "http://localhost:8080/0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("ya.ru"))
			rec := httptest.NewRecorder()
			c := e.NewContext(request, rec)
			h := handlers.New(tt.storage, testConfig)
			h.PostHandler(c)
			result := rec.Result()
			defer result.Body.Close()
			shortURL, _ := io.ReadAll(result.Body)

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.shortURL, string(shortURL))
		})
	}
}

func TestPostApiShortenHandler(t *testing.T) {
	type want struct {
		// contentType string
		statusCode int
		result     handlers.APIShortenResult
	}
	tests := []struct {
		name    string
		want    want
		body    handlers.APIShortenPayload
		storage *storage.SimpleStorage
	}{
		{
			name: "simple test #1",
			storage: &storage.SimpleStorage{
				URLToInt: map[string]int{"ya.ru": 0},
				IntToURL: []string{"ya.ru"},
			},

			body: handlers.APIShortenPayload{RawURL: "ya.ru"},
			want: want{
				statusCode: 201,
				result: handlers.APIShortenResult{
					ShortURL: "http://localhost:8080/0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			body, _ := json.Marshal(tt.body)
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(request, rec)
			h := handlers.New(tt.storage, testConfig)
			h.PostAPIShortenHandler(c)
			result := rec.Result()
			defer result.Body.Close()
			var resultPayload handlers.APIShortenResult
			err := json.NewDecoder(result.Body).Decode(&resultPayload)

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.NoError(t, err, "binding response error")
			assert.Equal(t, tt.want.result, resultPayload)
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
			h := handlers.New(tt.storage, testConfig)
			h.GetHandler(c)
			result := rec.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			//assert.Equal(t, tt.want.location, result.Header["Location"])
		})
	}
}
