package cmd

type Blob struct {
	Data string
	OID  string
}

func NewBlob(data string) *Blob {
	return &Blob{Data: data}
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
