package repository

import (
	"github.com/finalproject/src/domain/model"
)

/*
Load in memory a map with movies information where key is id movie
@return map
*/
func GetAllMovies() map[string]model.Movie {
	movie1 := newMovie("1", "Avatar", "Aventura")
	movie2 := newMovie("2", "El Aro", "Terror")
	movie3 := newMovie("3", "Aquaman", "Acción")
	movie4 := newMovie("4", "Spiderman", "Acción")
	movie5 := newMovie("5", "a", "Acción")
	movie6 := newMovie("6", "b", "Acción")
	movie7 := newMovie("7", "c", "Acción")
	movie8 := newMovie("8", "d", "Acción")
	movie9 := newMovie("9", "e", "Acción")
	movie10 := newMovie("10", "f", "Acción")
	movie11 := newMovie("11", "g", "Acción")
	movie12 := newMovie("12", "Thor: El Mundo Oscuro", "Acción")
	movie13 := newMovie("13", "Thor: El Mundo Oscuro", "Acción")
	movie14 := newMovie("14", "Thor: El Mundo Oscuro", "Acción")
	movie15 := newMovie("15", "Thor: El Mundo Oscuro", "Acción")
	movie16 := newMovie("16", "Thor: El Mundo Oscuro", "Acción")
	movie17 := newMovie("17", "Thor: El Mundo Oscuro", "Acción")
	movie18 := newMovie("18", "Thor: El Mundo Oscuro", "Acción")
	movie19 := newMovie("19", "Thor: El Mundo Oscuro", "Acción")
	movie20 := newMovie("20", "Thor: El Mundo Oscuro", "Acción")

	return map[string]model.Movie{
		movie1.Id:  movie1,
		movie2.Id:  movie2,
		movie3.Id:  movie3,
		movie4.Id:  movie4,
		movie5.Id:  movie5,
		movie6.Id:  movie6,
		movie7.Id:  movie7,
		movie8.Id:  movie8,
		movie9.Id:  movie9,
		movie10.Id: movie10,
		movie11.Id: movie11,
		movie12.Id: movie12,
		movie13.Id: movie13,
		movie14.Id: movie14,
		movie15.Id: movie15,
		movie16.Id: movie16,
		movie17.Id: movie17,
		movie18.Id: movie18,
		movie19.Id: movie19,
		movie20.Id: movie20,
	}
}

/*
Create new instance from movie
@arrayValues array with movie attributes
@return movie
*/
func newMovie(arrayValues ...string) model.Movie {
	return model.Movie{
		Id:     arrayValues[0],
		Name:   arrayValues[1],
		Gender: arrayValues[2],
	}
}
