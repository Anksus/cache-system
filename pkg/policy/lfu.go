package policy

type LFU struct {
}

func (l *LFU) Evict() (string, error) {
	return "", nil
}
