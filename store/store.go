package store

type Store interface {
	Get(key string) error
}
