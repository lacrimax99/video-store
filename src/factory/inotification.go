package factory

/*
Interface for sending notification
*/
type INotification interface {
	SendEmail(template string)
	SendWhatsapp(template string)
}
