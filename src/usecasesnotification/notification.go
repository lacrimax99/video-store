package usecasesnotification

import "fmt"

type INotification interface {
	SendEmail()
	SendWhatsapp()
}

type notification struct {
	Platform  string
	Template  string
	MovieName string
}

func (n *notification) SendEmail() {
	fmt.Printf("Email")
}

func (n *notification) SendWhatsapp() {
	fmt.Println("Whatsapp")
}
