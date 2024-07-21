package cmd

import (
	"crypto/sha1"
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
	hash := sha1.New()
	hash.Write([]byte(content))
	object.OID = fmt.Sprintf("%x", hash.Sum(nil))

	if err := db.writeObject(object.OID, content); err != nil {
		return err
	}

	return nil
}

func (db *Database) writeObject(oid, content string) error {
	objectPath := filepath.Join(db.Pathname, oid[:2], oid[2:])
	createDir(filepath.Join(db.Pathname, oid[:2]))

	compressedContent := compress(content)

	writeFile(objectPath, compressedContent)
	return nil
}
