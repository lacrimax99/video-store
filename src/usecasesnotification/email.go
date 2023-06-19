package usecasesnotification

type email struct {
	notification
}

func createEmail(platform string, template string, movieName string) INotification {
	return &email{
		notification: notification{
			Platform:  platform,
			Template:  template,
			MovieName: movieName,
		},
	}
}
