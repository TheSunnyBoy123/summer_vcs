package cmd

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const refPrefix = "ref: "

type Refs struct {
	pathname string
}

func NewRefs(pathname string) *Refs {
	return &Refs{pathname: pathname}
}

func (r *Refs) headPath() string {
	return filepath.Join(r.pathname, "HEAD")
}

func (r *Refs) headsPath() string {
	return filepath.Join(r.pathname, "refs", "heads")
}

func (r *Refs) branchPath(name string) string {
	return filepath.Join(r.headsPath(), name)
}

func (r *Refs) ReadHead() string {
	if !fileExists(r.headPath()) {
		return ""
	}
	contents := strings.TrimSpace(readFile(r.headPath()))
	if strings.HasPrefix(contents, refPrefix) {
		ref := strings.TrimPrefix(contents, refPrefix)
		refPath := filepath.Join(r.pathname, ref)
		if fileExists(refPath) {
			return strings.TrimSpace(readFile(refPath))
		}
		return ""
	}
	return contents
}

func (r *Refs) UpdateHead(oid string) {
	writeFile(r.headPath(), oid+"\n")
}

func (r *Refs) SetHeadToBranch(name string) {
	writeFile(r.headPath(), refPrefix+"refs/heads/"+name+"\n")
}

func (r *Refs) CurrentBranch() string {
	if !fileExists(r.headPath()) {
		return ""
	}
	contents := strings.TrimSpace(readFile(r.headPath()))
	if strings.HasPrefix(contents, refPrefix) {
		ref := strings.TrimPrefix(contents, refPrefix)
		parts := strings.Split(ref, "/")
		return parts[len(parts)-1]
	}
	return ""
}

func (r *Refs) ReadBranch(name string) string {
	path := r.branchPath(name)
	if !fileExists(path) {
		return ""
	}
	return strings.TrimSpace(readFile(path))
}

func (r *Refs) UpdateBranch(name, oid string) {
	if !dirExists(r.headsPath()) {
		createDir(r.headsPath())
	}
	writeFile(r.branchPath(name), oid+"\n")
}

func (r *Refs) DeleteBranch(name string) {
	path := r.branchPath(name)
	if fileExists(path) {
		deleteFile(path)
	}
}

func (r *Refs) ListBranches() []string {
	if !dirExists(r.headsPath()) {
		return nil
	}
	entries, _ := ioutil.ReadDir(r.headsPath())
	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return names
}

func (r *Refs) tagsPath() string {
	return filepath.Join(r.pathname, "refs", "tags")
}

func (r *Refs) tagPath(name string) string {
	return filepath.Join(r.tagsPath(), name)
}

func (r *Refs) ReadTag(name string) string {
	path := r.tagPath(name)
	if !fileExists(path) {
		return ""
	}
	return strings.TrimSpace(readFile(path))
}

func (r *Refs) UpdateTag(name, oid string) {
	if !dirExists(r.tagsPath()) {
		createDir(r.tagsPath())
	}
	writeFile(r.tagPath(name), oid+"\n")
}

func (r *Refs) DeleteTag(name string) {
	path := r.tagPath(name)
	if fileExists(path) {
		deleteFile(path)
	}
}

func (r *Refs) ListTags() []string {
	if !dirExists(r.tagsPath()) {
		return nil
	}
	entries, _ := ioutil.ReadDir(r.tagsPath())
	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return names
}
