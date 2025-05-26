package repository

import (
	"testing"
)

func TestAddPerson(t *testing.T) {
	persons := []Person{}

	person := Person{
		Name:        "Dupont",
		FirstName:   "Jean",
		PhoneNumber: "0123456789",
	}

	persons = AddPerson(persons, person)

	if len(persons) != 1 {
		t.Errorf("Expected 1 person in the list, got %d", len(persons))
	}

	if persons[0].Name != "Dupont" || persons[0].FirstName != "Jean" || persons[0].PhoneNumber != "0123456789" {
		t.Errorf("Person data doesn't match expected values")
	}

	persons = AddPerson(persons, person)

	if len(persons) != 1 {
		t.Errorf("Expected still 1 person in the list after adding duplicate, got %d", len(persons))
	}
}

func TestDeletePerson(t *testing.T) {
	person1 := Person{
		Name:        "Dupont",
		FirstName:   "Jean",
		PhoneNumber: "0123456789",
	}

	person2 := Person{
		Name:        "Martin",
		FirstName:   "Sophie",
		PhoneNumber: "0987654321",
	}

	persons := []Person{person1, person2}

	if len(persons) != 2 {
		t.Errorf("Expected 2 persons in the list, got %d", len(persons))
	}

	persons = DeletePerson(persons, person1)

	if len(persons) != 1 {
		t.Errorf("Expected 1 person in the list after deletion, got %d", len(persons))
	}

	if persons[0].Name != "Martin" || persons[0].FirstName != "Sophie" {
		t.Errorf("Wrong person was deleted")
	}

	personNotInList := Person{
		Name:        "Inconnu",
		FirstName:   "Pierre",
		PhoneNumber: "0000000000",
	}

	persons = DeletePerson(persons, personNotInList)

	if len(persons) != 1 {
		t.Errorf("Expected list to remain unchanged when deleting non-existent person")
	}
}
