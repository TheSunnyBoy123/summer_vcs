package cmd

import (
	"fmt"
	"path/filepath"
)

type SolObject interface {
	Type() string
	ToString() string
	GetOID() string
	SetOID(string)
}

type Database struct {
	Pathname string
}

func NewDatabase(pathname string) *Database {
	return &Database{
		Pathname: pathname,
	}
}

func (db *Database) Store(object SolObject) error {
	content := fmt.Sprintf("%s %d\x00%s", object.Type(), len(object.ToString()), object.ToString())

	oid := hashContents(content)

	content = compress(content)

	db.writeObject(oid, content)
	fmt.Println("Writing data:", content)
	return nil
}

func (db *Database) writeObject(oid, content string) error {
	createDir(filepath.Join(db.Pathname, oid[:2]))
	objectPath := filepath.Join(db.Pathname, oid[:2], oid[2:])
	writeFile(objectPath, content)
	return nil
}
