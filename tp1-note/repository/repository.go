package repository

type Person struct {
	Name        string
	FirstName   string
	PhoneNumber string
}

type Repository struct {
	persons []Person
}

func listRepository(persons []Person) []Person {
	return persons
}

func addPerson(persons []Person, person Person) []Person {
	return persons
}

func deletePerson(persons []Person, person Person) []Person {
	return persons
}

func modifyPerson(persons []Person, person Person) []Person {
	return persons
}

func searchPerson(persons []Person, person Person) []Person {
	return persons
}
