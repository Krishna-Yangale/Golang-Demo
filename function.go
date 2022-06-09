package main

import "fmt"

func main() {

	// rupees := convertDollar(54)
	// fmt.Println(rupees)

	fmt.Println(convertDollar(50))

}

func convertDollar(dollar float64) float64 {
	return dollar * 72
}
