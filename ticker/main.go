package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 10
	dur := 10
	timer := time.NewTimer(time.Duration(duration) * time.Second)
	tick := time.Tick(time.Duration(duration))

	for {
		select {
		case <-timer.C:
			return
		case <-tick:
			fmt.Println(dur)
			dur--

		}
	}
}
