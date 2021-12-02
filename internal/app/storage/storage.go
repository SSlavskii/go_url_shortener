package storage

import (
	"errors"
	"strconv"
)

type Storage interface {
	GetFullURLFromID(urlID int) (string, error)
	GetIDFromFullURL(url string) (string, error)
}

type SimpleStorage struct {
	URLToInt map[string]int
	IntToURL []string
}

func (s *SimpleStorage) GetFullURLFromID(urlID int) (string, error) {
	if urlID >= len(s.IntToURL) {
		return "ya.ru", errors.New("NO such id")
	} else {
		return s.IntToURL[urlID], nil
	}
}

func (s *SimpleStorage) GetIDFromFullURL(rawURL string) (string, error) {
	shortInt, ok := s.URLToInt[rawURL]
	if !ok {
		shortInt = len(s.IntToURL)
		s.URLToInt[rawURL] = shortInt
		s.IntToURL = append(s.IntToURL, rawURL)
	}
	return strconv.Itoa(shortInt), nil

}

func New() *SimpleStorage {
	var s SimpleStorage
	s.URLToInt = make(map[string]int)
	s.IntToURL = make([]string, 0)
	return &s
}
