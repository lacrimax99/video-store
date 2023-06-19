package repository

import (
	"time"

	"github.com/finalproject/src/domain/model"
)

/*
Load in memory a map with rental information where key is id
@return map
*/
func GetAllRental() map[string]model.Rental {
	var t time.Time
	rent1 := newRental("1", "1", "1", time.Now().AddDate(0, 0, -1), time.Now(), 100)
	rent2 := newRental("2", "2", "2", time.Now().AddDate(0, 0, -2), time.Now(), 200)
	rent3 := newRental("3", "3", "5", time.Now().AddDate(0, 0, -1), time.Now(), 300)
	rent4 := newRental("4", "4", "6", time.Now().AddDate(0, 0, -3), t, 400)
	rent5 := newRental("5", "5", "9", time.Now().AddDate(0, 0, -1), t, 500)

	return map[string]model.Rental{
		rent1.Id: rent1,
		rent2.Id: rent2,
		rent3.Id: rent3,
		rent4.Id: rent4,
		rent5.Id: rent5,
	}
}

/*
Create new instance from rental
@return rental
*/
func newRental(id string, idUser string, idMovie string, rentDate time.Time, deliveryDate time.Time, price int) model.Rental {
	return model.Rental{
		Id:           id,
		IdUser:       idUser,
		IdMovie:      idMovie,
		RentDate:     rentDate,
		DeliveryDate: deliveryDate,
		Price:        price,
	}
}
