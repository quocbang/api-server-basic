package database

type DataManager interface {
	Close() error
	Services
}
