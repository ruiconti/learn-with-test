package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age     int
	Brother string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name    string
				Brother string
			}{"Chris", "John"},
			[]string{"Chris", "John"},
		},
		{
			"Struct with multiple type fields",
			struct {
				Name    string
				Age     int
				Married bool
			}{"Chris", 33, false},
			[]string{"Chris"},
		},
		{
			"Struct with nested type",
			Person{
				"Chris",
				Profile{33, "John"},
			},
			[]string{"Chris", "John"},
		},
		{
			"Pointers passed",
			&Person{
				"Chris",
				Profile{33, "John"},
			},
			[]string{"Chris", "John"},
		},
		{
			"Slices passed",
			[]Profile{
				{33, "Johnny"},
				{35, "Mark"},
			},
			[]string{"Johnny", "Mark"},
		},
		{
			"Arrays passed",
			[2]Profile{
				{33, "Johnny"},
				{35, "Mark"},
			},
			[]string{"Johnny", "Mark"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("wanted %v got %v", test.ExpectedCalls, got)
			}
		})
	}

	// Maps need to be in a different test because Go doesn't guarantee
	// (key, value) pair order.
	// Hence, we need to check if values contain in it.
	t.Run("Maps passed", func(t *testing.T) {
		aMap := map[string]string{
			"A": "ACME Inc",
			"B": "Apple Inc",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "ACME Inc")
		assertContains(t, got, "Apple Inc")
	})

	t.Run("Channels passed", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{35, "Cato"}
			ch <- Profile{10, "Billy"}
			close(ch)
		}()

		var got []string
		want := []string{"Cato", "Billy"}

		walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})

	t.Run("Functions passed", func(t *testing.T) {
		fn := func() (Profile, Profile) {
			return Profile{35, "Mario"}, Profile{19, "Han"}
		}

		var got []string
		want := []string{"Mario", "Han"}

		walk(fn, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it did not.", haystack, needle)
	}
}
