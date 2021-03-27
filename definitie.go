package definitii

import "regexp"

/*  Reprezinta o singura definitie.
	Text - propriu al definitie formatat in
	reprezentarea interna a site-ului.
*/
type Definitie struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Text string `json:"internalRep"`
}

/*	Ofera doar textul definitiei de pe site. Elimina contextul adaugat.
	Filtrele:
	* Filtrele din formatul intern al site-ului (@, #, $, -).
	* Textul ajutator din paranteze patrate.
	* Blankspace
	* Mai mult de doua spatii inlantuite.
*/
func (def *Definitie) FormatContext() string {
	var expresie = regexp.MustCompile(`(@[^@]*@)|(#[^#]*#)|(\$[^$]*\$)`)  // Filtru formating
	rez := expresie.ReplaceAllString(def.Text, "")

	var expresieLimbaj = regexp.MustCompile(`( -.*)`)  // Filtru pentru segmentul de provenienta al cuvantului
	rez = expresieLimbaj.ReplaceAllString(rez, "")

	// filtru paranteze goale
	var expresieParanteze = regexp.MustCompile(`\(\)`)
	rez = expresieParanteze.ReplaceAllString(rez, "")

	// Elimina tot ce este intre paranteze patrate.
	var expresieParantezePatrate = regexp.MustCompile(`\[.*]\.+`)
	rez = expresieParantezePatrate.ReplaceAllString(rez, "")

	// Filtru pentru blank space pe linie sau la inceputul stringului.
	var expresieRemoveBlankLines = regexp.MustCompile(`(^\n)|(^\s+$)|(^\s+)`)
	rez = expresieRemoveBlankLines.ReplaceAllString(rez, "")

	// Spatii inlantuite, cel putin doua.
	var expresieSpatii = regexp.MustCompile(`\s{2,}`)
	rez = expresieSpatii.ReplaceAllString(rez, " ")

	var expresieStea = regexp.MustCompile(`\*\*`)
	rez = expresieStea.ReplaceAllString(rez, "")

	return rez
}

/*	Ofera textul definitiei de pe site, fara tagurile de formatare (@, #, $, &).
	Filtrele sunt bazate pe expresii regulare pentru:
	* Filtrele din formatul intern al site-ului (@, #, $, &).
	* Textul ajutator din paranteze patrate.
	* Blankspace
	* Mai mult de doua spatii inlantuite.
*/
func (def *Definitie) Format() string {
	var expresie = regexp.MustCompile(`[@#$&]`)  // Filtru formating
	rez := expresie.ReplaceAllString(def.Text, "")

	// filtru paranteze goale
	var expresieParanteze = regexp.MustCompile(`\(\)`)
	rez = expresieParanteze.ReplaceAllString(rez, "")

	// Filtru pentru blank space pe linie sau la inceputul stringului.
	var expresieRemoveBlankLines = regexp.MustCompile(`(^\n)|(^\s+$)|(^\s+)`)
	rez = expresieRemoveBlankLines.ReplaceAllString(rez, "")

	// Spatii inlantuite, cel putin doua.
	var expresieSpatii = regexp.MustCompile(`\s{2,}`)
	rez = expresieSpatii.ReplaceAllString(rez, " ")

	return rez
}
