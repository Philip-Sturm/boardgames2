package queens

// MoveAllowed prüft, ob ein Zug an der angegebenen Position erlaubt ist.
// Ein Zug ist erlaubt, wenn die Position innerhalb des Spielfelds liegt und die Zelle an dieser Position leer ist.
func (g *Game) MoveAllowed(row, col int) bool {
	if g.board.Get(row, col) == "" {
		return true
	}
	return false
}

// Solve löst das Spiel.
func (g *Game) Solve(row int) bool {
	// TODO
	return false
}
