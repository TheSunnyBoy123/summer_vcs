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
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	listEntries := ""

	for _, entry := range entries {
		if entry == nil {
			fmt.Println("Entry is nil")
			continue
		} else {
			fmt.Println("Entry is not nil", entry)
		}
		fmt.Println("Calling Mode() on entry:", entry)
		mode := entry.Mode()
		fmt.Println("Mode for entry:", mode)
		thisEntry := fmt.Sprintf(ENTRY_FORMAT, mode, entry.GetOID(), entry.GetName())
		fmt.Println("OID for " + entry.Name + " = " + entry.GetOID())
		listEntries += thisEntry
	}

	contents := fmt.Sprintf("tree %d\x00%s", len(listEntries), listEntries)
	oid := hashContents(contents)

	obj := &Tree{Entries: entries, OID: oid}
	// obj.SetOID("")
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
		thisEntry := fmt.Sprintf(ENTRY_FORMAT, entry.Mode(), entry.GetOID(), entry.GetName())
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
