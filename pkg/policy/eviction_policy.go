package policy

type EvictionPolicy interface {
	Evict() (string, error)
}

func GetEvictionPolicy(policy string) EvictionPolicy {
	mp := make(map[string]EvictionPolicy)
	mp["1"] = NewLRU()
	mp["2"] = NewLFU()
	if val, ok := mp[policy]; ok {
		return val
	}
	return nil
}
