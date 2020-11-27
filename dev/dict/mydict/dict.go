package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

// Money type
type Money int

var (
	errNokey  = errors.New("Not Found")
	errUpdate = errors.New("You can't update non-existing word")
	errDelete = errors.New("You can't delete non-existing word")
	errDup    = errors.New("Duplicated Keys")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, err := d[word]
	if err {
		return value, nil
	} else {
		return "", errNokey
	}
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNokey:
		d[word] = def
	case nil:
		return errDup
	}
	return nil
}

// Update a word to the dict
func (d Dictionary) Update(word, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case errNokey:
		return errUpdate
	case nil:
		d[word] = newDef
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNokey:
		return errUpdate
	case nil:
		delete(d, word)
	}
	return nil

}
