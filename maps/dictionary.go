package dictionary

// Dictionary type
type Dictionary map[string]string

// Search method for Dictionary type
func (d Dictionary) Search(word string) string {
	return d[word]
}

// Search function
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
