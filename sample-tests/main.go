// source https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// func main() {
// 	result := manipulate.Split("a,b,c", ",")
// 	fmt.Println(result)
// }

func main() {
	type T struct {
		I int
	}
	x := []*T{{1}, {2}, {3}, {4}}
	y := []*T{{1}, {2}, {4}, {5}}
	diff := cmp.Diff(x, y)
	fmt.Printf(diff)
}
