package unit

import (
	"testing"

	"github.com/DiegoDev2/Fleet/internal/core/repository"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryManager(t *testing.T) {

	tempDir := t.TempDir()

	manager, err := repository.NewManager(tempDir)
	assert.NoError(t, err)
	assert.NotNil(t, manager)

	err = manager.AddRepository("test-repo", "https://example.com/repo", "http", 100)
	assert.NoError(t, err)

	repos := manager.ListRepositories()
	assert.Len(t, repos, 1)
	assert.Equal(t, "test-repo", repos[0].Name)
	assert.Equal(t, "https://example.com/repo", repos[0].URL)
	assert.Equal(t, "http", repos[0].Type)
	assert.Equal(t, 100, repos[0].Priority)

	err = manager.RemoveRepository("test-repo")
	assert.NoError(t, err)

	repos = manager.ListRepositories()
	assert.Empty(t, repos)
}
