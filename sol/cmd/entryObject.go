package cmd

import "os"

type Entry struct {
	Name string
	OID  string
	Stat os.FileInfo
}

func NewEntry(name, oid string, stat os.FileInfo) *Entry {
	return &Entry{
		Name: name,
		OID:  oid,
		Stat: stat,
	}
}

func (e *Entry) Mode() string {
	REGULAR_MODE := "100644"
	EXECUTABLE_MODE := "100755"
	if e.Stat.Mode()&0111 != 0 {
		return EXECUTABLE_MODE
	}
	return REGULAR_MODE
}
