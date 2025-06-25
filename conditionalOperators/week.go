package main

import "fmt"

func main() {

	var dayOfTheWeek int

	fmt.Scan(&dayOfTheWeek)

	switch dayOfTheWeek {

	case 1:
		fmt.Println("День недели понедельник")
	case 2:
		fmt.Println("День недели вторник")
	case 3:
		fmt.Println("День недели среда")
	case 4:
		fmt.Println("День недели четверг")
	case 5:
		fmt.Println("День недели пятница")
	case 6:
		fmt.Println("День недели суббота")
	case 7:
		fmt.Println("День недели воскресенье")
	}

}
