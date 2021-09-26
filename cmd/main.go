package main

import (
	"fmt"
	"time"

	"./sensor"
)

func main() {
	fmt.Println("go")
	s := sensor.NewSensor()
	for {
		s.getmeasurement()
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
