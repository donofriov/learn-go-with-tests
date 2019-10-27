package dictionary

import "errors"

// Dictionary type
type Dictionary map[string]string

// ErrNotFound global error message
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search method for Dictionary type
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add method for Dictionary type
func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}

// Search function
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
