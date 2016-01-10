package main

import (
	"fmt"
	"sort"
)

type Vertex struct {
	X float64
	Y float64
}

type Person struct {
	Name  string
	Age   int
	Sex   string
	Likes []string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is a %v year old %v", p.Name, p.Age, p.Sex)
}

const M string = "male"
const F string = "female"

func main() {
	places := make(map[string]Vertex)
	places["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	places["home"] = Vertex{42.0, -76.0}
	places["work"] = Vertex{33.0, -82.0}
	fmt.Println(places)

	m := map[string]int{
		"alpha": 1,
		"beta":  2,
	}
	fmt.Println(m)

	people := map[string]*Person{
		"wv": {
			"Wilder Voigt",
			2,
			M,
			[]string{"monsters", "motorcycles", "cheese"},
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
			[]string{"medieval comics", "urine"},
		},
	}

	jv, ok := people["jv"]

	jv.Name = "J. Voigt"

	fmt.Println(jv)
	fmt.Println(people["mfm"])
	fmt.Println(len(people))

	var keys []string
	for p, person := range people {
		fmt.Printf("%q: %v\n", p, person)
		keys = append(keys, p)
	}
	sort.Strings(keys)
	fmt.Println("\nSORTED ORDER:")
	for _, k := range keys {
		fmt.Printf("%q: %v\n", k, people[k])
	}

	delete(people, "jv")

	jv, ok = people["jv"]
	if ok == false {
		fmt.Println("oh noes! jv is gone!")
	} else {
		fmt.Println(jv)
	}

	type Key struct {
		Path, Country string
	}

	hits := make(map[Key]int)
	hits[Key{"/doc/", "au"}]++
	hits[Key{"/doc/", "us"}] += 22
	hits[Key{"/doc/", "ca"}] += 11
	fmt.Println(hits)
}
