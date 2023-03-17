package reflections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	StreetName   string
	StreetNumber int
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 27},
			[]string{"Chris"},
		},
		{
			"struct non flat",
			Person{
				"Joao",
				Address{"Rua Dom Nunes", 13},
			},
			[]string{"Joao", "Rua Dom Nunes"},
		},
		{
			"struct with pointers",
			&Person{
				"Joao",
				Address{"Rua Dom Nunes", 13},
			},
			[]string{"Joao", "Rua Dom Nunes"},
		},
		{
			"slices",
			[]Address{
				{"Street X", 1},
				{"Street Y", 2},
			},
			[]string{"Street X", "Street Y"},
		},
		{
			"arrays",
			[2]Address{
				{"Street X", 1},
				{"Street Y", 2},
			},
			[]string{"Street X", "Street Y"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
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
		Walk(aMap, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Address)

		go func() {
			aChannel <- Address{"Rua X", 1}
			aChannel <- Address{"Rua Y", 2}
			close(aChannel)
		}()

		var got []string
		want := []string{"Rua X", "Rua Y"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, got []string, want string) {
	t.Helper()
	contains := false

	for _, x := range got {
		if x == want {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", got, want)
	}

}
