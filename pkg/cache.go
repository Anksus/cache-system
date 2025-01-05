package pkg

import "anksus/cache-system/pkg/policy"

type Cache struct {
	Store       []*KeyValue
	Limit       int
	Policy      policy.EvictionPolicy
	CurrentSize int
}

func NewCache(limit int, currentSize int, policyStr string) *Cache {
	return &Cache{
		Store:       make([]*KeyValue, 0),
		Limit:       limit,
		CurrentSize: currentSize,
		Policy:      policy.GetEvictionPolicy(policyStr),
	}
}
