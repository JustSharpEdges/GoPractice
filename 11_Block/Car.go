package main

import "fmt"

type Engine struct {
	Type    string
	Power   float64
	Volume  float64
	IsTurbo bool
}

func (e Engine) PrintEngineInfo() {
	fmt.Printf("Двигатель: %s, %.1f л, %.1f л.с.", e.Type, e.Volume, e.Power)
	if e.IsTurbo {
		fmt.Print(" (турбированный)")
	}
	fmt.Println()
}

type Car struct {
	Brand      string
	Model      string
	Year       int
	Mileage    int
	EngineInfo Engine
}

func (c Car) PrintCarInfo() {
	fmt.Printf("Автомобиль: %s %s %d года\n", c.Brand, c.Model, c.Year)
	fmt.Printf("Пробег: %d км\n", c.Mileage)
	c.EngineInfo.PrintEngineInfo()
}

func (c *Car) UpdateMileage(newMileage int) {
	if newMileage >= c.Mileage {
		c.Mileage = newMileage
		fmt.Printf("Пробег обновлен: %d км\n", c.Mileage)
	} else {
		fmt.Println("Новый пробег не может быть меньше текущего")
	}
}

func main() {
	engine := Engine{
		Type:    "бензиновый",
		Power:   950,
		Volume:  4.4,
		IsTurbo: true,
	}

	car := Car{
		Brand:      "AUDI",
		Model:      "RSQ8",
		Year:       2023,
		Mileage:    10000,
		EngineInfo: engine,
	}

	fmt.Println("=== Информация об автомобиле ===")
	car.PrintCarInfo()
	fmt.Println()

	fmt.Println("=== Обновление пробега ===")
	car.UpdateMileage(38000)
	car.UpdateMileage(30000)
	fmt.Println()

	fmt.Println("=== Обновленная информация ===")
	car.PrintCarInfo()
}
