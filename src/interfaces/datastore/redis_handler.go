package datastore

type RedisHandler interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Close() error
	Delete(key string) error
}
