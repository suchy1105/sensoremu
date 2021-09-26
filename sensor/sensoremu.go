package sensor

import (
	"math/rand"
)

//Sensor type of date
type Sensor struct {
	Value int
}

//Getmeasurement method
func (s *Sensor) Getmeasurement() {
	s.Value = rand.Intn(40)
}

//NewSensor generate new sensor reference
func NewSensor() *Sensor {
	return &Sensor{Value: 0}
}
