package cmd

import (
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

func (ws *Workspace) ListFiles() ([]string, error) {
	ignore := map[string]bool{
		".":    true,
		"..":   true,
		".sol": true,
	}

	files, err := ioutil.ReadDir(ws.Pathname)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range files {
		if !ignore[file.Name()] {
			result = append(result, file.Name())
		}
	}

	return result, nil
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
