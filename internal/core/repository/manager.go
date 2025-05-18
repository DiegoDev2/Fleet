package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/DiegoDev2/Fleet/pkg/manifest"
	"github.com/DiegoDev2/Fleet/pkg/repository"
)

type Manager struct {
	repositories map[string]Repository
	cacheDir     string
	mutex        sync.RWMutex
}

// NewManager crea un nuevo gestor de repositorios
func NewManager(cacheDir string) (*Manager, error) {
	// Asegurar que el directorio de caché exista
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("error al crear directorio de caché: %w", err)
	}

	return &Manager{
		repositories: make(map[string]Repository),
		cacheDir:     cacheDir,
	}, nil
}

func (m *Manager) AddRepository(name, url, repoType string, priority int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.repositories[name]; exists {
		return fmt.Errorf("ya existe un repositorio con el nombre %s", name)
	}

	// Crear el directorio para el repositorio en la caché
	repoCacheDir := filepath.Join(m.cacheDir, name)
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
	return nil
}

func (m *Manager) RemoveRepository(name string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.repositories[name]; !exists {
		return fmt.Errorf("no existe un repositorio con el nombre %s", name)
	}

	delete(m.repositories, name)

	repoCacheDir := filepath.Join(m.cacheDir, name)
	if err := os.RemoveAll(repoCacheDir); err != nil {
		return fmt.Errorf("error al eliminar directorio del repositorio: %w", err)
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

	return repo.Sync()
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
			if err := r.Sync(); err != nil {
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
			return manifest, nil
		}
	}

	return nil, fmt.Errorf("manifiesto no encontrado para la herramienta: %s", toolName)
}

func (m *Manager) ListRepositories() []repository.RepoInfo {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	repos := make([]repository.RepoInfo, 0, len(m.repositories))
	for _, repo := range m.repositories {
		repos = append(repos, repository.RepoInfo{
			Name:     repo.GetName(),
			URL:      repo.GetURL(),
			Type:     repo.GetType(),
			Priority: repo.GetPriority(),
			LastSync: repo.GetLastSync(),
		})
	}

	return repos
}

func (m *Manager) ListAllManifests() ([]*manifest.Manifest, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var allManifests []*manifest.Manifest
	manifestMap := make(map[string]*manifest.Manifest)

	repos := make([]Repository, 0, len(m.repositories))
	for _, repo := range m.repositories {
		repos = append(repos, repo)
	}
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].GetPriority() > repos[j].GetPriority()
	})

	for _, repo := range repos {
		manifests, err := repo.ListManifests()
		if err != nil {
			continue
		}

		for _, m := range manifests {

			if _, exists := manifestMap[m.Name]; !exists {
				manifestMap[m.Name] = m
				allManifests = append(allManifests, m)
			}
		}
	}

	return allManifests, nil
}
