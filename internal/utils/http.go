package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"vibe/internal/version"
)

// FetchFromURL downloads content from a URL with proper timeout handling
func FetchFromURL(urlStr string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set a reasonable user agent
	req.Header.Set("User-Agent", "vibe/"+version.GetVersion())

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d %s", resp.StatusCode, resp.Status)
	}

	return io.ReadAll(resp.Body)
}
