package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		want := "this is just a test"

		assertDefinition(t, dictionary, "test", want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("untest")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertErrors(t, err, ErrKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"a": "Abril"}

	t.Run("New word", func(t *testing.T) {
		ok := dictionary.Add("b", "Brevereiro")
		want := "Brevereiro"

		assertNoError(t, ok)
		assertDefinition(t, dictionary, "b", want)
	})

	t.Run("Existing word", func(t *testing.T) {
		err := dictionary.Add("b", "Bavaria")
		want := ErrKeyAlreadyExists

		assertErrors(t, err, want)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"c": "Casa"}

	t.Run("Updates an existing word", func(t *testing.T) {
		ok := dictionary.Update("c", "Cabral")
		want := "Cabral"

		assertNoError(t, ok)
		assertDefinition(t, dictionary, "c", want)
	})

	t.Run("Updates a non-existing word", func(t *testing.T) {
		err := dictionary.Update("d", "Dracula")
		want := ErrKeyNotFound

		assertErrors(t, err, want)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"e": "Eldorado", "f": "Fagulha"}

	t.Run("Deletes an existing word", func(t *testing.T) {
		ok := dictionary.Delete("f")
		assertNoError(t, ok)

		_, err := dictionary.Search("f")
		assertErrors(t, err, ErrKeyNotFound)
	})

	t.Run("Deletes a non-existing word", func(t *testing.T) {
		err := dictionary.Delete("k")
		want := ErrKeyNotFound

		assertErrors(t, err, want)
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error")
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, want string) {
	t.Helper()

	got, _ := dictionary.Search(word)

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
