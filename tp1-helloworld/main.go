package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SwannMartineau/go/tp1-helloworld/greeter"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "Required language")
	flag.Parse()
	// var greeting string = "Hello world"
	greeting, err := greeter.Greet(greeter.Language(lang))
	if err != nil {
		log.Fatalf("Error greeting: %v", err)
	}
	fmt.Println(greeting)
}
