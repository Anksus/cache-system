package pkg

type Cache struct {
	Store []*KeyValue
	Limit int64
}

func NewCache(limit int64) *Cache {
	return &Cache{
		Store: make([]*KeyValue, 0),
		Limit: limit,
	}
}
