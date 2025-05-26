package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English":               {lang: "en", want: "Hello world"},
		"French":                {lang: "fr", want: "Bonjour le monde"},
		"Empty":                 {lang: "", want: "Language not supported : \"\""},
		"Spanish not supported": {lang: "es", want: "Language not supported"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			greeting := greet(test.lang)
			if greeting != test.want {
				t.Errorf("Expected %s, got %s", test.want, greeting)
			}
		})
	}
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
