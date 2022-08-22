package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// basic channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Send Data to Channel"
		fmt.Println("Successfully send data to channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(1 * time.Second)
}

// channel as a parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Data Channel as a Parameter"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// channel in and out
func OnlyInChannel(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Param Data Channel"
}

func OnlyOutChannel(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyInChannel(channel)
	go OnlyOutChannel(channel)

	time.Sleep(5 * time.Second)
}

//buffer channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Data 1"
		channel <- "Data 2"
		channel <- "Data 3"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done!!!")
}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data ", data)
	}

	fmt.Println("Done!!!")
}

// select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2 ", data)
			counter++
		default:
			fmt.Println("Wait Received Data!!!")
		}

		if counter == 2 {
			break
		}
	}
}
