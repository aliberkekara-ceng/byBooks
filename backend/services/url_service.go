package services

import (
	"backend/models"
	"errors"
	"net/url"
	"strings"
)

type URLService interface {
	ProcessURL(req models.URLRequest) (string, error)
}

type urlService struct{}

func NewURLService() URLService {
	return &urlService{}
}

func (s *urlService) ProcessURL(req models.URLRequest) (string, error) {
	parsedURL, err := url.Parse(req.URL)
	if err != nil {
		return "", errors.New("invalid URL format")
	}

	if parsedURL.Scheme == "" {
		parsedURL, err = url.Parse("https://" + req.URL)
		if err != nil {
			return "", errors.New("invalid URL format")
		}
	}

	processedURL := parsedURL

	applyCanonical := func(u *url.URL) {
		u.RawQuery = ""
		u.Fragment = ""

		if strings.HasSuffix(u.Path, "/") && u.Path != "/" {
			u.Path = strings.TrimSuffix(u.Path, "/")
		}
	}

	applyRedirection := func(u *url.URL) {
		u.Host = "www.byfood.com"
	}

	switch req.Operation {
	case "canonical":
		applyCanonical(processedURL)
		return processedURL.String(), nil

	case "redirection":
		applyRedirection(processedURL)
		return strings.ToLower(processedURL.String()), nil

	case "all":
		applyCanonical(processedURL)
		applyRedirection(processedURL)
		return strings.ToLower(processedURL.String()), nil

	default:
		return "", errors.New("invalid operation type")
	}
}
