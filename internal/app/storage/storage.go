package storage

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Storager interface {
	GetFullURLFromID(urlID int) (string, error)
	GetIDFromFullURL(url string) (string, error)
}

type SimpleStorage struct {
	URLToInt    map[string]int
	IntToURL    []string
	StoragePath string
}

func (s *SimpleStorage) GetFullURLFromID(urlID int) (string, error) {
	if urlID >= len(s.IntToURL) {
		return "", errors.New("NO such id")
	}
	return s.IntToURL[urlID], nil
}

func (s *SimpleStorage) GetIDFromFullURL(rawURL string) (string, error) {
	shortInt, ok := s.URLToInt[rawURL]
	if !ok {
		shortInt = len(s.IntToURL)
		s.URLToInt[rawURL] = shortInt
		s.IntToURL = append(s.IntToURL, rawURL)
		if s.StoragePath != "" {
			file, err := os.OpenFile(s.StoragePath, os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				return "", err
			}
			defer file.Close()
			file.Write([]byte(strconv.Itoa(shortInt) + ";" + rawURL + "\n"))

		}
	}
	return strconv.Itoa(shortInt), nil

}

func New(storagePath string) (*SimpleStorage, error) {
	file, err := os.OpenFile(storagePath, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	urlToInt := make(map[string]int)
	intToURL := make([]string, 0)

	reader := bufio.NewReader(file)
	for {
		lineBytes, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		record := strings.Split(string(lineBytes), ";")
		urlID, _ := strconv.Atoi(record[0])
		urlToInt[record[1]] = urlID
		intToURL = append(intToURL, record[1])
	}

	return &SimpleStorage{
		URLToInt:    urlToInt,
		IntToURL:    intToURL,
		StoragePath: storagePath,
	}, nil
}
