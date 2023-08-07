package data_test

import (
	"math/rand"
	"testing"
)

func BenchmarkSliceIndexing(b *testing.B) {
	type entity struct {
		id int
	}

	slice := []entity{}
	for i := 0; i < ContainerSize; i++ {
		slice = append(slice, entity{id: rand.Int()})
	}

	for n := 0; n < b.N; n++ {
		_ = slice[rand.Intn(ContainerSize)]
	}
}

func BenchmarkSliceInsert(b *testing.B) {
	type entity struct {
		id int
	}
	slice := []entity{}

	for n := 0; n < b.N; n++ {
		slice = append(slice, entity{id: rand.Int()})
	}
}

func BenchmarkSliceDelete(b *testing.B) {
	type entity struct {
		id int
	}

	slice := []entity{}
	indexIds := []int{}
	for i := 0; i < ContainerSize; i++ {
		slice = append(slice, entity{id: rand.Int()})
		indexIds = append(indexIds, i)
	}

	for n := 0; n < b.N; n++ {
		if n < ContainerSize {
			index := rand.Intn(len(indexIds))

			slice = append(slice[:index], slice[index+1:]...)

			indexIds[index] = indexIds[len(indexIds)-1]
			indexIds = indexIds[:len(indexIds)-1]
		}
	}
}
