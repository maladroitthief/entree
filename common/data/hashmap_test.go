package data_test

import (
	"math/rand"
	"testing"
)

func BenchmarkHashmapIndexing(b *testing.B) {
	type entity struct {
		id int
	}
	hashmap := map[int]entity{}
	for i := 0; i < ContainerSize; i++ {
		hashmap[i] = entity{id: rand.Int()}
	}

	for n := 0; n < b.N; n++ {
		_ = hashmap[rand.Intn(ContainerSize)]
	}
}

func BenchmarkHashmapInsert(b *testing.B) {
	type entity struct {
		id int
	}
	hashmap := map[int]entity{}

	for n := 0; n < b.N; n++ {
		hashmap[n] = entity{id: rand.Int()}
	}
}

func BenchmarkHashDelete(b *testing.B) {
	type entity struct {
		id int
	}
	hashmap := map[int]entity{}
	indexIds := []int{}
	for i := 0; i < ContainerSize; i++ {
		hashmap[i] = entity{id: rand.Int()}
		indexIds = append(indexIds, i)
	}

	for n := 0; n < b.N; n++ {
		if n < ContainerSize {
			index := rand.Intn(len(indexIds))

			delete(hashmap, indexIds[index])

			indexIds[index] = indexIds[len(indexIds)-1]
			indexIds = indexIds[:len(indexIds)-1]
		}
	}
}
