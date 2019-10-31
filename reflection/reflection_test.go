package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Vince", "Vienna"},
			[]string{"Vince", "Vienna"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Vince", 30},
			[]string{"Vince"},
		},
		{
			"Nested fields",
			Person{
				"Vince",
				Profile{30, "Vienna"},
			},
			[]string{"Vince", "Vienna"},
		},
		{
			"Pointers to things",
			&Person{
				"Vince",
				Profile{30, "Vienna"},
			},
			[]string{"Vince", "Vienna"},
		},
		{
			"Slices",
			[]Profile{
				{28, "New York City"},
				{30, "Vienna"},
			},
			[]string{"New York City", "Vienna"},
		},
		{
			"Arrays",
			[2]Profile{
				{28, "New York City"},
				{30, "Vienna"},
			},
			[]string{"New York City", "Vienna"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
