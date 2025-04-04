package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	doWork()
}

func doWork() {
	for i := 0; i < 10; i++ {
		busyWork()
		time.Sleep(50 * time.Millisecond) // shorter sleep to keep CPU busy
	}
}

var result int // global variable to hold the computation result
func busyWork() {
	data := make([]string, 3)
	sum := 0
	// Loop for a significant number of iterations
	for i := 0; i < 10000000; i++ {
		// Perform some arithmetic operations
		sum += (i * i) % 100
		// Append to the slice to simulate memory allocation
		data = append(data, "magical pprof time")
		// Simulate some string operations
		if i%1000000 == 0 {
			data = append(data, "hello world")
		}
	}
	result = sum // store the result to avoid optimization
}
