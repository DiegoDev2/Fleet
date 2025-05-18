package manifest

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

type Platform struct {
	Architecture    map[string]Architecture   `yaml:"architecture" json:"architecture"`
	PackageManagers map[string]PackageManager `yaml:"package_managers,omitempty" json:"package_managers,omitempty"`
	Fallback        *Fallback                 `yaml:"fallback,omitempty" json:"fallback,omitempty"`
}

type Architecture struct {
	URL          string `yaml:"url" json:"url"`
	Type         string `yaml:"type" json:"type"`
	Checksum     string `yaml:"checksum" json:"checksum"`
	InstallSteps []Step `yaml:"install_steps" json:"install_steps"`
}

type Step struct {
	Mount   string `yaml:"mount,omitempty" json:"mount,omitempty"`
	Copy    string `yaml:"copy,omitempty" json:"copy,omitempty"`
	To      string `yaml:"to,omitempty" json:"to,omitempty"`
	Execute string `yaml:"execute,omitempty" json:"execute,omitempty"`
	Unmount string `yaml:"unmount,omitempty" json:"unmount,omitempty"`
}

type PackageManager struct {
	Package string `yaml:"package" json:"package"`
}

type Fallback struct {
	URL          string               `yaml:"url" json:"url"`
	Type         string               `yaml:"type" json:"type"`
	Checksum     string               `yaml:"checksum" json:"checksum"`
	Dependencies FallbackDependencies `yaml:"dependencies" json:"dependencies"`
	BuildSteps   []string             `yaml:"build_steps" json:"build_steps"`
}

type FallbackDependencies struct {
	Build   []string `yaml:"build" json:"build"`
	Runtime []string `yaml:"runtime" json:"runtime"`
}

type Dependency struct {
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}
