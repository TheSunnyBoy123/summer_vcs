package cmd

type Blob struct {
	OID  string
	Data string
}

func NewBlob(data string) *Blob {
	return &Blob{
		Data: data,
	}
}

func (b *Blob) Type() string {
	return "blob"
}

func (b *Blob) ToString() string {
	return b.Data
}
