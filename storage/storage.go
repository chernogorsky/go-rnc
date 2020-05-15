package storage

type storage interface {
	OpenStorage() (interface {}, error)
	Init()
	Close()
}

type devices struct {
	id int
	name string
}