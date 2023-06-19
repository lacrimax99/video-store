package factory

import "fmt"

/*
Define a aws structure and embed the inotification interface
*/
type aws struct {
	notification
}

/*
Create new object from AWS for sending notification
@param platform: Platform notification
@param movieName: Movie name
@param email: Email for sending notification
@param cellPhone: cellPhone for sending notification
*/
func createAws(platform string, movieName string, email string, cellPhone string) INotification {
	return &aws{
		notification: notification{
			Platform:  platform,
			MovieName: movieName,
			Email:     email,
			Cellphone: cellPhone,
		},
	}
}

/*
Send a email notification with aws pointer
*/
func (a *aws) SendEmail(template string) {
	// :TODO implement configuration to email aws
	fmt.Println("---------------------------------------")
	fmt.Printf("Email message to %s sent by %s \n", a.Email, a.Platform)
	fmt.Printf(template, a.MovieName)
	fmt.Println("---------------------------------------")
}

/*
Send a whatsapp notification with aws pointer
*/
func (a *aws) SendWhatsapp(template string) {
	// :TODO implement configuration to whatsapp aws
	fmt.Println("---------------------------------------")
	fmt.Printf("Whatsapp message to %s send by %s \n", a.Cellphone, a.Platform)
	fmt.Printf(template, a.MovieName)
	fmt.Println("---------------------------------------")
}
