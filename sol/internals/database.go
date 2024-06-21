package internals

// make databaseinterface
type databaseInterface interface {
	New() error   // create a new database
	Store() error // store the database
}

// make database struct
type Database struct {
	//path to database
	path string
}

// create a new database
func (d *Database) New(path string) error {
	d.path = path
	return nil
}

// store the database
func (d *Database) Store() error {
	return nil
}
