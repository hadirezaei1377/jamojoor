package bloom

import (
	"hash/fnv"
)

type BloomFilter struct {
	bitset []bool
	size   int
	hashes int
}

func NewBloomFilter(size, hashes int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		size:   size,
		hashes: hashes,
	}
}

func (bf *BloomFilter) hash(item string, seed int) int {
	h := fnv.New64a()
	h.Write([]byte(item))
	h.Write([]byte{byte(seed)})
	return int(h.Sum64() % uint64(bf.size))
}

func (bf *BloomFilter) Add(item string) {
	for i := 0; i < bf.hashes; i++ {
		index := bf.hash(item, i)
		bf.bitset[index] = true
	}
}

func (bf *BloomFilter) Check(item string) bool {
	for i := 0; i < bf.hashes; i++ {
		index := bf.hash(item, i)
		if !bf.bitset[index] {
			return false
		}
	}
	return true
}
