package bloom

import (
	"math"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filterSize uint
	hashCount  uint
	filter     []bool
}

func (bf *BloomFilter) Add(key string) {
	data := []byte(key)
	for i := uint(0); i < bf.hashCount; i += 1 {
		h, _ := murmur3.Sum128WithSeed(data, uint32(i))
		index := h % uint64(bf.filterSize)
		bf.filter[index] = true
	}
}

func (bf *BloomFilter) Contains(key string) bool {
	data := []byte(key)
	for i := uint(0); i < bf.hashCount; i += 1 {
		h, _ := murmur3.Sum128WithSeed(data, uint32(i))
		index := h % uint64(bf.filterSize)
		if !bf.filter[index] {
			return false
		}
	}
	return true
}

func calcFilterCapacity(n uint, e float64) uint {
	tmp := -((float64(n) * (math.Log(e))) / float64(math.Pow(math.Log(2), 2)))
	return uint(math.Ceil(tmp))
}

func calcHashCount(e float64) uint {
	return uint(math.Ceil(-1 * math.Log2(e)))
}

func Init(elementCount uint, tolerance float64) *BloomFilter {
	filterSize := calcFilterCapacity(elementCount, tolerance)
	hashCount := calcHashCount(tolerance)
	return &BloomFilter{
		filterSize: filterSize,
		hashCount:  hashCount,
		filter:     make([]bool, filterSize),
	}
}
