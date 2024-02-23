package belajargoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop() //kudu distop
	}()

	for time := range ticker.C {
		fmt.Println(time)
	} //ini pasti deadlock
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	} //ini pasti muter terus, mau kejadian berulang pake Ticker, kalau kejadiannya ingin sekali saja pake yang Timer
}
