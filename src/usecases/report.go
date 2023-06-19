package usecases

import (
	"fmt"
	"time"

	"github.com/finalproject/src/domain/model"
)

/*
seq returns another function
@return function closes over the variable i to form a closure
*/
func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
Report from movies delivery
*/
func GetMoviesDelivery(mapMovie map[string]model.Movie, mapUser map[string]model.User, mapRental map[string]model.Rental) {
	printHeadDelivery()

	next := seq()
	sum := 0
	for _, rental := range mapRental {
		if !rental.DeliveryDate.IsZero() {
			fmt.Printf("%d %20s  %21s %23s %24s %17d\n", next(), mapMovie[rental.IdMovie].Name, mapUser[rental.IdUser].Name, rental.RentDate.Format(time.DateOnly), rental.DeliveryDate.Format(time.DateOnly), rental.Price)
			sum += rental.Price
		}
	}
	printTotalDelivery(sum)
}

/*
Report from movies rent
*/
func GetMoviesRent(mapMovie map[string]model.Movie, mapUser map[string]model.User, mapRental map[string]model.Rental) {
	printHeadRent()

	next := seq()
	sum := 0
	for _, rental := range mapRental {
		if rental.DeliveryDate.IsZero() {
			fmt.Printf("%d %20s  %21s %23s %17d \n", next(), mapMovie[rental.IdMovie].Name, mapUser[rental.IdUser].Name, rental.RentDate.Format(time.DateOnly), rental.Price)
			sum += rental.Price
		}
	}
	printTotalRent(sum)
}

/*
Print head delivery of the report
*/
func printHeadDelivery() {
	fmt.Printf("\t\t\t\t Report Movies Delivery\n")
	fmt.Printf("Num \t\t Movie \t\t\t User \t\t    Rent Date \t\t Delivery Date \t\t   Price\n")
	fmt.Printf("================================================================================================================\n")
}

/*
Print total delivery
@param sum: total value of delivery
*/
func printTotalDelivery(sum int) {
	fmt.Printf("================================================================================================================\n")
	fmt.Printf("\t\t\t\t\t\t\t\t\t\t\t\t\t     %d", sum)
}

/*
Print head rent of the report
*/
func printHeadRent() {
	fmt.Printf("\n\n\t\t\t\t   Report Movies Rent \n")
	fmt.Printf("Num \t\t Movie \t\t\t User \t\t    Rent Date \t\t  Price \n")
	fmt.Printf("=======================================================================================\n")
}

/*
Print total rent
@param sum: total value of rent
*/
func printTotalRent(sum int) {
	fmt.Printf("=======================================================================================\n")
	fmt.Printf("\t\t\t\t\t\t\t\t\t\t    %d\n", sum)
}
