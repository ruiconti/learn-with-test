package main

import "testing"

/*
	Writing tests:
	1. needs to be in a file with a name like "xxx_test.go"
	2. test function must start with "Test"
	3. test function takes one argument only: t *testing.T
*/

/*
	Test cycle:
	1. Write a test
	2. Make the compiler pass
	3. Run the test, see that it fails and check the error message is meaningful
	4. Write enough code to make the test pass
	5. Refactor
*/

func TestHello(t *testing.T) {
	// t is a hook to testing framework

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		/*
			t.Helper() is needed to tell the test suite that this method is a helper.
			By doing this when it fails the line number reported will be in our function
			call rather than inside our test helper.
		*/
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		// these are subtests
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greetings in Portuguese", func(t *testing.T) {
		got := Hello("Elodie", portuguese)
		want := "Olá, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greetings in French", func(t *testing.T) {
		got := Hello("Rui", french)
		want := "Bonjour, Rui"
		assertCorrectMessage(t, got, want)
	})
}
