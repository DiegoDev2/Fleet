package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	internalManifest "github.com/DiegoDev2/Fleet/internal/core/manifest"
	internalValidate "github.com/DiegoDev2/Fleet/internal/core/manifest"
	"github.com/DiegoDev2/Fleet/pkg/manifest"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type GitRepository struct {
	name      string
	url       string
	cacheDir  string
	priority  int
	lastSync  time.Time
	branch    string
	manifests map[string]*manifest.Manifest
}

// NewGitRepository crea un nuevo repositorio Git
func NewGitRepository(name, url, cacheDir string, priority int) (*GitRepository, error) {
	return &GitRepository{
		name:      name,
		url:       url,
		cacheDir:  cacheDir,
		priority:  priority,
		branch:    "main", // Rama predeterminada
		manifests: make(map[string]*manifest.Manifest),
	}, nil
}

// GetName devuelve el nombre del repositorio
func (r *GitRepository) GetName() string {
	return r.name
}

// GetURL devuelve la URL del repositorio
func (r *GitRepository) GetURL() string {
	return r.url
}

// GetType devuelve el tipo del repositorio
func (r *GitRepository) GetType() string {
	return "git"
}

// GetPriority devuelve la prioridad del repositorio
func (r *GitRepository) GetPriority() int {
	return r.priority
}

// GetLastSync devuelve la fecha de la última sincronización
func (r *GitRepository) GetLastSync() time.Time {
	return r.lastSync
}

// SetBranch establece la rama a utilizar
func (r *GitRepository) SetBranch(branch string) {
	r.branch = branch
}

// Sync sincroniza el repositorio con la fuente remota
func (r *GitRepository) Sync() error {
	manifestsDir := filepath.Join(r.cacheDir, "manifests")

	// Verificar si el repositorio ya existe localmente
	repo, err := git.PlainOpen(r.cacheDir)
	if err != nil {
		// Si no existe, clonarlo
		repo, err = git.PlainClone(r.cacheDir, false, &git.CloneOptions{
			URL:           r.url,
			ReferenceName: plumbing.NewBranchReferenceName(r.branch),
			SingleBranch:  true,
			Progress:      os.Stdout,
		})
		if err != nil {
			return fmt.Errorf("error al clonar repositorio: %w", err)
		}
	} else {
		// Si existe, actualizarlo
		w, err := repo.Worktree()
		if err != nil {
			return fmt.Errorf("error al obtener worktree: %w", err)
		}

		err = w.Pull(&git.PullOptions{
			RemoteName:    "origin",
			ReferenceName: plumbing.NewBranchReferenceName(r.branch),
			Progress:      os.Stdout,
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			return fmt.Errorf("error al actualizar repositorio: %w", err)
		}
	}

	// Cargar los manifiestos
	r.manifests = make(map[string]*manifest.Manifest)
	manifests, err := internalManifest.ParseDirectory(manifestsDir)
	if err != nil {
		return fmt.Errorf("error al cargar manifiestos: %w", err)
	}

	// Validar los manifiestos
	for name, m := range manifests {
		errors := internalValidate.Validate(m)
		if len(errors) > 0 {
			fmt.Printf("Advertencia: manifiesto inválido %s en repositorio %s\n", name, r.name)
			continue
		}
		r.manifests[name] = m
	}

	r.lastSync = time.Now()
	return nil
}

func (r *GitRepository) GetManifest(toolName string) (*manifest.Manifest, error) {
	if m, exists := r.manifests[toolName]; exists {
		return m, nil
	}
	return nil, fmt.Errorf("manifiesto no encontrado: %s", toolName)
}

func (r *GitRepository) ListManifests() ([]*manifest.Manifest, error) {
	manifests := make([]*manifest.Manifest, 0, len(r.manifests))
	for _, m := range r.manifests {
		manifests = append(manifests, m)
	}
	return manifests, nil
}
