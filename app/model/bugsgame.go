package model

import (
	"math/rand"
	"time"

	"github.com/hculpan/go-sdl-lib/component"
)

type BugsGame struct {
	component.BaseGame

	BoardWidth      int
	BoardHeight     int
	Cycle           int
	Running         bool
	StartingBugs    int
	InitialSeedRate float64
	ReseedRate      float64
	BacteriaCount   int

	seedCarryover float64
	bugsCreated   int

	Board [][]bool
	Bugs  map[string]*Bug
}

var Bugs *BugsGame

func NewBugsGame(worldWidth, worldHeight int, startingBugs int, initialSeedRate float64, reseedRate float64) *BugsGame {
	rand.Seed(time.Now().UnixNano())

	result := BugsGame{
		BoardWidth:      worldWidth,
		BoardHeight:     worldHeight,
		StartingBugs:    startingBugs,
		InitialSeedRate: initialSeedRate,
		ReseedRate:      reseedRate,
	}

	Bugs = &result

	result.Board = make([][]bool, worldWidth)
	result.Bugs = map[string]*Bug{}

	result.Reset()

	return &result
}

func (s *BugsGame) Reset() {
	s.Cycle = 0
	s.Running = false
	s.seedCarryover = 0
	s.BacteriaCount = 0

	s.Bugs = map[string]*Bug{}

	for x := 0; x < s.BoardWidth; x++ {
		s.Board[x] = make([]bool, s.BoardHeight)
		for y := 0; y < s.BoardHeight; y++ {
			s.Board[x][y] = false
		}
	}

	s.PlaceInitialBacteria()
	s.PlaceInitialBugs()
}

func (s *BugsGame) Start() {
	s.Running = true
}

func (s *BugsGame) Stop() {
	s.Running = false
}

func (s *BugsGame) Update() error {
	if !s.Running {
		return nil
	}

	s.Cycle++

	for _, bug := range s.Bugs {
		s.MoveBug(bug)
		bug.Metabolize()
		foodCount := s.FeedAt(bug.X, bug.Y)
		s.BacteriaCount -= foodCount
		bug.ConsumeFood(foodCount * 40)

		if bug.Health <= 0 {
			delete(s.Bugs, bug.ID)
		} else if bug.CanReproduce() {
			s.bugsCreated++
			b1 := NewChildBug(s.bugsCreated, bug)
			s.Bugs[b1.ID] = b1
			s.bugsCreated++
			b2 := NewChildBug(s.bugsCreated, bug)
			s.Bugs[b2.ID] = b2
			delete(s.Bugs, bug.ID)
		}

	}

	s.PlaceNewBacteria()

	s.Running = len(s.Bugs) > 0

	return nil
}

func (s *BugsGame) FeedAt(x, y int) int {
	result := 0
	if x > 0 && y > 0 && s.Board[x-1][y-1] {
		result += 1
		s.Board[x-1][y-1] = false
	}
	if y > 0 && s.Board[x][y-1] {
		result += 1
		s.Board[x][y-1] = false
	}
	if x < s.BoardWidth-1 && y > 0 && s.Board[x+1][y-1] {
		result += 1
		s.Board[x+1][y-1] = false
	}
	if x > 0 && s.Board[x-1][y] {
		result += 1
		s.Board[x-1][y] = false
	}
	if s.Board[x][y] {
		result += 1
		s.Board[x][y] = false
	}
	if x < s.BoardWidth-1 && s.Board[x+1][y] {
		result += 1
		s.Board[x+1][y] = false
	}
	if x > 0 && y < s.BoardHeight-1 && s.Board[x-1][y+1] {
		result += 1
		s.Board[x-1][y+1] = false
	}
	if y < s.BoardHeight-1 && s.Board[x][y+1] {
		result += 1
		s.Board[x][y+1] = false
	}
	if x < s.BoardWidth-1 && y < s.BoardHeight-1 && s.Board[x+1][y+1] {
		result += 1
		s.Board[x+1][y+1] = false
	}
	return result
}

func (s *BugsGame) MoveBug(bug *Bug) {
	dir := bug.PickADirection()
	switch dir {
	case 0:
		bug.Y -= 2
	case 1:
		bug.X += 2
		bug.Y -= 1
	case 2:
		bug.X += 2
		bug.Y += 1
	case 3:
		bug.Y += 2
	case 4:
		bug.X -= 2
		bug.Y += 1
	case 5:
		bug.X -= 2
		bug.Y -= 1
	}

	if bug.X < 0 {
		bug.X += s.BoardWidth
	} else if bug.X >= s.BoardWidth {
		bug.X %= s.BoardWidth
	}

	if bug.Y < 0 {
		bug.Y += s.BoardHeight
	} else if bug.Y >= s.BoardHeight {
		bug.Y %= s.BoardHeight
	}
}

func (s *BugsGame) PlaceNewBacteria() {
	s.seedCarryover += s.ReseedRate
	for s.seedCarryover >= 1 {
		for cnt := 0; cnt < 10; cnt++ {
			x, y := rand.Intn(s.BoardWidth), rand.Intn(s.BoardHeight)
			if !s.Board[x][y] {
				s.Board[x][y] = true
				s.BacteriaCount++
				break
			}
		}
		s.seedCarryover -= 1
	}
}

func (s *BugsGame) PlaceInitialBacteria() {
	for x := 0; x < s.BoardWidth; x++ {
		for y := 0; y < s.BoardHeight; y++ {
			if !s.Board[x][y] {
				s.Board[x][y] = rand.Float64() < s.InitialSeedRate
				if s.Board[x][y] {
					s.BacteriaCount++
				}
			}
		}
	}
}

func (s *BugsGame) PlaceInitialBugs() {
	for n := 0; n < s.StartingBugs; n++ {
		x, y := rand.Intn(s.BoardWidth), rand.Intn(s.BoardHeight)
		bug := NewBug(n, x, y)
		s.Bugs[bug.ID] = bug
	}

	s.bugsCreated = s.StartingBugs
}
