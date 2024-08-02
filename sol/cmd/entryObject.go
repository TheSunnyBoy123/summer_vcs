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

func (e *Entry) Type() string {
	return "blob"
}

func (e *Entry) ToString() string {
	oid := e.GetOID()
	fileContents := readFile(".sol/objects/" + oid[:2] + "/" + oid[2:])
	split := strings.Split(fileContents, "\x00")
	return split[1]
}

// ParentDirectories returns the parent directories of the entry
func (e *Entry) ParentDirectories() []string {
	separator := os.PathSeparator
	// convert to string and split by separator
	components := strings.Split(e.Name, string(separator))
	// return all but the last component
	// fmt.Println("returning: ", components[:len(components)-1])
	return components[:len(components)-1]
}

func (e *Entry) Basename() string {
	separator := os.PathSeparator
	components := strings.Split(e.Name, string(separator))
	return components[len(components)-1]
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
