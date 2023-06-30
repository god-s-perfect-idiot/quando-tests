package main

import "fmt"

type Callback func()

func ProcessCallbacks(callbacks []Callback) {
	for _, callback := range callbacks {
		callback()
	}
}

// Example callback functions
func Callback1() {
	fmt.Println("Callback 1")
}

func Callback2() {
	fmt.Println("Callback 2")
}

func main() {
	// Create a slice of 100,000 callbacks
	callbacks := make([]Callback, 100000)
	for i := range callbacks {
		if i%2 == 0 {
			callbacks[i] = Callback1
		} else {
			callbacks[i] = Callback2
		}
	}

	// Process the callbacks
	ProcessCallbacks(callbacks)
}
