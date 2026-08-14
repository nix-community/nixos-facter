// Package cloud captures instance metadata from cloud providers for inclusion in the report.
package cloud

// Cloud groups metadata captured from cloud provider metadata services.
type Cloud struct {
	// Hetzner holds instance metadata fetched from the Hetzner metadata service,
	// converted from its YAML representation.
	Hetzner map[string]any `json:"hetzner,omitempty"`
}
