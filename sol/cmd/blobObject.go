package cmd

import (
	"fmt"
)

type Blob struct {
	Data string
	OID  string
}

func NewBlob(data string) *Blob {
	content := fmt.Sprintf("Blob %d\x00%s", len(data), data)
	oid := hashContents(content)
	return &Blob{Data: data, OID: oid}
}

func (b *Blob) Type() string {
	return "Blob"
}

func (b *Blob) ToString() string {
	return b.Data
}

func (b *Blob) GetOID() string {
	return b.OID
}

// func (b *Blob) SetOID(oid string) {

// }
