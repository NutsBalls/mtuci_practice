package main

import "fmt"

type Transport interface {
	Move() string
	Stop() string
}

type Car struct {
	Brand string
	Speed int
}

type Bike struct {
	Brand string
}

type Train struct {
	Size         string
	NumsOfWagons int
}

func (c Car) Move() string {
	return fmt.Sprintf("Автомобиль марки %s едет по дороге со скоростью %d", c.Brand, c.Speed)
}

func (c Car) Stop() string {
	return "Автомобиль остановился"
}

func (b Bike) Move() string {
	return fmt.Sprintf("Едет велосипед модели %s", b.Brand)
}

func (b Bike) Stop() string {
	return "Велосипед остановлен"
}

func (t Train) Move() string {
	return fmt.Sprintf("%s поезд движется по рельсам. У него %d вагонов", t.Size, t.NumsOfWagons)
}

func (t Train) Stop() string {
	return "Поезд прибыл на станцию"
}

func goTransport(t Transport) {
	fmt.Println(t.Move())
	fmt.Println(t.Stop())
}

func main() {
	car := Car{Brand: "BMW", Speed: 150}
	bike := Bike{Brand: "Stels"}
	train := Train{Size: "Большой", NumsOfWagons: 20}

	goTransport(car)
	goTransport(bike)
	goTransport(train)
}
