package repository

// Client define la interfaz para un cliente de repositorio
type Client interface {
	// Sync sincroniza el repositorio con la fuente remota
	Sync() error

	// GetInfo devuelve información básica sobre el repositorio
	GetInfo() RepoInfo
}
