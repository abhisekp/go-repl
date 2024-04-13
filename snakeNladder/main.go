package snakeNladder

import "math/rand"
import "fmt"

var dice *Dice

func init() {
	dice = NewDice()
}

type Dice struct {
}

func NewDice() *Dice {
	return &Dice{}
}

func (this *Dice) Throw() int {
	return rand.Intn(6) + 1
}

// ------

type Player struct {
	Id  int
	Pos int
}

func (this *Player) Move(diceVal int) {
	newPos := this.Pos + diceVal
	if newPos < 100 {
		this.Pos = newPos
	}
}

// ---

type Cell struct {
	Entity Pos
}

func (this *Cell) UpdatePlayerPos(player *Player) bool {
	fmt.Println(this.Entity)
	if this.Entity.Head == this.Entity.Tail {
		return false
	}

	if this.Entity.Head > this.Entity.Tail {
		fmt.Println("Encountered a Snake")
	} else {
		fmt.Println("Encountered a Ladder")
	}
	fmt.Printf("Updated Player Pos from %d to %d\n", player.Pos+1, this.Entity.Tail+1)
	player.Pos = this.Entity.Tail

	return true
}

// ----

type Pos struct {
	Head, Tail int
}

// ---

type Board struct {
	Grid [100]Cell
}

type BoardOpts struct {
	Snakes, Ladders []Pos
}

func (this *Board) Init(options BoardOpts) {
	// Update Cells with Snakes
	for _, snake := range options.Snakes {
		this.Grid[snake.Head] = Cell{
			Entity: snake,
		}
	}

	// Update Cells with Ladder
	for _, ladder := range options.Ladders {
		this.Grid[ladder.Head] = Cell{
			Entity: ladder,
		}
	}
}

func (this *Board) MovePlayer(player *Player, diceVal int) {
	// Update player position based on dice value
	player.Move(diceVal)
	fmt.Println("Player position updated based on dice:", player.Pos+1)

	this.Grid[player.Pos].UpdatePlayerPos(player)
}

func NewBoard(options BoardOpts) *Board {
	board := Board{}

	if len(options.Snakes) == 0 || len(options.Ladders) == 0 {
		board.Init(options)
	} else {
		board.Init(BoardOpts{
			Snakes: []Pos{
				{92, 8},
				{78, 68},
				{56, 42},
				{31, 1},
				{69, 33},
			},
			Ladders: []Pos{
				{65, 95},
				{46, 79},
				{37, 98},
				{17, 27},
				{9, 67},
				{1, 67},
				{2, 67},
				{3, 67},
				{4, 67},
				{5, 67},
				{6, 67},
			},
		})
	}

	return &board
}

// ----

type Game struct {
	Players          []Player
	CurrentPlayerIdx int
	Board            *Board
}

func (this *Game) Run() {
	for {
		this.Play()

		if this.Players[this.CurrentPlayerIdx].Pos == 100 { // Winning
			fmt.Println("\n\nWinning Player: ", this.CurrentPlayerIdx)
			break
		}

		// Next Player
		this.CurrentPlayerIdx = (this.CurrentPlayerIdx + 1) % len(this.Players)
		fmt.Println("\nNext Player: ", this.Players[this.CurrentPlayerIdx].Id)
	}
}

func (this *Game) Play() {
	diceVal := dice.Throw()
	fmt.Printf("Player %d, Throws Dice: %d\n", this.Players[this.CurrentPlayerIdx].Id, diceVal+1)

	this.Board.MovePlayer(&this.Players[this.CurrentPlayerIdx], diceVal)
}

type GameOpts struct {
	TotalPlayers    int
	Snakes, Ladders []Pos
}

func NewGame(options ...GameOpts) *Game {
	option := GameOpts{
		TotalPlayers: 2,
	}
	if len(options) > 1 {
		option = options[0]
	}

	players := make([]Player, 0, option.TotalPlayers)
	for i := range option.TotalPlayers {
		players = append(players, Player{
			Id: i + 1,
		})
	}
	return &Game{
		Players:          players,
		CurrentPlayerIdx: 0,
		Board:            NewBoard(BoardOpts{Snakes: option.Snakes, Ladders: option.Ladders}),
	}
}
