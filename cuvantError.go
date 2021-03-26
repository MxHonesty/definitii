package definitii

/** Eroare returnata daca cuvantul dat nu a fost gasit */
type CuvantNotFoundError struct {}

func(m *CuvantNotFoundError) Error() string {
	return "Cuvantul nu a fost gasit"
}