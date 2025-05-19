package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/DiegoDev2/Fleet/pkg/manifest"
)

type Manager struct {
	repositories map[string]Repository
	cacheDir     string
	cache        *Cache
	mutex        sync.RWMutex
}

func NewManager(cacheDir string) (*Manager, error) {

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("error al crear directorio de caché: %w", err)
	}

	cache, err := NewCache(cacheDir)
	if err != nil {
		return nil, fmt.Errorf("error al inicializar caché: %w", err)
	}

	return &Manager{
		repositories: make(map[string]Repository),
		cacheDir:     cacheDir,
		cache:        cache,
	}, nil
}

func (m *Manager) AddRepository(name, url, repoType string, priority int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.repositories[name]; exists {
		return fmt.Errorf("ya existe un repositorio con el nombre %s", name)
	}

	repoCacheDir := filepath.Join(m.cacheDir, "repos", name)
	if err := os.MkdirAll(repoCacheDir, 0755); err != nil {
		return fmt.Errorf("error al crear directorio para el repositorio: %w", err)
	}

	var repo Repository
	var err error
	switch repoType {
	case "git":
		repo, err = NewGitRepository(name, url, repoCacheDir, priority)
	case "http":
		repo, err = NewHTTPRepository(name, url, repoCacheDir, priority)
	default:
		return fmt.Errorf("tipo de repositorio no soportado: %s", repoType)
	}

	if err != nil {
		return err
	}

	m.repositories[name] = repo

	m.cache.AddRepo(name, url, repoType)
	if err := m.cache.SaveMetadata(); err != nil {
		fmt.Printf("Warning: Failed to save cache metadata: %v\n", err)
	}

	return nil
}

func (m *Manager) RemoveRepository(name string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.repositories[name]; !exists {
		return fmt.Errorf("no existe un repositorio con el nombre %s", name)
	}

	delete(m.repositories, name)

	repoCacheDir := filepath.Join(m.cacheDir, "repos", name)
	if err := os.RemoveAll(repoCacheDir); err != nil {
		return fmt.Errorf("error al eliminar directorio del repositorio: %w", err)
	}

	m.cache.RemoveRepo(name)
	if err := m.cache.SaveMetadata(); err != nil {
		fmt.Printf("Warning: Failed to save cache metadata: %v\n", err)
	}

	return nil
}

func (m *Manager) SyncRepository(name string) error {
	m.mutex.RLock()
	repo, exists := m.repositories[name]
	m.mutex.RUnlock()

	if !exists {
		return fmt.Errorf("repositorio no encontrado: %s", name)
	}

	if err := repo.Sync(); err != nil {
		return err
	}

	manifests, err := repo.ListManifests()
	if err != nil {
		return fmt.Errorf("error al listar manifiestos: %w", err)
	}

	for _, mf := range manifests {

		checksum := fmt.Sprintf("sha256:%p", mf) // this is a placeholder, you should use a real checksum

		manifestData, err := manifest.ToYAML(mf)
		if err != nil {
			fmt.Printf("Warning: Failed to convert manifest to YAML: %v\n", err)
			continue
		}

		if err := m.cache.AddManifest(mf.Name, name, checksum, manifestData); err != nil {
			fmt.Printf("Warning: Failed to add manifest to cache: %v\n", err)
		}
	}

	if err := m.cache.SaveMetadata(); err != nil {
		fmt.Printf("Warning: Failed to save cache metadata: %v\n", err)
	}

	return nil
}

func (m *Manager) SyncAllRepositories() error {
	m.mutex.RLock()
	repos := make([]Repository, 0, len(m.repositories))
	for _, repo := range m.repositories {
		repos = append(repos, repo)
	}
	m.mutex.RUnlock()

	var wg sync.WaitGroup
	errCh := make(chan error, len(repos))

	for _, repo := range repos {
		wg.Add(1)
		go func(r Repository) {
			defer wg.Done()
			if err := m.SyncRepository(r.GetName()); err != nil {
				errCh <- fmt.Errorf("error al sincronizar %s: %w", r.GetName(), err)
			}
		}(repo)
	}

	wg.Wait()
	close(errCh)

	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("errores al sincronizar repositorios: %v", errors)
	}

	return nil
}

func (m *Manager) GetManifest(toolName string) (*manifest.Manifest, error) {

	cachedManifest, err := m.cache.GetManifest(toolName)
	if err == nil {
		return cachedManifest, nil
	}

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	repos := make([]Repository, 0, len(m.repositories))
	for _, repo := range m.repositories {
		repos = append(repos, repo)
	}
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].GetPriority() > repos[j].GetPriority()
	})

	for _, repo := range repos {
		manifest, err := repo.GetManifest(toolName)
		if err == nil && manifest != nil {
			// Guardar en caché para futuras consultas
			manifestData, err := manifest.ToYAML(manifest)
			if err == nil {
				checksum := fmt.Sprintf("sha256:%p", manifest) // Placeholder
				if data, ok := manifestData.([]byte); ok {
					m.cache.AddManifest(toolName, repo.GetName(), checksum, data)
				} else {
					fmt.Printf("Warning: manifestData is not of type []byte\n")
				}
				m.cache.SaveMetadata()
			}
			return manifest, nil
		}
	}

	return nil, fmt.Errorf("manifiesto no encontrado para la herramienta: %s", toolName)
}

// SearchManifests busca manifiestos que coincidan con un patrón
func (m *Manager) SearchManifests(pattern string) ([]*manifest.Manifest, error) {
	allManifests, err := m.ListAllManifests()
	if err != nil {
		if e, ok := err.(error); ok {
			return nil, e
		}
		return nil, fmt.Errorf("unexpected error type: %v", err)
	}

	var results []*manifest.Manifest
	manifests, ok := allManifests.([]*manifest.Manifest)
	if !ok {
		return nil, fmt.Errorf("unexpected type for allManifests, expected []*manifest.Manifest")
	}
	for _, m := range manifests {

		if strings.Contains(strings.ToLower(m.Name), strings.ToLower(pattern)) ||
			strings.Contains(strings.ToLower(m.Description), strings.ToLower(pattern)) {
			results = append(results, m)
			continue
		}

		for _, cat := range m.Categories {
			if strings.Contains(strings.ToLower(cat), strings.ToLower(pattern)) {
				results = append(results, m)
				break
			}
		}
	}

	return results, nil
}

func (m *Manager) ListAllManifests() (any, any) {
	panic("unimplemented")
}

func (m *Manager) GetRepositoryByName(name string) (Repository, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	repo, exists := m.repositories[name]
	return repo, exists
}

func (m *Manager) ClearCache() error {
	return m.cache.ClearCache()
}
