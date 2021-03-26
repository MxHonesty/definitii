package definitii

import (
	"testing"
)

// Functie de test pentru functia getUrl.
func TestGetUrl(t *testing.T) {
	var cuvant string
	cuvant = "test"
	rezultat := getUrl(cuvant)
	if rezultat != "https://dexonline.ro/definitie/test/json" {
		t.Error("Url format incorect")
	}
}

func TestIsValid(t *testing.T) {
	defs := make([]Definitie, 1, 1)  // Exemplu corect
	defs = append(defs, Definitie{Text: "test"})
	var cuvant = Cuvant{
		Type:        "test",
		Word:        "test",
		Definitions: defs,
	}

	valid := cuvant.isValid()
	if valid == false {
		t.Error("Expected true, input is valid")
	}

	defsInvalid := make([]Definitie, 0, 1) // Lipsesc definitiile.
	cuvantInvalid := Cuvant{
		Definitions: defsInvalid,
		Type:        "test",
		Word:        "test",
	}
	valid2 := cuvantInvalid.isValid()
	if valid2 == true {
		t.Error("Expected false, input is invalid")
	}

	defsValid := make([]Definitie, 0, 10)
	defsValid = append(defsValid, Definitie{Text: "test"})

	cuvantFaraType := Cuvant{  // Formam un cuvant fara Type.
		Definitions: defsValid,
		Word: "test",
	}
	if cuvantFaraType.isValid() == true {
		t.Error("Expected false, input is invalid")
	}

	cuvantFaraWord := Cuvant{  // Formam un cuvant fara Word
		Definitions: defsValid,
		Type: "test",
	}
	if cuvantFaraWord.isValid() == true {
		t.Error("Expected false, input is invalid")
	}

}

func TestFetchCuvant(t *testing.T) {
	var cuvant string
	cuvant = "cuvant"

	rez, err := FetchCuvant(cuvant)
	if err != nil {
		t.Error("Unexpected Error")
	} else {
		if rez.isValid() != true {
			t.Error("Expected true, input is valid")
		}
		if rez.Word != "cuvant" {
			t.Error("wrong Word field")
		}
		if rez.Type != "searchResults" {
			t.Error("Wrong Type field")
		}
	}

	rez, err = FetchCuvant("ababababa")  // Cuvant inexistent
	if err == nil {
		t.Error("Expected error")
	} else if err.Error() != "Cuvantul nu a fost gasit" {
		t.Error("Expected error message not found.")
	}
}