package policy

type EvictionPolicy interface {
	Evict() (string, error)
}
