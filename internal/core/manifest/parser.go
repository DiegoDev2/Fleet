package manifest

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/DiegoDev2/Fleet/pkg/manifest"
	"gopkg.in/yaml.v3"
)

func ParseFile(filePath string) (*manifest.Manifest, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al leer archivo de manifiesto: %w", err)
	}

	return Parse(data)
}

func Parse(data []byte) (*manifest.Manifest, error) {
	var manifest manifest.Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("error al analizar manifiesto: %w", err)
	}

	return &manifest, nil
}

func ParseDirectory(dirPath string) (map[string]*manifest.Manifest, error) {
	manifests := make(map[string]*manifest.Manifest)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}

		manifest, err := ParseFile(path)
		if err != nil {
			return fmt.Errorf("error al analizar %s: %w", path, err)
		}

		manifests[manifest.Name] = manifest
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error al recorrer directorio: %w", err)
	}

	return manifests, nil
}
