package manifest

import "gopkg.in/yaml.v3"

// Manifest representa la definición de una herramienta
type Manifest struct {
	Name         string              `yaml:"name" json:"name"`
	Version      string              `yaml:"version" json:"version"`
	Description  string              `yaml:"description" json:"description"`
	Homepage     string              `yaml:"homepage" json:"homepage"`
	License      string              `yaml:"license" json:"license"`
	Categories   []string            `yaml:"categories" json:"categories"`
	Platforms    map[string]Platform `yaml:"platforms" json:"platforms"`
	Dependencies []Dependency        `yaml:"dependencies" json:"dependencies"`
	PostInstall  []string            `yaml:"post_install" json:"post_install"`
}

func (m *Manifest) ToYAML(manifest *Manifest) (any, any) {
	panic("unimplemented")
}

// Platform representa detalles de instalación específicos de la plataforma
type Platform struct {
	Architecture    map[string]Architecture   `yaml:"architecture" json:"architecture"`
	PackageManagers map[string]PackageManager `yaml:"package_managers,omitempty" json:"package_managers,omitempty"`
	Fallback        *Fallback                 `yaml:"fallback,omitempty" json:"fallback,omitempty"`
}

// Architecture representa detalles de instalación específicos de la arquitectura
type Architecture struct {
	URL          string `yaml:"url" json:"url"`
	Type         string `yaml:"type" json:"type"`
	Checksum     string `yaml:"checksum" json:"checksum"`
	InstallSteps []Step `yaml:"install_steps" json:"install_steps"`
}

// Step representa un paso de instalación
type Step struct {
	Mount   string `yaml:"mount,omitempty" json:"mount,omitempty"`
	Copy    string `yaml:"copy,omitempty" json:"copy,omitempty"`
	To      string `yaml:"to,omitempty" json:"to,omitempty"`
	Execute string `yaml:"execute,omitempty" json:"execute,omitempty"`
	Unmount string `yaml:"unmount,omitempty" json:"unmount,omitempty"`
}

// PackageManager representa la instalación específica del gestor de paquetes
type PackageManager struct {
	Package string `yaml:"package" json:"package"`
}

// Fallback representa el método de instalación alternativo
type Fallback struct {
	URL          string               `yaml:"url" json:"url"`
	Type         string               `yaml:"type" json:"type"`
	Checksum     string               `yaml:"checksum" json:"checksum"`
	Dependencies FallbackDependencies `yaml:"dependencies" json:"dependencies"`
	BuildSteps   []string             `yaml:"build_steps" json:"build_steps"`
}

// FallbackDependencies representa dependencias para la instalación alternativa
type FallbackDependencies struct {
	Build   []string `yaml:"build" json:"build"`
	Runtime []string `yaml:"runtime" json:"runtime"`
}

// Dependency representa una dependencia de herramienta
type Dependency struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}

func ToYAML(m *Manifest) ([]byte, error) {
	return yaml.Marshal(m)
}
