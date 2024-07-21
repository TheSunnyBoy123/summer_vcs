package cmd

type Entry struct {
	Name string
	OID  string
}

func NewEntry(name, oid string) *Entry {
	return &Entry{
		Name: name,
		OID:  oid,
	}
}
