package manifest

import (
	"fmt"
	"strings"

	"github.com/DiegoDev2/Fleet/pkg/manifest"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string // The field where the validation error occurred
	Message string // The error message describing the issue
}

// Error implements the error interface for ValidationError
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Validate validates a manifest object and returns a list of validation errors
func Validate(m *manifest.Manifest) []ValidationError {
	var errors []ValidationError

	// Check if the manifest name is provided
	if m.Name == "" {
		errors = append(errors, ValidationError{Field: "name", Message: "is required"})
	}

	// Check if the manifest version is provided
	if m.Version == "" {
		errors = append(errors, ValidationError{Field: "version", Message: "is required"})
	}

	// Check if at least one platform is defined
	if len(m.Platforms) == 0 {
		errors = append(errors, ValidationError{Field: "platforms", Message: "at least one platform is required"})
	}

	// Iterate through each platform and validate its data
	for platform, platformData := range m.Platforms {
		// Ensure that at least one of architecture, package managers, or fallback is defined
		if len(platformData.Architecture) == 0 && len(platformData.PackageManagers) == 0 && platformData.Fallback == nil {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("platforms.%s", platform),
				Message: "must specify at least one architecture, package manager, or fallback",
			})
		}

		// Validate each architecture within the platform
		for arch, archData := range platformData.Architecture {
			// Check if the URL is provided
			if archData.URL == "" {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("platforms.%s.architecture.%s.url", platform, arch),
					Message: "is required",
				})
			}

			// Check if the checksum is provided and properly formatted
			if archData.Checksum == "" {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("platforms.%s.architecture.%s.checksum", platform, arch),
					Message: "is required",
				})
			} else if !strings.Contains(archData.Checksum, ":") {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("platforms.%s.architecture.%s.checksum", platform, arch),
					Message: "must be in the format 'algorithm:hash'",
				})
			}

			// Check if at least one installation step is defined
			if len(archData.InstallSteps) == 0 {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("platforms.%s.architecture.%s.install_steps", platform, arch),
					Message: "at least one installation step is required",
				})
			}
		}

		// Validate fallback data if it exists
		if platformData.Fallback != nil {
			// Check if the fallback URL is provided
			if platformData.Fallback.URL == "" {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("platforms.%s.fallback.url", platform),
					Message: "is required",
				})
			}

			// Check if the fallback checksum is provided
			if platformData.Fallback.Checksum == "" {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("platforms.%s.fallback.checksum", platform),
					Message: "is required",
				})
			}

			// Check if at least one build step is defined for the fallback
			if len(platformData.Fallback.BuildSteps) == 0 {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("platforms.%s.fallback.build_steps", platform),
					Message: "at least one build step is required",
				})
			}
		}
	}

	// Return the list of validation errors
	return errors
}
