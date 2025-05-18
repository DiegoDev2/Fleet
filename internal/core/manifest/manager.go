package manifest

import (
	"fmt"

	"github.com/DiegoDev2/Fleet/pkg/manifest"
)

type Manager struct {
	manifests map[string]*manifest.Manifest
}

func NewManager() *Manager {
	return &Manager{
		manifests: make(map[string]*manifest.Manifest),
	}
}

func (m *Manager) LoadManifestsFromDirectory(dirPath string) error {
	manifests, err := ParseDirectory(dirPath)
	if err != nil {
		return fmt.Errorf("error al cargar manifiestos: %w", err)
	}

	for name, manifest := range manifests {
		errors := Validate(manifest)
		if len(errors) > 0 {
			errorMessages := ""
			for _, err := range errors {
				errorMessages += fmt.Sprintf("\n  - %s", err.Error())
			}
			return fmt.Errorf("manifiesto inválido para %s:%s", name, errorMessages)
		}

		m.manifests[name] = manifest
	}

	return nil
}

func (m *Manager) GetManifest(name string) (*manifest.Manifest, bool) {
	manifest, exists := m.manifests[name]
	return manifest, exists
}

func (m *Manager) ListManifests() map[string]*manifest.Manifest {
	return m.manifests
}

func (m *Manager) AddManifest(manifest *manifest.Manifest) error {

	errors := Validate(manifest)
	if len(errors) > 0 {
		errorMessages := ""
		for _, err := range errors {
			errorMessages += fmt.Sprintf("\n  - %s", err.Error())
		}
		return fmt.Errorf("manifiesto inválido para %s:%s", manifest.Name, errorMessages)
	}

	m.manifests[manifest.Name] = manifest
	return nil
}

func (m *Manager) RemoveManifest(name string) bool {
	_, exists := m.manifests[name]
	if exists {
		delete(m.manifests, name)
		return true
	}
	return false
}

func (m *Manager) GetManifestsByCategory(category string) []*manifest.Manifest {
	var result []*manifest.Manifest

	for _, manifest := range m.manifests {
		for _, cat := range manifest.Categories {
			if cat == category {
				result = append(result, manifest)
				break
			}
		}
	}

	return result
}
