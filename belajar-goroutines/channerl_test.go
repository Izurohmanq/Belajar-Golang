package belajargoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // jangan lupa di-close

	//channel <- "Eko" // masukin data ke channel

	//data := <-channel // mengambil data dari channel terus masukkin ke variable data

	//fmt.Println(data)      // bisa gini
	//fmt.Println(<-channel) // bisa langsung

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Asep kayu"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Asep kayu"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) { // ini menerima data
	time.Sleep(2 * time.Second)
	channel <- "Asep kayu"
}
func OnlyOut(channel <-chan string) { // ini keluarin data
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
