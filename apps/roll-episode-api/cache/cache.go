package cache

type Cache interface {
	Set(key string, value []byte)
	Get(key string) ([]byte, bool)
}

type Map struct {
	cache map[string][]byte
}

func (m *Map) Set(key string, value []byte) {
	if m.cache == nil {
		m.cache = make(map[string][]byte)
	}

	m.cache[key] = value
}

func (m *Map) Get(key string) ([]byte, bool) {
	if m.cache == nil {
		return []byte{}, false
	}

	b, c := m.cache[key]
	return b, c
}

