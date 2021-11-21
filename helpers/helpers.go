package helpers

// Prüft eine Liste aus Strings, ob alle ihre Einträge gleich 's' sind.
func AllElementsEqualTo(list []string, s string) bool {
	for _, element := range list {
		if element != s {
			return false
		}
	}
	return true
}
