package main

import (
	"fmt"
	"time"
)

func main() {
	f := GetProblemFactory()

	for dice := 1; dice <= 10; dice++ {
		p := f.Get(Difficult, 12, dice)
		g := NewGame(p)

		start := time.Now()
		solutions := g.Solve()
		runtime := time.Since(start)

		//for _, sol := range solutions {
		//	fmt.Println(sol.String())
		//}
		fmt.Printf("%s\tFound %d solutions in %s\n", p, len(solutions), runtime)
	}
}
