package belajargoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second) //ingin mengirim suatu kejadian tapi nanti gitu di masa yang akan datang, cuman ini cara traditional
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}
func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second) //ingin mengirim suatu kejadian tapi nanti gitu di masa yang akan datang, cuman ini cara cepat
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())

	group.Wait()
}

//intinya timer ini digunakan untuk delay
