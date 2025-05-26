package main

import (
	"github.com/SwannMartineau/go/tp1-note/repository"
)

func main() {
	person1 := repository.Person{
		Name:        "Dupont",
		FirstName:   "Jean",
		PhoneNumber: "0123456789",
	}

	person2 := repository.Person{
		Name:        "Martin",
		FirstName:   "Sophie",
		PhoneNumber: "0987654321",
	}

	persons := []repository.Person{person1, person2}

	repository.ListRepository(persons)

	persons2 := []repository.Person{}

	repository.ListRepository(persons2)

	persons2 = repository.AddPerson(persons2, person1)

	repository.ListRepository(persons2)

	persons2 = repository.AddPerson(persons2, person1)

	repository.ListRepository(persons2)

	persons2 = repository.DeletePerson(persons2, person1)

	repository.ListRepository(persons2)

	persons = repository.ModifyPerson(persons, person1.Name, person1.FirstName, "66666666")

	repository.ListRepository(persons)

	repository.SearchPerson(persons, person2)

	repository.SearchPerson(persons, person1)
}
