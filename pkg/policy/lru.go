package policy

type LRU struct {
}

func (l *LRU) Evict() (string, error) {
	return "", nil
}

func NewLRU() *LRU {
	return &LRU{}
}
