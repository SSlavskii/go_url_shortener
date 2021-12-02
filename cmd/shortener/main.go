package main

import (
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
)

var urlToInt = make(map[string]int)
var intToURL = make([]string, 0)

func GetHandler(e echo.Context) error {
	index, err := strconv.Atoi(path.Base(e.Param("url_id")))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	} else if index >= len(intToURL) {
		return echo.NewHTTPError(400, "NO such id")
	} else {
		e.Response().Header().Set(echo.HeaderLocation, intToURL[index])
		return e.String(http.StatusTemporaryRedirect, "")
	}
}

func PostHandler(e echo.Context) error {
	defer e.Request().Body.Close()
	rawURL, err := io.ReadAll(e.Request().Body)
	shortInt, ok := urlToInt[string(rawURL)]
	if !ok {
		shortInt = len(intToURL)
		urlToInt[string(rawURL)] = shortInt
		intToURL = append(intToURL, string(rawURL))
	}
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	e.Response().Header().Add("Content-Type", "application/json")
	e.Response().Header().Add("Accept-Charset", "utf-8")
	e.Response().WriteHeader(201)
	return e.String(201, "http://localhost:8080/"+strconv.Itoa(shortInt))
}

func main() {
	e := echo.New()

	e.GET("/:url_id", GetHandler)
	e.POST("/", PostHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
