package snl

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// --- Data Structures ---

// Square represents a single cell on the board.
type Square struct {
	destination *Square // Where the square leads (snake, ladder, or normal).
}

// Board represents the game board.
type Board struct {
	squares       []Square
	snakeHeads    map[int]*Square // Quick lookup for snake heads.
	ladderBottoms map[int]*Square // Quick lookup for ladder bottoms.
}

// NewBoard creates a new 100-square game board.
func NewBoard() *Board {
	board := &Board{
		squares:       make([]Square, 100),
		snakeHeads:    make(map[int]*Square),
		ladderBottoms: make(map[int]*Square),
	}

	// Initialize a basic board without snakes and ladders.
	for i := range board.squares {
		board.squares[i].destination = &board.squares[i+1] // Normal movement
	}

	return board
}

// getDestination determines where a player lands after a move (handling snakes/ladders).
func (b *Board) getDestination(startSquare *Square) *Square {
	nextSquare := startSquare.destination

	// Check for snakes or ladders recursively (in case of chains).
	return b.getDestination(nextSquare)
}

// addSnake adds a snake to the board.
func (b *Board) addSnake(head, tail int) {
	headSquare := &b.squares[head-1]
	tailSquare := &b.squares[tail-1]
	headSquare.destination = tailSquare
	b.snakeHeads[head] = headSquare
}

// addLadder adds a ladder to the board.
func (b *Board) addLadder(start, end int) {
	startSquare := &b.squares[start-1]
	endSquare := &b.squares[end-1]
	startSquare.destination = endSquare
	b.ladderBottoms[start] = startSquare
}

// Player represents a player in the game.
type Player struct {
	name     string
	Position *Square // Current square on the board.
}

// rollDice simulates a dice roll.
func (p *Player) rollDice() int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
	return rand.Intn(6) + 1
}

// --- Game Logic ---

// Game represents the Snakes and Ladders game.
type Game struct {
	board         *Board
	Players       []Player
	CurrentPlayer int
}

// NewGame sets up a new game.
func NewGame() *Game {
	return &Game{
		board: NewBoard(),
	}
}

func (g *Game) Setup() {
	reader := bufio.NewReader(os.Stdin)

	// Get snake details
	fmt.Print("Enter number of snakes: ")
	var numSnakes int
	fmt.Scanln(&numSnakes)

	for i := 0; i < numSnakes; i++ {
		fmt.Printf("Enter snake head and tail (space separated): ")
		var head, tail int
		fmt.Scanln(&head, &tail)
		g.board.addSnake(head, tail)
	}

	// Get ladder details
	fmt.Print("Enter number of ladders: ")
	var numLadders int
	fmt.Scanln(&numLadders)

	for i := 0; i < numLadders; i++ {
		fmt.Printf("Enter ladder start and end (space separated): ")
		var start, end int
		fmt.Scanln(&start, &end)
		g.board.addLadder(start, end)
	}

	// Get player details
	fmt.Print("Enter number of players: ")
	var numPlayers int
	fmt.Scanln(&numPlayers)

	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter player %d name: ", i+1)
		var name string
		fmt.Fscanln(reader, &name)
		g.Players = append(g.Players, Player{name: name, Position: 0})
	}
}

func (g *Game) PlayTurn() {
	player := &g.Players[g.CurrentPlayer]

	roll := player.rollDice()
	oldPosition := player.Position
	newPosition := oldPosition + roll

	if newPosition > 100 {
		newPosition = oldPosition
	} else {
		newPosition = g.board.getDestination(newPosition)
	}

	player.Position = newPosition

	fmt.Printf("%s rolled a %d and moved from %d to %d\n", player.name, roll, oldPosition, newPosition)

	if newPosition == 100 {
		fmt.Printf("%s wins the game\n", player.name)
	}

	g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
}

func main() {

}
