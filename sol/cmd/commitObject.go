package cmd

import (
	"fmt"
	"strings"
)

type Commit struct {
	parent  string
	OID     string
	tree    *Tree
	author  *Author
	message string
}

func NewCommit(parent string, tree *Tree, author *Author, message string) *Commit {
	return &Commit{
		parent:  parent,
		tree:    tree,
		author:  author,
		message: message,
	}
}

func (c *Commit) Type() string {
	return "Commit"
}

func (c *Commit) GetOID() string {
	return c.OID
}

func (c *Commit) GetParent() string {
	return c.parent
}

func (c *Commit) SetOID(oid string) {
	content := fmt.Sprintf("%s %d\x00%s", c.Type(), len(c.ToString()), c.ToString())

	c.OID = hashContents(content)
}

func (c *Commit) ToString() string {
	lines := []string{}

	lines = append(lines, fmt.Sprintf("tree %s", c.tree.GetOID()))
	if c.parent != "" {
		lines = append(lines, fmt.Sprintf("parent %s", c.parent))
	}
	lines = append(lines, fmt.Sprintf("author %s", c.author.ToString()))
	lines = append(lines, fmt.Sprintf("committer %s", c.author.ToString()))
	lines = append(lines, "")
	lines = append(lines, c.message)

	return strings.Join(lines, "\n")
}
