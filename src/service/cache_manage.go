package service

type cacheManager interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	GetString(key string) string
}

type simpleMapCache struct {
	mapObj map[string]interface{}
}

func (m *simpleMapCache) Set(key string, value interface{}) {
	m.mapObj[key] = value
}

func (m *simpleMapCache) Get(key string) interface{} {
	return m.mapObj[key]
}

func (m *simpleMapCache) GetString(key string) string {
	return m.mapObj[key].(string)
}

func newSimpleMapCache() cacheManager {
	return &simpleMapCache{
		mapObj: make(map[string]interface{}),
	}
}

var Cache cacheManager = newSimpleMapCache()

func init() {

	//println("cache_manage.go init ...")
}
