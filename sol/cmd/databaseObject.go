package cmd

import (
	"fmt"
	"path/filepath"
)

// type SolObject interface {
// 	Type() string
// 	ToString() string
// 	GetOID() string
// 	SetOID(string)
// }

type Database struct {
	Pathname string
}

func NewDatabase(pathname string) *Database {
	return &Database{
		Pathname: pathname,
	}
}

func (db *Database) Store(object SolObject) (string, error) {
	content := fmt.Sprintf("%s %d\x00%s", object.Type(), len(object.ToString()), object.ToString())

	oid := hashContents(content)

	compressed := compress(content)

	db.writeObject(oid, compressed)

	object.SetOID(oid)
	return oid, nil
}

func (db *Database) writeObject(oid, content string) error {
	objectPath := filepath.Join(db.Pathname, oid[:2], oid[2:])
	if fileExists(objectPath) {
		fmt.Println("Object already exists")
		return nil
	}
	if !dirExists(filepath.Join(db.Pathname, oid[:2])) {
		createDir(filepath.Join(db.Pathname, oid[:2]))
	}
	writeFile(objectPath, content)
	return nil
}
