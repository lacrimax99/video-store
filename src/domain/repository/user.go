package repository

import "github.com/finalproject/src/domain/model"

/*
Load in memory a map with user information where key is id user
@return map
*/
func GetAllUser() map[string]model.User {
	user1 := newUser("1", "Alex", "alex@gmail.com", "7895522001")
	user2 := newUser("2", "Alexa", "alexander@gmail.com", "7895522002")
	user3 := newUser("3", "Fred", "fred@gmail.com", "7895522003")
	user4 := newUser("4", "Esteban", "esteban@gmail.com", "7895522004")
	user5 := newUser("5", "Sam", "sam@gmail.com", "7895522005")

	return map[string]model.User{
		user1.Id: user1,
		user2.Id: user2,
		user3.Id: user3,
		user4.Id: user4,
		user5.Id: user5,
	}
}

/*
Create new instance from user
@arrayValues array with user attributes
@return user
*/
func newUser(arrayValues ...string) model.User {
	return model.User{
		Id:        arrayValues[0],
		Name:      arrayValues[1],
		Email:     arrayValues[2],
		Cellphone: arrayValues[3],
	}
}
