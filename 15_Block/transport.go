package main

import "fmt"

type Transport interface {
	Move() string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Move() string {
	return fmt.Sprintf("Автомобиль %s едет по дороге", c.Model)
}

func (c Car) Stop() string {
	return fmt.Sprintf("Автомобиль %s остановился", c.Model)
}

type Airplane struct {
	FlightNumber string
}

func (a Airplane) Move() string {
	return fmt.Sprintf("Самолет рейса %s взлетает", a.FlightNumber)
}

func (a Airplane) Stop() string {
	return fmt.Sprintf("Самолет рейса %s приземлился", a.FlightNumber)
}

func testTransport(t Transport) {
	fmt.Println(t.Move())
	fmt.Println(t.Stop())
}

func main() {
	car := Car{Model: "МАШИНА"}
	plane := Airplane{FlightNumber: "1"}

	testTransport(car)

	testTransport(plane)
	transports := []Transport{car, plane}
	for _, t := range transports {
		fmt.Println(t.Move())
	}
}
