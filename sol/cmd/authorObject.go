package cmd

type Author struct {
	name      string
	email     string
	timestamp string
}

func NewAuthor(name, email, timestamp string) *Author {
	return &Author{
		name:      name,
		email:     email,
		timestamp: timestamp,
	}
}

func (a *Author) ToString() string {
	return a.name + " <" + a.email + "> " + a.timestamp
}
