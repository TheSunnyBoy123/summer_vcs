package cmd

import (
	"fmt"
)

type Blob struct {
	Data string
	OID  string
}

func NewBlob(data string) *Blob {
	return &Blob{Data: data}
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

func (b *Blob) SetOID(oid string) {
	content := fmt.Sprintf("%s %d\x00%s", b.Type(), len(b.ToString()), b.ToString())

	b.OID = hashContents(content)
}
