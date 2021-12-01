package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlHandler(t *testing.T) {
	type want struct {
		// contentType string
		statusCode int
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
			request := httptest.NewRequest(http.MethodPost, tt.request, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UrlHandler)
			h.ServeHTTP(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
		})
	}
}
