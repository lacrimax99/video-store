package factory

import "fmt"

const AWS = "AWS"
const TWILIO = "TWILIO"

/*
NewNotification is a factory function that creates the specified notification
@param platform: Platform notification
@param movieName: Movie name
@param email: Email for sending notification
@param cellPhone: cellPhone for sending notification
*/
func NewNotification(platform string, movieName string, email string, cellPhone string) (INotification, error) {

	if platform == AWS {
		return createAws(platform, movieName, email, cellPhone), nil
	}
	if platform == TWILIO {
		return createTwilio(platform, movieName, email, cellPhone), nil
	}

	return nil, fmt.Errorf("No such notification platform")

}
