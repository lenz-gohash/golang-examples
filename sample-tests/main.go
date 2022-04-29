// source https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

package main

import (
	"fmt"

	"github.com/lenz-gohash/golang-examples/sample-tests/arithmetic"
	"github.com/lenz-gohash/golang-examples/sample-tests/manipulate"
)

func main() {
	splitResult := manipulate.Split("a,b,c", ",")
	fmt.Println("manipulate.Split(\"a,b,c\", \",\") =", splitResult)

	addResult := arithmetic.Add(7, -8)
	fmt.Println("arithmetic.Add(7, -8) =", addResult)

}
