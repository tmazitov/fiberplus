package services

type Database struct{}

func (d *Database) Test() string {
	return "Hello world!"
}

type Services struct {
	Database Database
}
