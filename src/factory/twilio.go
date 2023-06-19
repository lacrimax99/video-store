package factory

import "fmt"

/*
Define a twilio structure and embed the inotification interface
*/
type twilio struct {
	notification
}

/*
Create new object from twilio for sending notification
@param platform: Platform notification
@param movieName: Movie name
@param email: Email for sending notification
@param cellPhone: cellPhone for sending notification
*/
func createTwilio(platform string, movieName string, email string, cellPhone string) INotification {
	return &twilio{
		notification: notification{
			Platform:  platform,
			MovieName: movieName,
			Email:     email,
			Cellphone: cellPhone,
		},
	}
}

/*
Send a email notification with twilio pointer 
*/
func (t *twilio) SendEmail(template string) {
	// :TODO implement configuration to email twilio
	fmt.Println("---------------------------------------")
	fmt.Printf("Email message to %s sent by %s \n", t.Email, t.Platform)
	fmt.Printf(template, t.MovieName)
	fmt.Println("---------------------------------------")
}

/*
Send a whatsapp notification with twilio pointer 
*/
func (t *twilio) SendWhatsapp(template string) {
	// :TODO implement configuration to whatsapp twilio
	fmt.Println("---------------------------------------")
	fmt.Printf("Whatsapp message to %s send by %s \n", t.Cellphone, t.Platform)
	fmt.Printf(template, t.MovieName)
	fmt.Println("---------------------------------------")
}
