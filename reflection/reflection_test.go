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
}
