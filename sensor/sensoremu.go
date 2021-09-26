package sensor

import (
	"fmt"
	"math/rand"
)

//Sensor type of date
type Sensor struct {
	value int
}

func (s *Sensor) getmeasurement() {
	fmt.Println("get measurement sensor")
	s.value = rand.Intn(40)
}

//NewSensor generate new sensor reference
func NewSensor() *Sensor {
	return &Sensor{value: 0}
}
