package manifest

import (
	"github.com/DiegoDev2/Fleet/pkg/manifest"
)

type ManifestSource struct {
	Type     string
	Location string
	RepoName string
}

type ManifestWithMetadata struct {
	Manifest *manifest.Manifest
	Source   ManifestSource
	Cached   bool
	Modified string
}

type ValidationContext struct {
	StrictMode bool
	Platform   string
	Arch       string
}

type ManifestIndex struct {
	ByName     map[string]*manifest.Manifest
	ByCategory map[string][]*manifest.Manifest
}

func NewManifestIndex() *ManifestIndex {
	return &ManifestIndex{
		ByName:     make(map[string]*manifest.Manifest),
		ByCategory: make(map[string][]*manifest.Manifest),
	}
}

func (idx *ManifestIndex) AddToIndex(m *manifest.Manifest) {
	idx.ByName[m.Name] = m

	for _, category := range m.Categories {
		idx.ByCategory[category] = append(idx.ByCategory[category], m)
	}
}
