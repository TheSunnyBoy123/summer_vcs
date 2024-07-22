package cmd

import (
	"fmt"
	"strings"
)

type Tree struct {
	Entries []*Entry
	OID     string
}

func NewTree(entries []*Entry) *Tree {
	return &Tree{Entries: entries}
}

func (t *Tree) Type() string {
	return "Tree"
}

func (t *Tree) ToString() string {
	var sb strings.Builder
	for _, entry := range t.Entries {
		sb.WriteString(fmt.Sprintf("%s %s\n", entry.OID, entry.Name))
	}
	return sb.String()
}

func (t *Tree) GetOID() string {
	return t.OID
}

func (t *Tree) SetOID(oid string) {
	t.OID = oid
}
