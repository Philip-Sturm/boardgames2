package queens

import "fmt"

func ExampleGame_MoveAllowed() {
	g1 := NewGame()
	g2 := NewGame()
	g3 := NewGame()

	fmt.Println(g1.MoveAllowed(0, 0))
	fmt.Println(g2.MoveAllowed(0, 0))
	fmt.Println(g3.MoveAllowed(0, 0))
	fmt.Println()

	g1.Set(2, 2)
	g2.Set(0, 2)
	g3.Set(2, 0)

	fmt.Println(g1.MoveAllowed(0, 0))
	fmt.Println(g1.MoveAllowed(0, 1))
	fmt.Println(g1.MoveAllowed(1, 0))
	fmt.Println(g1.MoveAllowed(1, 1))
	fmt.Println()

	fmt.Println(g2.MoveAllowed(0, 0))
	fmt.Println(g2.MoveAllowed(0, 1))
	fmt.Println(g2.MoveAllowed(1, 0))
	fmt.Println(g2.MoveAllowed(1, 1))
	fmt.Println()

	fmt.Println(g3.MoveAllowed(0, 0))
	fmt.Println(g3.MoveAllowed(0, 1))
	fmt.Println(g3.MoveAllowed(1, 0))
	fmt.Println(g3.MoveAllowed(1, 1))
	fmt.Println()

	// Output:
}
