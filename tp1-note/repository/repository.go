package repository

import "fmt"

type Person struct {
	Name        string
	FirstName   string
	PhoneNumber string
}

type Repository struct {
	Persons []Person
}

func ListRepository(persons []Person) []Person {
	if len(persons) == 0 {
		fmt.Println("Aucune personne dans la liste.")
		return persons
	}
	fmt.Println("Liste des personnes :")
	for _, person := range persons {
		fmt.Printf("Nom: %s, Prénom: %s, Téléphone: %s\n", person.Name, person.FirstName, person.PhoneNumber)
	}
	return persons
}

func AddPerson(persons []Person, person Person) []Person {
	return append(persons, person)
}

func DeletePerson(persons []Person, person Person) []Person {
	return persons
}

func ModifyPerson(persons []Person, person Person) []Person {
	return persons
}

func SearchPerson(persons []Person, person Person) []Person {
	return persons
}
