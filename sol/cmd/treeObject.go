package cmd

import (
	"fmt"
	"sort"
)

const (
	MODE         = "100644"       // Example mode, adjust as needed
	ENTRY_FORMAT = "%s %s %s\x00" // Format for packing entries
)

type Tree struct {
	Entries []*Entry
	OID     string
}

func NewTree(entries []*Entry) *Tree {
	obj := &Tree{Entries: entries, OID: ""}
	obj.SetOID("")
	return obj
}

func (t *Tree) Type() string {
	return "Tree"
}

func (t *Tree) ToString() string {
	// Sort entries by name
	sort.Slice(t.Entries, func(i, j int) bool {
		return t.Entries[i].Name < t.Entries[j].Name
	})

	listEntries := ""
	for _, entry := range t.Entries {
		thisEntry := fmt.Sprintf(ENTRY_FORMAT, entry.Mode(), entry.GetOID(), entry.Name)
		fmt.Println("OID for " + entry.Name + " = " + entry.GetOID())
		listEntries += thisEntry
	}

	// fmt.Println("List entries:", listEntries)
	return listEntries
}

func (t *Tree) GetOID() string {
	return t.OID
}

func (t *Tree) SetOID(oid string) {
	content := fmt.Sprintf("Tree %d\x00%s", len(t.ToString()), t.ToString())
	t.OID = hashContents(content)
}

// add a method to return "tree" when called
