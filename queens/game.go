package queens

import "boardgames/board"

type Game struct {
	board *board.Board
}

// NewGame erstellt ein neues Spiel mit einem leeren 8x8-Spielfeld.
func NewGame() *Game {
	// Hinweis:
	// Erstellen Sie ein neues Board mit 8 Zeilen und 8 Spalten.
	// TODO
	return &Game{}
}

// Get liefert den Inhalt der Zelle an der angegebenen Position zurück.
// Liefert einen leeren String, falls die Position außerhalb des Spielfelds liegt.
func (g *Game) Get(row, col int) string {
	// Hinweis:
	// Verwenden Sie die Get-Methode des Boards, um den Inhalt der Zelle zu erhalten.
	// TODO
	return ""
}

// Set setzt eine Dame in der Zelle an der angegebenen Position.
func (g *Game) Set(row, col int) {
	// Hinweis:
	// Entscheiden Sie sich für ein Symbol, um eine Dame darzustellen.
	// Setzen Sie dieses Symbol mit der Set-Methode des Boards.
	// TODO
}

// Unset entfernt eine Dame aus der Zelle an der angegebenen Position.
func (g *Game) Unset(row, col int) {
	// Hinweis:
	// Entfernen Sie die Dame, indem Sie die Zelle mit einem leeren String überschreiben.
	// TODO
}

// String gibt eine textuelle Darstellung des Spiels zurück.
func (g Game) String() string {
	// Hinweis:
	// Verwenden Sie die String-Methode des Boards.
	// TODO
	return ""
}
