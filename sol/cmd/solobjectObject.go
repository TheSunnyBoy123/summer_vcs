// go:build exclude

package cmd

type SolObject interface {
	Type() string
	ToString() string
	GetOID() string
	Mode() string
	GetName() string
}
