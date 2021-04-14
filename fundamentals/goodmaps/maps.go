package maps

const (
	ErrKeyNotFound      = DictionaryErr("word not in dictionary")
	ErrKeyAlreadyExists = DictionaryErr("word already exists in dictionary")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrKeyNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrKeyNotFound:
		d[word] = definition
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrKeyNotFound:
		return ErrKeyNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrKeyNotFound:
		return ErrKeyNotFound
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

// It made me wonder: How could we generalize this? Or in Gopher way this is
// not something to pursue?
