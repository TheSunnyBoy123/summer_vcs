package cmd

import (
	"fmt"
	"path/filepath"
)

type Refs struct {
	pathname string
}

func NewRefs(pathname string) *Refs {
	return &Refs{
		pathname: pathname,
	}
}

func (r *Refs) Type() string {
	return "Refs"
}

func (r *Refs) UpdateHead(oid string) {
	if !fileExists(r.headPath()) {
		createFiles([]string{filepath.Join(r.headPath(), "HEAD")})
	}
	writeFile(r.headPath(), oid)

}

func (r *Refs) headPath() string {
	return filepath.Join(r.pathname, "HEAD")
}

func (r *Refs) ReadHead() string {
	fmt.Println("Reading HEAD: ", r.headPath())
	if fileExists(r.headPath()) {
		return readFile(r.headPath())
	}
	return ""
}
