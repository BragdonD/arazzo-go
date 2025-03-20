package arazzo

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Loader handles file loading with configurable options.
type Loader struct {
	httpClient  *http.Client
	allowRemote bool
	allowLocal  bool
}

// Option defines a functional option for configuring Loader.
type Option func(*Loader)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(l *Loader) {
		l.httpClient = client
	}
}

// AllowRemoteLookup enables remote file loading.
func AllowRemoteLookup() Option {
	return func(l *Loader) {
		l.allowRemote = true
	}
}

// AllowLocalLookup enables local file loading.
func AllowLocalLookup() Option {
	return func(l *Loader) {
		l.allowLocal = true
	}
}

// NewLoader creates a new Loader with optional configurations.
func NewLoader(opts ...Option) *Loader {
	l := &Loader{
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		allowRemote: false,
		allowLocal:  false,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// IsRemoteFile checks if the given URL is remote.
func IsRemoteFile(rawURL string) bool {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return parsedURL.Scheme != "" && parsedURL.Host != ""
}

// LoadFile loads a file, supporting both local and remote paths.
func (l *Loader) LoadFile(path string) ([]byte, error) {
	if IsRemoteFile(path) {
		if !l.allowRemote {
			return nil, fmt.Errorf("remote file lookup is not allowed")
		}
		return l.loadRemoteFile(path)
	}
	if !l.allowLocal {
		return nil, fmt.Errorf("local file lookup is not allowed")
	}
	return l.loadLocalFile(path)
}

// loadLocalFile reads a file from the local filesystem.
func (l *Loader) loadLocalFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read local file: %w", err)
	}
	return data, nil
}

// loadRemoteFile fetches a file from a remote URL using the configured HTTP client.
func (l *Loader) loadRemoteFile(rawURL string) ([]byte, error) {
	resp, err := l.httpClient.Get(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch remote file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read remote file: %w", err)
	}
	return data, nil
}
