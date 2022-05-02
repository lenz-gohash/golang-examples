// source https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

package main

import (
	// "github.com/lenz-gohash/golang-examples/sample-tests/arithmetic"
	// "github.com/lenz-gohash/golang-examples/sample-tests/arithmetic"
	"fmt"

	"github.com/lenz-gohash/golang-examples/sample-tests/people"
)

func main() {
	// splitResult := manipulate.Split("a,b,c", ",")
	// fmt.Println("manipulate.Split(\"a,b,c\", \",\") =", splitResult)

	// addResult := arithmetic.Add(7, -8)
	// fmt.Println("arithmetic.Add(7, -8) =", addResult)

	person := people.Person{
		Name:        "John",
		Age:         35,
		WeightInLbs: 250,
		HeightInCm:  200,
	}

	giant := people.Giant{
		NameInAncientGreek: "Γιάννης",
		WeightInTons:       1.5,
		HeightInKms:        12.5,
		AgeInCenturies:     11,
	}

	fmt.Println("person.BMI =", people.Bmi(&person))
	fmt.Println("giant.BMI =", people.Bmi(&giant))

}
