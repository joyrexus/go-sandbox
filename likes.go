package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
	Sex string
	Likes []string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is a %v year old %v", p.Name, p.Age, p.Sex)
}

const M string = "male"
const F string = "female"

func main() {

	people := map[string]*Person{
		"wv": {
			"Wilder Voigt",
			2,
			M,
			[]string{"monsters", "biking", "motorcycles", "cheese"},
		},
		"jv": {
			"Jason Voigt",
			46,
			M,
			[]string{"walking", "biking", "cartography"},
		},
		"mfm": {
			"Maggie Fritz-Morkin",
			32,
			F,
			[]string{"medieval comics", "urine", "cheese"},
		},
	}

	likes := map[string][]string{}
	for _, person := range people {
		for _, l := range person.Likes {
			likes[l] = append(likes[l], person.Name)
		}
	}
	for like, people := range likes {
		n := len(people)
		p, l := "people", "like"
		if n == 1 {
			p, l = "person", "likes"
		}
		fmt.Printf("%d %v %v %v: %v\n", n, p, l, like, people)
	}
	fmt.Println(len(likes["bacon"]), "people like bacon!")

}
