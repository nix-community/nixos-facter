package cloud_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/numtide/nixos-facter/pkg/cloud"
	"github.com/stretchr/testify/require"
)

// metadata document as served by http://169.254.169.254/hetzner/v1/metadata,
// composed from the sample values used in hcloud-go's metadata client tests:
// https://github.com/hetznercloud/hcloud-go/blob/main/hcloud/metadata/client_test.go
const hetznerMetadataFixture = `availability-zone: fsn1-dc14
hostname: my-server
instance-id: 123456
local-ipv4: ""
public-ipv4: 127.0.0.1
region: eu-central
public-keys: []
private-networks:
    - ip: 10.0.0.2
      alias_ips: [10.0.0.3, 10.0.0.4]
      interface_num: 1
      mac_address: 86:00:00:2a:7d:e0
      network_id: 1234
      network_name: nw-test1
      network: 10.0.0.0/8
      subnet: 10.0.0.0/24
      gateway: 10.0.0.1
    - ip: 192.168.0.2
      alias_ips: []
      interface_num: 2
      mac_address: 86:00:00:2a:7d:e1
      network_id: 4321
      network_name: nw-test2
      network: 192.168.0.0/16
      subnet: 192.168.0.0/24
      gateway: 192.168.0.1
`

func TestHetznerMetadata(t *testing.T) {
	t.Parallel()

	rq := require.New(t)

	var requestedPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPath = r.URL.Path
		_, _ = w.Write([]byte(hetznerMetadataFixture))
	}))
	defer server.Close()

	metadata, err := cloud.HetznerMetadata(t.Context(), server.URL+"/hetzner/v1/metadata")
	rq.NoError(err)
	rq.Equal("/hetzner/v1/metadata", requestedPath)

	// scalar fields survive the yaml decode
	rq.Equal(123456, metadata["instance-id"])
	rq.Equal("my-server", metadata["hostname"])
	rq.Equal("eu-central", metadata["region"])
	rq.Equal("fsn1-dc14", metadata["availability-zone"])
	rq.Equal("127.0.0.1", metadata["public-ipv4"])

	// nested structures are decoded, not flattened
	networks, ok := metadata["private-networks"].([]any)
	rq.True(ok, "private-networks should decode to a list")
	rq.Len(networks, 2)

	first, ok := networks[0].(map[string]any)
	rq.True(ok, "private network entries should decode to maps")
	rq.Equal("86:00:00:2a:7d:e0", first["mac_address"])
	rq.Equal("10.0.0.1", first["gateway"])

	// the whole document must serialise cleanly to JSON for the report
	data, err := json.Marshal(cloud.Cloud{Hetzner: metadata})
	rq.NoError(err)
	rq.Contains(string(data), `"mac_address":"86:00:00:2a:7d:e0"`)
	rq.Contains(string(data), `"alias_ips":["10.0.0.3","10.0.0.4"]`)
}

func TestHetznerMetadataUnavailable(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, err := cloud.HetznerMetadata(t.Context(), server.URL+"/hetzner/v1/metadata")
	require.Error(t, err)
}

func TestHetznerMetadataInvalidYaml(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("\t not: valid: yaml"))
	}))
	defer server.Close()

	_, err := cloud.HetznerMetadata(t.Context(), server.URL+"/hetzner/v1/metadata")
	require.Error(t, err)
}
