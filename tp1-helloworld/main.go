package main

import "fmt"

func main() {
	// var gtreeting string = "Hello world"
	greeting := greet("en")
	fmt.Println(greeting)
}

type language string

//var m map[string]int
//m := make(map[string]int)
/* m := map[string]int{
    "a": 1,
    "b": 2
} */
// v := m["a"]
// m["c"] = 3 // ajouter une clé
// v, ok := m["c"]
// delete(m, "c") // supprimer une clé
// length := len(m) // taille de la map

var phrasebook = map[language]string{
	"fr": "Bonjour le monde",
	"en": "Hello world",
	"es": "Language not supported",
}

func greet(l language) string {
	// var gtreeting string = "Hello world"
	greeting, exists := phrasebook[l]
	if !exists {
		return fmt.Sprintf("Language not supported : %q", l)
	}
	return greeting
}
