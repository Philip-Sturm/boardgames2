package queens

// MoveAllowed prüft, ob ein Zug an der angegebenen Position erlaubt ist.
// Ein Zug ist erlaubt, wenn die Position innerhalb des Spielfelds liegt und die Zelle an dieser Position leer ist.
func (g *Game) MoveAllowed(row, col int) bool {
	// Hinweis:
	// Benutzen Sie die Methoden von `Board` um zu überprüfen, ob die Zeilen, Spalten und Diagonalen frei sind.
	// Bestimmen Sie vorher (für die Diagonalen) zum gegebenen (row/col)-Paar, in welcher Spalte die entsprechenden
	// Diagonalen beginnen.

	r, cl, cr := row, col, col
	for ; r > 0; r, cl, cr = r-1, cl-1, cr+1 {
	}

	return !g.board.RowContains(row, "*") &&
		!g.board.ColContains(col, "*") &&
		!g.board.DiagDownRightContains(cl, "*") &&
		!g.board.DiagDownLeftContains(cr, "*")
}

// Solve löst das Spiel.
func (g *Game) Solve(row int) bool {
	// Hinweis:
	// Gehen Sie rekursiv vor.
	// Als Abbruchbedingung können Sie verwenden, dass row den Wert 8 verwendet hat.
	// Gehen Sie dann in der aktuellen Zeile alle Felder durch, setzen Sie eine Dame (falls erlaubt),
	// und rufen Sie die Lösungsfunktion mit der nächsten Zeile auf.
	if row == 8 {
		return true
	}

	for col := 0; col < 8; col++ {
		if g.MoveAllowed(row, col) {
			g.Set(row, col)
			if g.Solve(row + 1) {
				return true
			}
			g.Unset(row, col)
		}
	}
	return false
}
