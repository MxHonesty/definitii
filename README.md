# definitii
Un library pentru <b>definitii</b> de pe https://dexonline.ro/ scris in Go.

Utilizare
---
go.mod
```go
require github.com/MxHonesty/definitii v1.2.0  // sau orice alta versiune curenta din tags
```
.go file
```go
  import "github.com/MxHonesty/definitii"
```
#### Functia FetchCuvant
```go
  cuvant, err := definitii.FetchCuvant("cuvant")
```
* <b>cuvant</b> - Reprezentare a listei de definitii de pe site, obtinute prin https://dexonline.ro/definitie/cuvant/json.
* Tipul de erori posibile: <b>CuvantNotFoundError</b>, <b>nil</b>.


Reprezentarile Cuvintelor si Definitiilor
---

```go
type Cuvant struct {
	Type string `json:"type"`
	Word string `json:"word"`
	Definitions []Definitie `json:"definitions"` // Pastreaza lista de definitii
}
```

``` go
type Definitie struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Text string `json:"internalRep"`  // Textul definitiei
}
```
Text contine definitia cuvantului in <b> reprezentarea interna </b> dexonline. Pentru un format mai usor de citit
sunt utilizate urmatoarele metode ale structului Definitie.
#### Metodele format si formatContext
Aplicate asupra unei singure definitii, metodele ofera versiuni cu grad diferit de formatare asupra textului.
```go
	cuvant.Definitions[0].Text  // Reprezentarea interna a definitiei pentru dexonline.
	cuvant.Definitions[0].format()  // Elimina tagurile din reprezentarea interna.
	cuvant.Definitions[0].formatContext()  // Elimina tot contextul insotit de taguri.
```
```
@TEST^2,@ $teste,$ #s. n.# Înveliș calcaros sau chitinos care protejează corpul unor moluște, crustacee etc. - Din #fr.# @test.@
TEST^2, teste, s. n. Înveliș calcaros sau chitinos care protejează corpul unor moluște, crustacee etc. - Din fr. test.
Înveliș calcaros sau chitinos care protejează corpul unor moluște, crustacee etc.
```

Exemplu utilizare
---
go.mod
```go
module awesomeProject1

go 1.16

require github.com/MxHonesty/definitii v1.2.0
```

main.go
```go
package main

import (
	"fmt"
	"github.com/MxHonesty/definitii"
)

func main() {
	cuvant, err := definitii.FetchCuvant("masina")
	if err != nil {
		fmt.Println(err)  // Cuvantul nu a fost gasit.
	} else {
		fmt.Println(cuvant.Definitions[0].Format())  // Afiseaza PRIMA definitie
	}
}
```

Output:
```
MAȘÍNĂ, mașini, s. f. 1. Sistem tehnic alcătuit din piese cu mișcări determinate, care transformă o formă de energie în altă formă de energie sau în lucru mecanic util; ...
```
