package definitii

/** Patreaza date pentru un cuvant. */
type Cuvant struct {
	Type string `json:"type"`
	Word string `json:"word"`
	Definitions []Definitie `json:"definitions"` // Pastreaza lista de definitii
}
