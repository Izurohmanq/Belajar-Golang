package belajargoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			time.Sleep(5 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GOroutine", totalGoroutine)
}
func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			time.Sleep(5 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	// -1 itu pengen liat default threadnya
	// kata pa eko mendingan horizontal scalling untuk thread
	// gak usah khawatir, karena udah ada go scheduler
	fmt.Println("total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GOroutine", totalGoroutine)
}
