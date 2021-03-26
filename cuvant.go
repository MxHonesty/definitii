package definitii

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/** Patreaza date pentru un cuvant. */
type Cuvant struct {
	Type string `json:"type"`
	Word string `json:"word"`
	Definitions []Definitie `json:"definitions"` // Pastreaza lista de definitii
}

/* Formeaza URL-ul pentru cuvantul dat, respectand formatul de adresare al site-ului.
	cuvant - string cuvantul pentru care dorim sa formam url.
 */
func getUrl(cuvant string) string {
	return fmt.Sprintf("https://dexonline.ro/definitie/%s/json", cuvant)
}

/* Extrage definitiile cuvantului oferite prin https://dexonline.ro/definitie/
	si returneaza o reprezentare a lor sub forma de structura Cuvant.
	Daca extragerea a fost efectuata cu succes, err intoarce nil.
	Returneaza o eroare de tip CuvantNotFoundError daca nu se poate
	citi cuvantul sau nu se poate citi continutul paginii.
*/
func FetchCuvant(cuvant string) (rez Cuvant, err error) {
	url := getUrl(cuvant)

	resp, getErr := http.Get(url)
	if getErr != nil {
		return rez, &CuvantNotFoundError{} // Returnam eroarea de la Get.
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return rez, &CuvantNotFoundError{}
	}

	jsonErr := json.Unmarshal(body, &rez)
	if jsonErr != nil {
		return rez, &CuvantNotFoundError{}
	}

	if rez.isValid() == false {  // Daca nu a fost gasit cuvantul.
		return rez, &CuvantNotFoundError{}
	}

	return rez, nil
}

/** Determina daca cuvantul este valid.
	Returneaza true daca este valid, false daca nu este valid.
	Un cuvant este valid daca are cel putin o definitie, un field Type
	si un field Word nevid.
*/
func (cuv *Cuvant) isValid() (rez bool) {
	rez = true
	if cuv.Type == "" {
		rez = false
	}
	if cuv.Word == "" {
		rez = false
	}
	if len(cuv.Definitions) == 0 {
		rez = false
	}
	return
}