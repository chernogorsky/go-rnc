package storage

type storage interface {
	Open() (interface {}, error)
	Init()
	Close()
}
