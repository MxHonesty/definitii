package definitii

import (
	"testing"
)

func TestFormat(t *testing.T) {
	cuvant, _ := FetchCuvant("test")
	if cuvant.Definitions[0].Format() != "TEST^2, teste, s. n. Înveliș calcaros sau chitinos care protejează corpul unor moluște, crustacee etc. - Din fr. test." {
		t.Error("Unexpected text")
	}
	if cuvant.Definitions[0].FormatContext() != "Înveliș calcaros sau chitinos care protejează corpul unor moluște, crustacee etc." {
		t.Error("Unexpected text")
	}
}