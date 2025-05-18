package repository

import (
	"time"

	"github.com/DiegoDev2/Fleet/pkg/manifest"
)

// Repository define la interfaz que deben implementar todos los repositorios
type Repository interface {
	// GetName devuelve el nombre del repositorio
	GetName() string

	// GetURL devuelve la URL del repositorio
	GetURL() string

	// GetType devuelve el tipo del repositorio
	GetType() string

	// GetPriority devuelve la prioridad del repositorio
	GetPriority() int

	// GetLastSync devuelve la fecha de la última sincronización
	GetLastSync() time.Time

	// Sync sincroniza el repositorio con la fuente remota
	Sync() error

	// GetManifest obtiene un manifiesto específico del repositorio
	GetManifest(toolName string) (*manifest.Manifest, error)

	// ListManifests devuelve todos los manifiestos disponibles en el repositorio
	ListManifests() ([]*manifest.Manifest, error)
}
