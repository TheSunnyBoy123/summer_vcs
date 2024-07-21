package cmd

import (
	"bytes"
	"encoding/hex"
	"sort"
)

type Tree struct {
	OID     string
	Entries []*Entry
}

func NewTree(entries []*Entry) *Tree {
	return &Tree{
		Entries: entries,
	}
}

func (t *Tree) Type() string {
	return "tree"
}

func (t *Tree) ToString() string {
	const mode = "100644"
	var buffer bytes.Buffer

	// Sort entries by name
	sort.Slice(t.Entries, func(i, j int) bool {
		return t.Entries[i].Name < t.Entries[j].Name
	})

	// Serialize entries
	for _, entry := range t.Entries {
		oidBytes, err := hex.DecodeString(entry.OID)
		if err != nil {
			continue // Handle error appropriately in real code
		}
		if len(oidBytes) != 20 {
			continue // Ensure OID is 20 bytes long
		}
		buffer.WriteString(mode + " " + entry.Name + "\x00")
		buffer.Write(oidBytes)
	}

	return buffer.String()
}
