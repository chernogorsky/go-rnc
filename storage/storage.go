package storage

type Device struct {
	Id int
	Name string
}

type storage interface {
	OpenStorage() (interface {}, error)
	Init()
	Close()
	GetDevices() ([] Device, error)
}

