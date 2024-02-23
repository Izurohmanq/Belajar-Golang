package belajargoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() //harus nambahin ini kalau ingin menggunakan Waitgroup

	group.Add(1)

	fmt.Println("Hoolla")

	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}