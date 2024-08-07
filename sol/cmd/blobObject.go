package cmd

import (
	"fmt"
)

type Blob struct {
	Data string
	OID  string
}

func NewBlob(data string) *Blob {
	content := fmt.Sprintf("blob %d\x00%s", len(data), data)
	oid := hashContents(content)
	return &Blob{Data: data, OID: oid}
}

func (b *Blob) Type() string {
	return "blob"
}

func (b *Blob) ToString() string {
	return b.Data
}

func (b *Blob) GetOID() string {
	return b.OID
}

func (b *Blob) SetOID(oid string) {
	b.OID = oid
}

func (b *Blob) Mode() string {
	return "100644"
}

func (b *Blob) GetName() string {
	return "Some blob"
}
