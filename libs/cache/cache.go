package cache

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string, output interface{}) error
	Update(key string, value interface{}) error
	Delete(key string)
}
