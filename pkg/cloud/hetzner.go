package cloud

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"gopkg.in/yaml.v3"
)

// HetznerMetadataURL is the link-local endpoint of the Hetzner metadata service.
// https://docs.hetzner.cloud/reference/cloud#description/server-metadata
const HetznerMetadataURL = "http://169.254.169.254/hetzner/v1/metadata"

// hetznerTimeout bounds the metadata request.
// The service is link-local and should normally answer in milliseconds, but
// we need to ensure we don't hang the scan when the flag is enabled on a
// machine that isn't a Hetzner instance.
const hetznerTimeout = 5 * time.Second

// hetznerMaxResponseSize bounds the response body we are prepared to read.
const hetznerMaxResponseSize = 1 << 20

// HetznerMetadata fetches instance metadata from the given URL and converts
// the YAML response into a JSON-serialisable map.
func HetznerMetadata(ctx context.Context, url string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(ctx, hetznerTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create metadata request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hetzner metadata: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch hetzner metadata: unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, hetznerMaxResponseSize))
	if err != nil {
		return nil, fmt.Errorf("failed to read hetzner metadata response: %w", err)
	}

	var metadata map[string]any

	if err = yaml.Unmarshal(body, &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse hetzner metadata as yaml: %w", err)
	}

	return metadata, nil
}
