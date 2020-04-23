package cache

type CacheIFace interface {
	Get(string) (string, error)
	Set(string, string) error
}
