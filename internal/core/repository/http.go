package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	internalValidate "github.com/DiegoDev2/Fleet/internal/core/manifest"
	"github.com/DiegoDev2/Fleet/pkg/manifest"
	"gopkg.in/yaml.v3"
)

// HTTPRepository implementa la interfaz Repository para repositorios HTTP
type HTTPRepository struct {
	name      string
	url       string
	cacheDir  string
	priority  int
	lastSync  time.Time
	manifests map[string]*manifest.Manifest
}

// NewHTTPRepository crea un nuevo repositorio HTTP
func NewHTTPRepository(name, url, cacheDir string, priority int) (*HTTPRepository, error) {
	return &HTTPRepository{
		name:      name,
		url:       url,
		cacheDir:  cacheDir,
		priority:  priority,
		manifests: make(map[string]*manifest.Manifest),
	}, nil
}

// GetName devuelve el nombre del repositorio
func (r *HTTPRepository) GetName() string {
	return r.name
}

// GetURL devuelve la URL del repositorio
func (r *HTTPRepository) GetURL() string {
	return r.url
}

// GetType devuelve el tipo del repositorio
func (r *HTTPRepository) GetType() string {
	return "http"
}

// GetPriority devuelve la prioridad del repositorio
func (r *HTTPRepository) GetPriority() int {
	return r.priority
}

// GetLastSync devuelve la fecha de la última sincronización
func (r *HTTPRepository) GetLastSync() time.Time {
	return r.lastSync
}

// Sync sincroniza el repositorio con la fuente remota
func (r *HTTPRepository) Sync() error {
	// Crear directorio de caché si no existe
	manifestsDir := filepath.Join(r.cacheDir, "manifests")
	if err := os.MkdirAll(manifestsDir, 0755); err != nil {
		return fmt.Errorf("error al crear directorio de manifiestos: %w", err)
	}

	// Obtener el índice de manifiestos
	indexURL := fmt.Sprintf("%s/index.json", r.url)
	resp, err := http.Get(indexURL)
	if err != nil {
		return fmt.Errorf("error al obtener índice de manifiestos: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error al obtener índice de manifiestos: código de estado %d", resp.StatusCode)
	}

	var index struct {
		Manifests []string `json:"manifests"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		return fmt.Errorf("error al decodificar índice de manifiestos: %w", err)
	}

	// Descargar cada manifiesto
	r.manifests = make(map[string]*manifest.Manifest)
	for _, manifestName := range index.Manifests {
		manifestURL := fmt.Sprintf("%s/manifests/%s.yaml", r.url, manifestName)
		manifestResp, err := http.Get(manifestURL)
		if err != nil {
			fmt.Printf("Error al descargar manifiesto %s: %v\n", manifestName, err)
			continue
		}

		if manifestResp.StatusCode != http.StatusOK {
			fmt.Printf("Error al descargar manifiesto %s: código de estado %d\n", manifestName, manifestResp.StatusCode)
			manifestResp.Body.Close()
			continue
		}

		// Leer el contenido del manifiesto
		manifestData, err := ioutil.ReadAll(manifestResp.Body)
		manifestResp.Body.Close()
		if err != nil {
			fmt.Printf("Error al leer manifiesto %s: %v\n", manifestName, err)
			continue
		}

		// Guardar el manifiesto en caché
		manifestPath := filepath.Join(manifestsDir, manifestName+".yaml")
		if err := ioutil.WriteFile(manifestPath, manifestData, 0644); err != nil {
			fmt.Printf("Error al guardar manifiesto %s en caché: %v\n", manifestName, err)
			continue
		}

		// Parsear el manifiesto
		var m manifest.Manifest
		if err := yaml.Unmarshal(manifestData, &m); err != nil {
			fmt.Printf("Error al parsear manifiesto %s: %v\n", manifestName, err)
			continue
		}

		// Validar el manifiesto
		errors := internalValidate.Validate(&m)
		if len(errors) > 0 {
			fmt.Printf("Advertencia: manifiesto inválido %s en repositorio %s\n", manifestName, r.name)
			continue
		}

		r.manifests[m.Name] = &m
	}

	r.lastSync = time.Now()
	return nil
}

// GetManifest obtiene un manifiesto específico del repositorio
func (r *HTTPRepository) GetManifest(toolName string) (*manifest.Manifest, error) {
	if m, exists := r.manifests[toolName]; exists {
		return m, nil
	}
	return nil, fmt.Errorf("manifiesto no encontrado: %s", toolName)
}

// ListManifests devuelve todos los manifiestos disponibles en el repositorio
func (r *HTTPRepository) ListManifests() ([]*manifest.Manifest, error) {
	manifests := make([]*manifest.Manifest, 0, len(r.manifests))
	for _, m := range r.manifests {
		manifests = append(manifests, m)
	}
	return manifests, nil
}
