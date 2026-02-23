package queens

import "boardgames/board"

type Game struct {
	board *board.Board
}

// NewGame erstellt ein neues Spiel mit einem leeren 8x8-Spielfeld.
func NewGame() *Game {
	return &Game{
		board: board.New(8, 8),
	}
}

// Get liefert den Inhalt der Zelle an der angegebenen Position zurück.
// Liefert einen leeren String, falls die Position außerhalb des Spielfelds liegt.
func (g *Game) Get(row, col int) string {
	return g.board.Get(row, col)
}

// Set setzt eine Dame in der Zelle an der angegebenen Position.
func (g *Game) Set(row, col int) {
	g.board.Set(row, col, "*")
}

// Unset entfernt eine Dame aus der Zelle an der angegebenen Position.
func (g *Game) Unset(row, col int) {
	g.board.Set(row, col, " ")
}

// String gibt eine textuelle Darstellung des Spiels zurück.
func (g Game) String() string {
	return g.board.String()
}
