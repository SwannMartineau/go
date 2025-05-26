package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	lang := language("en")
	expected := "Hello world"
	greeting := greet(lang)
	if greeting != expected {
		t.Errorf("Expected %s, got %s", expected, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	expected := "Bonjour le monde"
	greeting := greet(lang)
	if greeting != expected {
		t.Errorf("Expected %s, got %s", expected, greeting)
	}
}

func TestGreet_Spanish(t *testing.T) {
	lang := language("es")
	expected := "Language not supported"
	greeting := greet(lang)
	if greeting != expected {
		t.Errorf("Expected %s, got %s", expected, greeting)
	}
}
