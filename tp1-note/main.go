package main

import (
	"flag"
	"fmt"

	"github.com/SwannMartineau/go/tp1-note/repository"
)

func main() {
	actionPtr := flag.String("action", "", "Action à effectuer (ajouter, rechercher, lister)")
	nomPtr := flag.String("nom", "", "Nom de famille de la personne")
	prenomPtr := flag.String("prenom", "", "Prénom de la personne")
	telPtr := flag.String("tel", "", "Numéro de téléphone")

	flag.Parse()

	persons := loadDirectory()

	newPerson := repository.Person{
		Name:        *nomPtr,
		FirstName:   *prenomPtr,
		PhoneNumber: *telPtr,
	}

	switch *actionPtr {
	case "ajouter":
		if *nomPtr == "" {
			fmt.Println("Erreur: Le nom est requis pour ajouter une personne")
		}

		persons = repository.AddPerson(persons, newPerson)
		fmt.Printf("Personne ajoutée: %s %s, Tél: %s\n", *prenomPtr, *nomPtr, *telPtr)
		repository.ListRepository(persons)

	case "rechercher":
		if *nomPtr == "" && *prenomPtr == "" {
			fmt.Println("Erreur: Le nom ou le prénom est requis pour rechercher une personne")
		}

		repository.SearchPerson(persons, newPerson)

	case "lister":
		repository.ListRepository(persons)

	case "supprimer":
		persons = repository.DeletePerson(persons, newPerson)
		repository.ListRepository(persons)

	case "modifier":
		persons = repository.ModifyPerson(persons, newPerson.Name, newPerson.FirstName, *telPtr)
		repository.ListRepository(persons)

	default:
		fmt.Println("Action non reconnue. Utilisez 'ajouter', 'rechercher', 'lister', 'supprimer' ou 'modifier'")
	}
}

func loadDirectory() []repository.Person {
	return []repository.Person{
		{Name: "Dupont", FirstName: "Jean", PhoneNumber: "0123456789"},
		{Name: "Martin", FirstName: "Sophie", PhoneNumber: "0987654321"},
		{Name: "Durand", FirstName: "Alice", PhoneNumber: "0654321987"},
	}
}
