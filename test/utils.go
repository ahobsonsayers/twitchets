package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// ProjectDirectory returns the directory of the project being worked on,
// by walking the tree upwards until it finds a go.mod file.
func ProjectDirectory(t *testing.T) string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		require.NoError(t, err, "failed to get path of current working directory")
	}

	directory := workingDirectory
	for directory != "/" {
		_, err := os.Stat(filepath.Join(directory, "go.mod"))
		if err == nil {
			break
		}
		directory = filepath.Dir(directory)
	}
	require.NotEqual(t, "failed find project directory", directory, "/")

	return directory
}

// ProjectDirectoryJoin returns a path joined to the project directory.
// See `ProjectDirectory` for more information
func ProjectDirectoryJoin(t *testing.T, pathElems ...string) string {
	pathElems = append(
		[]string{ProjectDirectory(t)},
		pathElems...,
	)
	return filepath.Join(pathElems...)
}
