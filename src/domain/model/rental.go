package model

import "time"

/*
Rental entity
*/
type Rental struct {
	Id           string
	IdUser       string
	IdMovie      string
	RentDate     time.Time
	DeliveryDate time.Time
	Price        int
}
