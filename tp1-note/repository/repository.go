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

func ListRepository(ps []Person) []Person {
	if len(ps) == 0 {
		fmt.Println("Aucune personne dans la liste.")
		return ps
	}
	if len(ps) == 1 {
		fmt.Println("Une seule personne dans la liste.")
	} else {
		fmt.Println("Liste des personnes :")
	}
	for _, p := range ps {
		fmt.Printf("Nom: %s, Prénom: %s, Téléphone: %s\n", p.Name, p.FirstName, p.PhoneNumber)
	}
	return ps
}

func AddPerson(persons []Person, person Person) []Person {
	for _, p := range persons {
		if p.Name == person.Name && p.FirstName == person.FirstName && p.PhoneNumber == person.PhoneNumber {
			fmt.Println("Cette personne existe déjà.")
			return persons
		}
	}
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
