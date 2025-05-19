package unit

import (
	"testing"

	"github.com/DiegoDev2/Fleet/internal/core/manifest"
	pkgManifest "github.com/DiegoDev2/Fleet/pkg/manifest"
	"github.com/stretchr/testify/assert"
)

func TestValidateManifest(t *testing.T) {
	// Test valid manifest
	validManifest := &pkgManifest.Manifest{
		Name:    "test-tool",
		Version: "1.0.0",
		Platforms: map[string]pkgManifest.Platform{
			"darwin": {
				Architecture: map[string]pkgManifest.Architecture{
					"amd64": {
						URL:      "https://example.com/test.tar.gz",
						Type:     "tar.gz",
						Checksum: "sha256:abcdef",
						InstallSteps: []pkgManifest.Step{
							{Execute: "echo test"},
						},
					},
				},
			},
		},
	}

	errors := manifest.Validate(validManifest)
	assert.Empty(t, errors, "Valid manifest should have no validation errors")

	invalidManifest1 := &pkgManifest.Manifest{
		Version: "1.0.0",
		Platforms: map[string]pkgManifest.Platform{
			"darwin": {
				Architecture: map[string]pkgManifest.Architecture{
					"amd64": {
						URL:      "https://example.com/test.tar.gz",
						Type:     "tar.gz",
						Checksum: "sha256:abcdef",
						InstallSteps: []pkgManifest.Step{
							{Execute: "echo test"},
						},
					},
				},
			},
		},
	}

	errors = manifest.Validate(invalidManifest1)
	assert.NotEmpty(t, errors, "Invalid manifest should have validation errors")
	assert.Equal(t, "name", errors[0].Field, "Should detect missing name")

	invalidManifest2 := &pkgManifest.Manifest{
		Name: "test-tool",
		Platforms: map[string]pkgManifest.Platform{
			"darwin": {
				Architecture: map[string]pkgManifest.Architecture{
					"amd64": {
						URL:      "https://example.com/test.tar.gz",
						Type:     "tar.gz",
						Checksum: "sha256:abcdef",
						InstallSteps: []pkgManifest.Step{
							{Execute: "echo test"},
						},
					},
				},
			},
		},
	}

	errors = manifest.Validate(invalidManifest2)
	assert.NotEmpty(t, errors, "Invalid manifest should have validation errors")
	assert.Equal(t, "version", errors[0].Field, "Should detect missing version")
}
