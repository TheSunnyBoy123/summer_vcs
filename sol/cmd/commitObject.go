package cmd

import (
	"fmt"
	"strings"
)

type Commit struct {
	OID     string
	tree    *Tree
	author  *Author
	message string
}

func NewCommit(tree *Tree, author *Author, message string) *Commit {
	return &Commit{
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

func (c *Commit) SetOID(oid string) {
	content := fmt.Sprintf("%s %d\x00%s", c.Type(), len(c.ToString()), c.ToString())

	c.OID = hashContents(content)
}

func (c *Commit) ToString() string {
	lines := []string{"Tree " + c.tree.GetOID(), "Author " + c.author.ToString(), "Committer " + c.author.ToString(), c.message}
	return strings.Join(lines, "\n")
}
