package storage

type Device struct {
	id int
	name string
}

type storage interface {
	OpenStorage() (interface {}, error)
	Init()
	Close()
	GetDevices() ([] Device, error)
}

