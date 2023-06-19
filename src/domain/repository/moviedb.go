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
	movie5 := newMovie("5", "Acorralado", "Acción")
	movie6 := newMovie("6", "Aelita", "Aventura")
	movie7 := newMovie("7", "Buscando a Nemo", "Animación")
	movie8 := newMovie("8", "A Roma con Amor", "Comedia")
	movie9 := newMovie("9", "Metropia", "Animación")
	movie10 := newMovie("10", "Abel", "Comedia")
	movie11 := newMovie("11", "Apocalypto", "Aventura")
	movie12 := newMovie("12", "Billy Elliot", "Drama")
	movie13 := newMovie("13", "Ex Machina", "Ciencia Ficción")
	movie14 := newMovie("14", "Matrix", "Ciencia Ficción")
	movie15 := newMovie("15", "Akira", "Animación")
	movie16 := newMovie("16", "La Visita", "Terror")
	movie17 := newMovie("17", "Alien", "Terror")
	movie18 := newMovie("18", "Corre", "Drama")
	movie19 := newMovie("19", "La Monja", "Terror")
	movie20 := newMovie("20", "John Carter", "Aventura")

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
