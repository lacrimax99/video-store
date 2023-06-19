package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/finalproject/src/domain/model"
	"github.com/finalproject/src/domain/repository"
	usecases "github.com/finalproject/src/usecases"
)

//Map
var mapMovies map[string]model.Movie
var mapUsers map[string]model.User
var mapRental map[string]model.Rental

/*
Main package containing the main function
*/
func main() {
	var opt int
	printLabel()

	_, err := fmt.Scan(&opt)
	if err != nil || opt < 1 || opt > 2 {
		fmt.Printf("Error al leer la opción: %v", err)
		os.Exit(1)
	}

	mapMovies = repository.GetAllMovies()
	mapUsers = repository.GetAllUser()
	mapRental = repository.GetAllRental()

	usecases.GetMoviesDelivery(mapMovies, mapUsers, mapRental)
	usecases.GetMoviesRent(mapMovies, mapUsers, mapRental)
	createGoRoutineNotification(getPlatform(opt))

}

/*
Get platform option
@param opt: platform option
@return platform name
*/
func getPlatform(opt int) string {
	if opt == 1 {
		return "AWS"
	}
	return "TWILIO"
}

/*
Print system message
*/
func printLabel() {
	fmt.Println("Digite plataforma a enviar notificación")
	fmt.Println("1 -> AWS")
	fmt.Println("2 -> TWILIO")
}

/*
Process a system notification to send to the customers
@platform platform notification
*/
func createGoRoutineNotification(platform string) {
	var wg sync.WaitGroup
	wg.Add(1)
	usecases.ProccessNotification(&wg, platform, mapMovies, mapUsers, mapRental)
	wg.Wait()
}
