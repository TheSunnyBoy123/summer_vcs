package cmd

import (
	"os"
	"strings"
)

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

func (e *Entry) ParentDirectories() string {
	name := e.Name
	os_seperator := string(os.PathSeparator)
	components := strings.Split(name, os_seperator)
	return strings.Join(components[:len(components)-1], os_seperator)
}

func (e *Entry) GetName() string {
	return e.Name
}

func (e *Entry) GetOID() string {
	return e.OID
}

func (e *Entry) Mode() string {
	// if stat shows file is executable
	if e.Stat.Mode()&0111 != 0 {
		return "100755"
	}
	return "100644"
	// fmt.Println("Stat = ", e.Stat)
	// fmt.Println("Mode = ", e.Stat.Mode()
}
