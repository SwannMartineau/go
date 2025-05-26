package greeter

import "fmt"

type Language string

var phrasebook = map[Language]string{
	"fr": "Bonjour le monde",
	"en": "Hello world",
	"es": "Language not supported",
}

func Greet(l Language) (string, error) {
	greeting, exists := phrasebook[l]
	if !exists {
		return "", fmt.Errorf("unsupported language : %q", l)
	}
	return greeting, nil
}
