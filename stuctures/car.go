package main

import "fmt"

type Engine struct {
	Model      string
	HorsePower int
}

type Car struct {
	Brand  string
	Model  string
	Year   int
	Engine Engine
}

func infoAboutCar(c Car) {
	fmt.Printf("Это машина бренда %s модели %s, выпущенная в %d году имеет двигатель модели %s, с мощностью %d л. с.\n", c.Brand, c.Model, c.Year, c.Engine.Model, c.Engine.HorsePower)
}

func (c *Car) updateEngine(newModel string, newHorsePower int) {
	c.Engine.Model = newModel
	c.Engine.HorsePower = newHorsePower

	fmt.Printf("Двигатель был обновлен до модели %s, с мощностью %d л. с.\n", newModel, newHorsePower)
}

func main() {
	engine := Engine{Model: "v5", HorsePower: 110}
	car := Car{Brand: "Lada", Model: "Sport", Year: 2021, Engine: engine}

	infoAboutCar(car)

	car.updateEngine("v10", 300)
	infoAboutCar(car)
}
