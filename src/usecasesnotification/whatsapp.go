package usecasesnotification

type whatsapp struct {
	notification
}

func createWhatsapp(platform string, template string, movieName string) INotification {
	return &whatsapp{
		notification: notification{
			Platform:  platform,
			Template:  template,
			MovieName: movieName,
		},
	}
}
