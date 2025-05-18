package unit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DiegoDev2/Fleet/internal/core/manifest"
	"github.com/stretchr/testify/assert"
)

func TestParseManifest(t *testing.T) {

	tempDir := t.TempDir()
	manifestPath := filepath.Join(tempDir, "test.yaml")

	manifestContent := `  
name: test-tool  
version: 1.0.0  
description: "A test tool"  
homepage: "https://example.com"  
license: "MIT"  
categories:  
  - test  
  - example  
platforms:  
  darwin:  
    architecture:  
      amd64:  
        url: "https://example.com/test-1.0.0-darwin-amd64.tar.gz"  
        type: "tar.gz"  
        checksum: "sha256:abcdef1234567890abcdef1234567890"  
        install_steps:  
          - execute: "tar -xzf test-1.0.0-darwin-amd64.tar.gz"  
          - copy: "test"   
            to: "/usr/local/bin/"  
dependencies:  
  - name: "dependency1"  
    version: ">=2.0.0"  
post_install:  
  - "test --version"  
`

	err := os.WriteFile(manifestPath, []byte(manifestContent), 0644)
	assert.NoError(t, err)

	// Probar ParseFile
	m, err := manifest.ParseFile(manifestPath)
	assert.NoError(t, err)
	assert.NotNil(t, m)

	// Verificar que los datos se parsearon correctamente
	assert.Equal(t, "test-tool", m.Name)
	assert.Equal(t, "1.0.0", m.Version)
	assert.Equal(t, "A test tool", m.Description)
	assert.Equal(t, "MIT", m.License)
	assert.Len(t, m.Categories, 2)
	assert.Contains(t, m.Categories, "test")
	assert.Contains(t, m.Categories, "example")

	// Verificar plataformas
	assert.Contains(t, m.Platforms, "darwin")
	darwinPlatform := m.Platforms["darwin"]

	// Verificar arquitecturas
	assert.Contains(t, darwinPlatform.Architecture, "amd64")
	amd64Arch := darwinPlatform.Architecture["amd64"]
	assert.Equal(t, "https://example.com/test-1.0.0-darwin-amd64.tar.gz", amd64Arch.URL)
	assert.Equal(t, "tar.gz", amd64Arch.Type)

	// Verificar pasos de instalación
	assert.Len(t, amd64Arch.InstallSteps, 2)
	assert.Equal(t, "tar -xzf test-1.0.0-darwin-amd64.tar.gz", amd64Arch.InstallSteps[0].Execute)

	// Verificar dependencias
	assert.Len(t, m.Dependencies, 1)
	assert.Equal(t, "dependency1", m.Dependencies[0].Name)
	assert.Equal(t, ">=2.0.0", m.Dependencies[0].Version)

	// Verificar post-install
	assert.Len(t, m.PostInstall, 1)
	assert.Equal(t, "test --version", m.PostInstall[0])
}

func TestParseDirectory(t *testing.T) {
	// Crear un directorio temporal con múltiples manifiestos
	tempDir := t.TempDir()

	// Crear primer manifiesto
	manifest1Path := filepath.Join(tempDir, "tool1.yaml")
	manifest1Content := `  
name: tool1  
version: 1.0.0  
platforms:  
  darwin:  
    architecture:  
      amd64:  
        url: "https://example.com/tool1.tar.gz"  
        type: "tar.gz"  
        checksum: "sha256:abcdef"  
        install_steps:  
          - execute: "echo Installing tool1"  
`
	err := os.WriteFile(manifest1Path, []byte(manifest1Content), 0644)
	assert.NoError(t, err)

	// Crear segundo manifiesto
	manifest2Path := filepath.Join(tempDir, "tool2.yaml")
	manifest2Content := `  
name: tool2  
version: 2.0.0  
platforms:  
  linux:  
    architecture:  
      amd64:  
        url: "https://example.com/tool2.tar.gz"  
        type: "tar.gz"  
        checksum: "sha256:123456"  
        install_steps:  
          - execute: "echo Installing tool2"  
`
	err = os.WriteFile(manifest2Path, []byte(manifest2Content), 0644)
	assert.NoError(t, err)

	// Probar ParseDirectory
	manifests, err := manifest.ParseDirectory(tempDir)
	assert.NoError(t, err)
	assert.Len(t, manifests, 2)

	// Verificar que ambos manifiestos se cargaron correctamente
	assert.Contains(t, manifests, "tool1")
	assert.Contains(t, manifests, "tool2")

	// Verificar contenido de los manifiestos
	assert.Equal(t, "1.0.0", manifests["tool1"].Version)
	assert.Equal(t, "2.0.0", manifests["tool2"].Version)
}
