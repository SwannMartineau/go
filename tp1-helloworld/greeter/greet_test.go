package greeter

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang        Language
		want        string
		expectError bool
		errValue    string
	}

	var tests = map[string]testCase{
		"English": {lang: "en", want: "Hello world"},
		"French":  {lang: "fr", want: "Bonjour le monde"},
		"Empty":   {lang: "", want: "", expectError: true, errValue: ""},
		"Spanish": {lang: "es", want: "Language not supported"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			greeting, err := Greet(test.lang)

			if test.expectError {
				if err == nil {
					t.Errorf("Expected error for language %q, got nil", test.lang)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for language %q, got: %v", test.lang, err)
				}
				if greeting != test.want {
					t.Errorf("Expected greeting %q, got %q", test.want, greeting)
				}
			}
		})
	}
}

func TestGreet_English(t *testing.T) {
	lang := Language("en")
	expected := "Hello world"
	greeting, err := Greet(lang)
	if greeting != expected {
		t.Errorf("Expected %q, got %q", expected, greeting)
	}
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGreet_French(t *testing.T) {
	lang := Language("fr")
	expected := "Bonjour le monde"
	greeting, err := Greet(lang)
	if greeting != expected {
		t.Errorf("Expected %q, got %q", expected, greeting)
	}
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGreet_Spanish(t *testing.T) {
	lang := Language("es")
	expected := "Language not supported"
	greeting, err := Greet(lang)
	if greeting != expected {
		t.Errorf("Expected %q, got %q", expected, greeting)
	}
	if err != nil {
		t.Errorf("Expected no error for Spanish, got %v", err)
	}
}
