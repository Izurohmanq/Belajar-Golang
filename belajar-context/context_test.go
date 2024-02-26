package belajarcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

//context.background() => membuat context kosong. Tidak pernah dibatalkan
//tidak pernah timeout, dan tidak memiliki value apapun.
//biasanya digunakan dalam function atau test

//context.TOTO(), belum jelas arahnya ke mana

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextValue(t *testing.T) {
	contextA := context.Background()

	// contextB := context.WithValue(parent, "value", "value")
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f")) //dapat
	fmt.Println(contextF.Value("c")) //dapat karena parentnya C jadi weh ada
	fmt.Println(contextF.Value("b")) //gada

	fmt.Println(contextA.Value("b")) //gada
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}

		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second) //simulasi aja ditambah 2 detik, kalau kecepatn entar sama aja

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

//ceritanya ada goroutine leak, dengan menggunakan context jadi tidak ada goroutine leak

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel() // mengirim sinyal cancel ke context

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second) //simulasi aja ditambah 2 detik, kalau kecepatn entar sama aja

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) //harus pake time.Now()
	defer cancel()                                                             // mengirim sinyal cancel ke context

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second) //simulasi aja ditambah 2 detik, kalau kecepatn entar sama aja

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
} // biasanya Deadline itu fix gitu waktunya mau kapan
