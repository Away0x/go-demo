package cache

/**
通过 sync.RWMutex 封装缓存算法, 使缓存支持并发读写
*/

type Getter interface {
	Get(key string) interface{}
}

type GetFunc func(key string) interface{}

func (f GetFunc) Get(key string) interface{} {
	return f(key)
}

type Cache struct {
	mainCache *safeCache
	getter    Getter
}

func NewCache(getter Getter, cache ICache) *Cache {
	return &Cache{
		mainCache: newSafeCache(cache),
		getter:    getter,
	}
}

func (t *Cache) Get(key string) interface{} {
	val := t.mainCache.get(key)
	if val != nil {
		return val
	}

	if t.getter != nil {
		val = t.getter.Get(key)
		if val == nil {
			return nil
		}
		t.mainCache.set(key, val)
		return val
	}

	return nil
}

func (t *Cache) Set(key string, val interface{}) {
	if val == nil {
		return
	}
	t.mainCache.set(key, val)
}

func (t *Cache) Stat() *Stat {
	return t.mainCache.stat()
}
