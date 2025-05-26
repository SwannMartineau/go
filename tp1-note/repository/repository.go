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
	if len(persons) == 1 {
		fmt.Println("Une seule personne dans la liste.")
	} else {
		fmt.Println("Liste des personnes :")
	}
	for _, person := range persons {
		fmt.Printf("Nom: %s, prénom: %s, Téléphone: %s\n", person.Name, person.FirstName, person.PhoneNumber)
	}
	return persons
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
	for i, p := range persons {
		if p.Name == person.Name && p.FirstName == person.FirstName && p.PhoneNumber == person.PhoneNumber {
			//Slice jusqu'à l'index i (exclu) et prendre i+1 et spread le reste en arguments individuels
			persons = append(persons[:i], persons[i+1:]...)
			return persons
		}
	}
	return persons
}

func ModifyPerson(persons []Person, lastName string, firstName string, phoneNumber string) []Person {
	found := false
	for i, p := range persons {
		if p.Name == lastName && p.FirstName == firstName {
			persons[i].PhoneNumber = phoneNumber
			found = true
		}
	}
	if !found {
		fmt.Println("Cette personne n'existe pas")
	} else {
		fmt.Printf("Le numéro de téléphone de %s %s a été modifié\n", firstName, lastName)
	}
	return persons
}

func SearchPerson(persons []Person, person Person) []Person {
	for i, p := range persons {
		if p.Name == person.Name && p.FirstName == person.FirstName && p.PhoneNumber == person.PhoneNumber {
			fmt.Printf("Personne trouvé à l'index %d ! Nom: %s, prénom: %s, Téléphone: %s\n", i, p.Name, p.FirstName, p.PhoneNumber)
			return persons
		}
	}
	fmt.Println("Cette personne n'existe pas.")
	return persons
}
