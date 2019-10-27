package dictionary

import "errors"

// Dictionary type
type Dictionary map[string]string

// Search method for Dictionary type
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}

// Search function
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
