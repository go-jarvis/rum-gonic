package client

type Meta map[string][]string

func (meta Meta) Get(key string) []string {
	return meta[key]
}

func (meta Meta) Set(key string, value ...string) {
	meta[key] = value
}

func (meta Meta) Add(key string, value ...string) {
	val, ok := meta[key]
	if ok {
		val = append(val, value...)
	}
	meta.Set(key, val...)
}

func (meta Meta) Del(key string) {
	delete(meta, key)
}
