package usecasesnotification

import "fmt"

func newNotification(platform string, template string, movieName string) (INotification, error) {

	if platform == "AWS" {
		return createEmail(platform, template, movieName), nil
	}
	if platform == "TWILIO" {
		return createWhatsapp(platform, template, movieName), nil
	}

	return nil, fmt.Errorf("No such notification platform")

}
