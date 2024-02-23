package belajargoroutines

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)

}

func TestMap(t *testing.T) {
	data := &sync.Map{}

	group := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
} // mau membuat tipe data map tapi diakses dengan go routine sekaligus jangan pake bawaan map golang, tapi pake sync.Map karena sudah aman dari Race Condition
