package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Workspace struct {
	Pathname string
}

func NewWorkspace(pathname string) *Workspace {
	return &Workspace{
		Pathname: pathname,
	}
}

func (ws *Workspace) getRelativePath(path string) (string, error) {
	relPath, err := filepath.Rel(ws.Pathname, path)
	if err != nil {
		return "", err
	}
	return relPath, nil
}

func (ws *Workspace) ListFiles(dir string) ([]string, error) {
	ignore := map[string]bool{
		".":    true,
		"..":   true,
		".sol": true,
	}

	if dir == "" {
		dir = ws.Pathname
	}

	var listFilesRecursive func(string) ([]string, error)
	listFilesRecursive = func(path string) ([]string, error) {
		fmt.Println("Called for path: " + path)
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, err
		}

		var paths []string
		for _, file := range files {
			if ignore[file.Name()] {
				continue
			}

			fullPath := filepath.Join(path, file.Name())
			if file.IsDir() {
				subPaths, err := listFilesRecursive(fullPath)
				if err != nil {
					return nil, err
				}
				paths = append(paths, subPaths...)
			} else {
				relPath, err := ws.getRelativePath(fullPath)
				if err != nil {
					return nil, err
				}
				paths = append(paths, relPath)
			}
		}

		return paths, nil
	}

	return listFilesRecursive(dir)
}

// return stat of file executable or not
func (ws *Workspace) StatFile(path string) os.FileInfo {
	fullPath := filepath.Join(ws.Pathname, path)
	info, _ := os.Stat(fullPath)
	return info
}

func (ws *Workspace) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
