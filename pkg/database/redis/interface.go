package redis

type RedisEngine interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}
