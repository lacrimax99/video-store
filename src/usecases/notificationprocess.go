package usecases

import (
	"fmt"
	"sync"
	"time"

	"github.com/finalproject/src/domain/model"
	factory "github.com/finalproject/src/factory"
)

const TEMPLATE_EMAIL = "Estimado cliente,\nEl prestamo de la pelicula %s ha vencido,\nEsperamos su respuesta\n"
const TEMPLATE_WHATSAPP = "Estimado cliente, el prestamo de la pelicula %s ha vencido.\nEn caso de alguna información adicional contactar con nuestros accesores\n"

func ProccessNotification(wg *sync.WaitGroup, platform string, mapMovie map[string]model.Movie, mapUser map[string]model.User, mapRental map[string]model.Rental) {
	fmt.Println("NOTIFICATIONS")
	timeNow := time.Now()
	for _, rental := range mapRental {
		if rental.DeliveryDate.IsZero() && timeNow.Sub(rental.RentDate).Hours() >= 24 {
			notification, _ := factory.NewNotification(platform, mapMovie[rental.IdMovie].Name, mapUser[rental.IdUser].Email, mapUser[rental.IdUser].Cellphone)
			notification.SendEmail(TEMPLATE_EMAIL)
			notification.SendWhatsapp(TEMPLATE_WHATSAPP)
		}
	}
	defer wg.Done()
}
