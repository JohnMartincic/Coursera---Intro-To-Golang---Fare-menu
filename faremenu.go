package main

import "fmt"

func main() {
	var origin, destination, cabinClass string
	const enteredMsg = "You've entered "

	fmt.Println("*** Silver Karma Airfare Calculator ***")

	validOriginEntered := false
	var originCity City
	var originErr error

	for !validOriginEntered {
		fmt.Println("Enter origin code: ")
		fmt.Scanln(&origin)

		originCity, originErr = getCityFromCode(origin)

		if originErr == nil {
			fmt.Println(enteredMsg + originCity.cityName)
			validOriginEntered = true
		} else {
			fmt.Println(originErr)
		}
	}

	validDestinationEntered := false
	var destinationCity City
	var destinationErr error

	for !validDestinationEntered {
		fmt.Println("Enter destination code: ")
		fmt.Scanln(&destination)

		destinationCity, destinationErr = getCityFromCode(destination)

		if destinationErr == nil {
			fmt.Println(enteredMsg + destinationCity.cityName)
			validDestinationEntered = true
		} else {
			fmt.Println(destinationErr)
		}
	}

	validCabinClassEntered := false
	var enteredCabinClass CabinClass
	var enteredCabinClassErr error

	for !validCabinClassEntered {
		fmt.Println("Enter cabin class code: ")
		fmt.Scanln(&cabinClass)

		enteredCabinClass, enteredCabinClassErr = getCabinClassFromCode(cabinClass)

		if enteredCabinClassErr == nil {
			fmt.Println(enteredMsg + enteredCabinClass.className + " class")
			validCabinClassEntered = true
		} else {
			fmt.Println(enteredCabinClassErr)
		}
	}
}
