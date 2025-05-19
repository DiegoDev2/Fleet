package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	internalManifest "github.com/DiegoDev2/Fleet/internal/core/manifest"
	"github.com/DiegoDev2/Fleet/pkg/manifest"
)

type Cache struct {
	cacheDir string
	mutex    sync.RWMutex
	metadata CacheMetadata
}

type CacheMetadata struct {
	LastUpdate time.Time                   `json:"last_update"`
	Repos      map[string]RepoMetadata     `json:"repos"`
	Manifests  map[string]ManifestMetadata `json:"manifests"`
}

type RepoMetadata struct {
	LastSync time.Time `json:"last_sync"`
	URL      string    `json:"url"`
	Type     string    `json:"type"`
}

type ManifestMetadata struct {
	RepoName string    `json:"repo_name"`
	Path     string    `json:"path"`
	LastSync time.Time `json:"last_sync"`
	Checksum string    `json:"checksum"`
}

func NewCache(cacheDir string) (*Cache, error) {

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("error creating cache directory: %w", err)
	}

	cache := &Cache{
		cacheDir: cacheDir,
		metadata: CacheMetadata{
			LastUpdate: time.Now(),
			Repos:      make(map[string]RepoMetadata),
			Manifests:  make(map[string]ManifestMetadata),
		},
	}

	metadataPath := filepath.Join(cacheDir, "metadata.json")
	if _, err := os.Stat(metadataPath); err == nil {
		data, err := os.ReadFile(metadataPath)
		if err == nil {
			if err := json.Unmarshal(data, &cache.metadata); err != nil {
				fmt.Printf("Warning: Failed to parse cache metadata: %v\n", err)

			}
		}
	}

	return cache, nil
}

func (c *Cache) SaveMetadata() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.metadata.LastUpdate = time.Now()
	data, err := json.MarshalIndent(c.metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling cache metadata: %w", err)
	}

	metadataPath := filepath.Join(c.cacheDir, "metadata.json")
	if err := os.WriteFile(metadataPath, data, 0644); err != nil {
		return fmt.Errorf("error writing cache metadata: %w", err)
	}

	return nil
}

func (c *Cache) AddRepo(name, url, repoType string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.metadata.Repos[name] = RepoMetadata{
		LastSync: time.Now(),
		URL:      url,
		Type:     repoType,
	}
}

func (c *Cache) RemoveRepo(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.metadata.Repos, name)

	for manifestName, metadata := range c.metadata.Manifests {
		if metadata.RepoName == name {
			delete(c.metadata.Manifests, manifestName)
		}
	}
}

func (c *Cache) GetManifestPath(toolName string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	metadata, exists := c.metadata.Manifests[toolName]
	if !exists {
		return "", false
	}

	return metadata.Path, true
}

func (c *Cache) AddManifest(toolName, repoName, checksum string, manifestData []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	manifestsDir := filepath.Join(c.cacheDir, "manifests")
	if err := os.MkdirAll(manifestsDir, 0755); err != nil {
		return fmt.Errorf("error creating manifests directory: %w", err)
	}

	manifestPath := filepath.Join(manifestsDir, toolName+".yaml")
	if err := os.WriteFile(manifestPath, manifestData, 0644); err != nil {
		return fmt.Errorf("error writing manifest file: %w", err)
	}

	c.metadata.Manifests[toolName] = ManifestMetadata{
		RepoName: repoName,
		Path:     manifestPath,
		LastSync: time.Now(),
		Checksum: checksum,
	}

	return nil
}

func (c *Cache) GetManifest(toolName string) (*manifest.Manifest, error) {
	c.mutex.RLock()
	metadata, exists := c.metadata.Manifests[toolName]
	c.mutex.RUnlock()

	if !exists {
		return nil, fmt.Errorf("manifest not found in cache: %s", toolName)
	}

	// Parse the manifest file
	return internalManifest.ParseFile(metadata.Path)
}

func (c *Cache) ListManifests() (map[string]*manifest.Manifest, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	manifests := make(map[string]*manifest.Manifest)
	for toolName, metadata := range c.metadata.Manifests {
		m, err := internalManifest.ParseFile(metadata.Path)
		if err != nil {
			fmt.Printf("Warning: Failed to parse manifest %s: %v\n", toolName, err)
			continue
		}
		manifests[toolName] = m
	}

	return manifests, nil
}

func (c *Cache) ClearCache() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := os.RemoveAll(c.cacheDir); err != nil {
		return fmt.Errorf("error clearing cache: %w", err)
	}

	if err := os.MkdirAll(c.cacheDir, 0755); err != nil {
		return fmt.Errorf("error recreating cache directory: %w", err)
	}

	// Reset metadata
	c.metadata = CacheMetadata{
		LastUpdate: time.Now(),
		Repos:      make(map[string]RepoMetadata),
		Manifests:  make(map[string]ManifestMetadata),
	}

	return nil
}
