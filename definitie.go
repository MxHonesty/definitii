package definitii

/*  Reprezinta o singura definitie.
	Text - propriu al definitie formatat in
	reprezentarea interna a site-ului.
*/
type Definitie struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Text string `json:"internalRep"`
}

/*	TODO: Versiunea formatata a textului din reprezentarea interna a
	Elimina tot ce este intre $, #, @ din text.
	site-ului. De utilizat pentru display direct al continutului Text.
*/
func (def *Definitie) format() string {
	return def.Text
	
}
