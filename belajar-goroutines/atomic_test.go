package belajargoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1) //mau memanipulasi data yg tipe datanya primitif seperti angka, tidak usah menggunakan mutex pake atomic
				// kecuali datanya struct baru cocoknya pake mutex
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
