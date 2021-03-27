package definitii

import (
	"testing"
)

func TestFormat(t *testing.T) {
	cuvant, _ := FetchCuvant("test")
	if cuvant.Definitions[0].format() != "TEST^2, teste, s. n. Înveliș calcaros sau chitinos care protejează corpul unor moluște, crustacee etc. - Din fr. test." {
		t.Error("Unexpected text")
	}
	if cuvant.Definitions[0].formatContext() != "Înveliș calcaros sau chitinos care protejează corpul unor moluște, crustacee etc." {
		t.Error("Unexpected text")
	}
}