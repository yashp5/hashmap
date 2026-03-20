package hashmap

type HashMap interface {
	Put(key string, value any)
	Get(key string) (any, bool)
	Delete(key string) bool
}

func hash(key string) int {
	fnvPrime := uint64(1099511628211)
	hash := uint64(14695981039346656037)
	for i := range len(key) {
		hash ^= uint64(key[i])
		hash *= fnvPrime
	}
	return int(hash & 0x7fffffffffffffff)
}
