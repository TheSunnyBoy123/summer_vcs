package cmd

type SolObject interface {
	Type() string
	ToString() string
	GetOID() string
	SetOID(string)
	Mode() string
	GetName() string
}
