package definitii

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

/** Functia este un placeholder pentru
	exemplificarea structurii testelor
*/
func TestPlaceholder(t *testing.T) {
	// Request la url si parse la json. Arunca erori daca intervine ceva.
	url := "https://dexonline.ro/definitie/test/json"

	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()  // Close Body on function exit

	body, err := ioutil.ReadAll(resp.Body)  // body []byte
	if err != nil {
		t.Error(err)
	}

	var cuvantCurent Cuvant
	jsonErr := json.Unmarshal(body, &cuvantCurent)
	if jsonErr != nil{
		t.Error(jsonErr)
	}

	log.Println(cuvantCurent.Definitions)

}