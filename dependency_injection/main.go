package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func Greet(writer io.Writer, word string) {
	fmt.Fprintf(writer, "Hello, %s \n", word)
}

func Countdown(writer io.Writer, count int, sleeper SpySleeper) {
	for i := count; i >= 1; i-- {
		fmt.Fprintf(writer, "%d \n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(writer, "Go!")
}

func main() {
	Greet(os.Stdout, "Elodie")
	//Countdown(os.Stdout, 3)
}
