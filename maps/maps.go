package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExist    = DictionaryErr("could not add the word because it already exist in the dictionary")
	ErrWordNotExist = DictionaryErr("could not update the word because the word doesn't exist in the dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
