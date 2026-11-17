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
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chipotle"},
			[]string{"Chipotle"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chipotle", "Chipotleland"},
			[]string{"Chipotle", "Chipotleland"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chipotle", 24},
			[]string{"Chipotle"},
		},
		{
			"nested fields",
			Person{
				"Chipotle",
				Profile{24, "Chipotleland"},
			},
			[]string{"Chipotle", "Chipotleland"},
		},
		{
			"pointers to things",
			&Person{
				"Chipotle",
				Profile{24, "Chipotleland"},
			},
			[]string{"Chipotle", "Chipotleland"},
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
