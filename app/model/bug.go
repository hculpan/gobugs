package model

import (
	"fmt"
	"math/rand"
)

type Bug struct {
	ID        string
	X         int
	Y         int
	Health    int
	Age       int
	Direction int
	Class     int

	Genes      [6]int
	GeneWeight [6]int
	GeneTotal  int
}

func NewBug(bugNumber, x, y int) *Bug {
	result := Bug{}

	result.X, result.Y = x, y

	result.Health = 400
	result.Age = 0
	result.Direction = rand.Intn(6)
	result.ID = fmt.Sprintf("bug-%05d", bugNumber)
	result.RandomGenes()

	fmt.Printf("NewBug: %s\n", result.ToString())
	return &result
}

func NewChildBug(bugNumber int, bug *Bug) *Bug {
	result := Bug{}

	result.X, result.Y = bug.X, bug.Y

	result.Health = bug.Health / 2
	result.Direction = rand.Intn(6)
	result.ID = fmt.Sprintf("bug-%05d", bugNumber)
	result.Age = 0

	for i := 0; i < 6; i++ {
		result.Genes[i] = bug.Genes[i]
	}

	result.RandomGenes()

	fmt.Printf("ChildBug: %s\n", result.ToString())
	return &result
}

func (b Bug) ToString() string {
	return fmt.Sprintf("ID: %s, Pos: [%d, %d], Health: %d, Age: %d", b.ID, b.X, b.Y, b.Health, b.Age)
}

func (b *Bug) Metabolize() {
	b.Health--
	b.Age++
}

func (b *Bug) CanReproduce() bool {
	return b.Age >= 800 && b.Health >= 1000
}

func (b *Bug) ConsumeFood(amount int) {
	b.Health += amount
	b.Health %= 1500
}

func (b *Bug) TotalGenes() {
	for i := 0; i < 6; i++ {
		b.GeneWeight[i] = b.Genes[i] * b.Genes[i] // Get rid of negative values
		b.GeneTotal += b.GeneWeight[i]
	}

	forwardMove := float64(b.GeneWeight[0]+b.GeneWeight[1]+b.GeneWeight[5]) / float64(b.GeneTotal)
	if forwardMove >= 0.75 {
		b.Class = 3
	} else if forwardMove >= .50 {
		b.Class = 2
	} else if forwardMove >= .25 {
		b.Class = 1
	} else {
		b.Class = 0
	}
}

func (b *Bug) RandomGenes() {
	b.GeneTotal = 0
	for i := 0; i < 6; i++ {
		b.Genes[i] = rand.Intn(11) - 5
	}

	b.TotalGenes()
}

func (b *Bug) MutateGenes(number int) {
	geneIndex := rand.Intn(6)
	if rand.Float64() < 0.50 {
		b.Genes[geneIndex] += 1
	} else {
		b.Genes[geneIndex] -= 1
	}

	b.TotalGenes()
}

func (b *Bug) PickADirection() int {
	n := rand.Intn(b.GeneTotal)
	for i := 0; i < 6; i++ {
		n -= b.GeneWeight[i]
		if n <= 0 {
			b.Direction = (b.Direction + i) % 6
			return b.Direction
		}
	}

	return 5
}
