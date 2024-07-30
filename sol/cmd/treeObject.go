// go:build excludex

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
	Entries map[string]*Entry
	OID     string
}

func NewTree() *Tree {

	entries := make(map[string]*Entry)
	return &Tree{Entries: entries, OID: ""}

	// sort.Slice(entries, func(i, j int) bool {
	// 	return entries[i].Name < entries[j].Name
	// })

	// listEntries := ""

	// for _, entry := range entries {
	// 	if entry == nil {
	// 		// fmt.Println("Entry is nil")
	// 		continue
	// 	} else {
	// 		// fmt.Println("Entry is not nil", entry)
	// 	}
	// 	// fmt.Println("Calling Mode() on entry:", entry)
	// 	mode := entry.Mode()
	// 	// fmt.Println("Mode for entry:", mode)
	// 	thisEntry := fmt.Sprintf(ENTRY_FORMAT, mode, entry.GetOID(), entry.GetName())
	// 	// fmt.Println("OID for " + entry.Name + " = " + entry.GetOID())
	// 	listEntries += thisEntry
	// }

	// contents := fmt.Sprintf("tree %d\x00%s", len(listEntries), listEntries)
	// oid := hashContents(contents)

	// obj := &Tree{Entries: entries, OID: oid}
	// // obj.SetOID("")
	// return obj
}

// class method Build to create a tree object which iterates over each entry and adds to the tree
func (t *Tree) Build(entries []*Entry) *Tree {
	// sort entries by name
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	root := NewTree()
	for _, entry := range entries {
		fmt.Println("root.AddEntry(", entry.ParentDirectories(), entry, ")")
		root.AddEntry(entry.ParentDirectories(), entry)
	}
	return root
}

func (t *Tree) AddEntry(parentDirectories []string, entry *Entry) {
	if len(parentDirectories) == 0 {
		t.Entries[entry.Basename()] = entry
	}

}

func (t *Tree) Type() string {
	return "tree"
}

func (t *Tree) ToString() string {
	// sort entries by name
	keys := make([]string, 0, len(t.Entries))
	for k := range t.Entries {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	listEntries := ""
	for _, k := range keys {
		entry := t.Entries[k]
		thisEntry := fmt.Sprintf(ENTRY_FORMAT, entry.Mode(), entry.GetOID(), entry.GetName())
		listEntries += thisEntry
	}

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
