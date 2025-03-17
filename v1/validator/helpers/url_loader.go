package helpers

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// FileScheme represents the "file" URL scheme for local file access.
const (
	FileScheme  = "file"
	HttpsScheme = "https"
)

type URLLoader interface {
	Load(url string) (any, error)
}

type HTTPURLLoader struct {
	httpClient  *http.Client
	httpsClient *http.Client
}

type HTTPURLLoaderOption func(*HTTPURLLoader)

func WithHTTPInsecureSkipVerify() HTTPURLLoaderOption {
	return func(h *HTTPURLLoader) {
		h.httpsClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
}

func WithHTTPSInsecureSkipVerify() HTTPURLLoaderOption {
	return func(h *HTTPURLLoader) {
		h.httpsClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
}

func (l *HTTPURLLoader) Load(rawUrl string) (any, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}
	var reader io.Reader

	// Check if the URL represents a local file.
	// A local file is identified by:
	// - Having no scheme (e.g., "myfile.yaml",
	// "/home/user/myfile.yaml")
	// - Using the "file" scheme (e.g., "file:///home/user/file.yaml")
	if parsedURL.Scheme == "" || parsedURL.Scheme == FileScheme {
		path := filepath.FromSlash(parsedURL.Path)
		reader, err = os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %v", err)
		}
	} else {
		var client *http.Client
		if strings.HasPrefix(rawUrl, HttpsScheme) {
			client = (*http.Client)(l.httpsClient)
		} else {
			client = (*http.Client)(l.httpClient)
		}

		resp, err := client.Get(rawUrl)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			_ = resp.Body.Close()
			return nil, fmt.Errorf("%s returned status code %d", rawUrl, resp.StatusCode)
		}
		defer resp.Body.Close()
		reader = resp.Body
	}

	return jsonschema.UnmarshalJSON(reader)
}

func NewHTTPURLLoader(opts ...HTTPURLLoaderOption) *HTTPURLLoader {
	httpLoader := &HTTPURLLoader{
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		httpsClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(httpLoader)
	}
	return httpLoader
}

func NewCompilerLoader(opts ...HTTPURLLoaderOption) jsonschema.SchemeURLLoader {
	return jsonschema.SchemeURLLoader{
		"file":  jsonschema.FileLoader{},
		"http":  NewHTTPURLLoader(opts...),
		"https": NewHTTPURLLoader(opts...),
	}
}
