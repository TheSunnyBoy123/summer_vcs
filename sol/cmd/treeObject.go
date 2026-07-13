package cmd

import (
	"fmt"
	"sort"
)

const (
	ENTRY_FORMAT = "%s %s %s\x00"
)

type Tree struct {
	Entries map[string]SolObject
	OID     string
	name    string
}

func NewTree(name string) *Tree {
	entries := make(map[string]SolObject)
	return &Tree{Entries: entries, OID: "", name: name}
}

func BuildTree(entries []*Entry) *Tree {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	root := NewTree("")
	for _, entry := range entries {
		root.AddEntry(entry.ParentDirectories(), entry)
	}
	return root
}

func (t *Tree) AddEntry(parentDirectories []string, entry *Entry) {
	if len(parentDirectories) == 0 {
		t.Entries[entry.Basename()] = entry
	} else {
		dirName := parentDirectories[0]
		subtree, ok := t.Entries[dirName]
		if !ok {
			subtree = NewTree(dirName)
			t.Entries[dirName] = subtree
		}
		subtree.(*Tree).AddEntry(parentDirectories[1:], entry)
	}
}

func (t *Tree) Type() string {
	return "tree"
}

func (t *Tree) ToString() string {
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

func (t *Tree) GetName() string {
	return t.name
}

func (t *Tree) Mode() string {
	return "40000"
}

func (t *Tree) GetOID() string {
	return t.OID
}

func (t *Tree) SetOID(oid string) {
	content := fmt.Sprintf("tree %d\x00%s", len(t.ToString()), t.ToString())
	t.OID = hashContents(content)
}
