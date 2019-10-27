package dictionary

// Dictionary type
type Dictionary map[string]string

// ErrNotFound, ErrWordExists global error message variable
const (
	ErrNotFound         = Err("could not find the word you were looking for")
	ErrWordExists       = Err("cannot add word because it already exists")
	ErrWordDoesNotExist = Err("cannot update word because it does not exist")
)

// Err custom string type
type Err string

// Error method for Err type
func (e Err) Error() string {
	return string(e)
}

// Search method for Dictionary type
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Search function
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// Add method for Dictionary type
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update word definition method for Dictionary type
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
