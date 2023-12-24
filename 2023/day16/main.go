package main

import (
	"aoc-go/utils"
	"errors"
	"log"
)

type Direction int

const (
	west  Direction = iota
	north Direction = iota
	east  Direction = iota
	south Direction = iota
)

type Position struct {
	row int
	col int
}

type Item struct {
	p      Position
	moving Direction
}

type Grid [][]rune

func (g *Grid) move(pos Position, d Direction) (Position, error) {
	switch d {
	case north:
		if pos.row > 0 {
			pos.row -= 1
			return pos, nil
		}
	case south:
		if pos.row < len(*g)-1 {
			pos.row += 1
			return pos, nil
		}
	case east:
		if pos.col < len((*g)[0])-1 {
			pos.col += 1
			return pos, nil
		}
	case west:
		if pos.col > 0 {
			pos.col -= 1
			return pos, nil
		}
	}
	return Position{}, errors.New("out of grid")
}

func (g *Grid) next(item Item) []Item {
	next_items := []Item{}

	log.Println("checking", item)

	switch symbol := (*g)[item.p.row][item.p.col]; symbol {
	case '/':
		switch item.moving {
		case south:
			if next, err := g.move(item.p, west); err == nil {
				next_items = append(next_items, Item{p: next, moving: west})
			}
		case east:
			if next, err := g.move(item.p, north); err == nil {
				next_items = append(next_items, Item{p: next, moving: north})
			}
		case north:
			if next, err := g.move(item.p, east); err == nil {
				next_items = append(next_items, Item{p: next, moving: east})
			}
		case west:
			if next, err := g.move(item.p, south); err == nil {
				next_items = append(next_items, Item{p: next, moving: south})
			}
		}

	case '\\':
		switch item.moving {
		case south:
			if next, err := g.move(item.p, east); err == nil {
				next_items = append(next_items, Item{p: next, moving: east})
			}
		case east:
			if next, err := g.move(item.p, south); err == nil {
				next_items = append(next_items, Item{p: next, moving: south})
			}
		case north:
			if next, err := g.move(item.p, west); err == nil {
				next_items = append(next_items, Item{p: next, moving: west})
			}
		case west:
			if next, err := g.move(item.p, north); err == nil {
				next_items = append(next_items, Item{p: next, moving: north})
			}
		}

	case '-':
		switch item.moving {
		case north, south:
			if next, err := g.move(item.p, west); err == nil {
				next_items = append(next_items, Item{p: next, moving: west})
			}
			if next, err := g.move(item.p, east); err == nil {
				next_items = append(next_items, Item{p: next, moving: east})
			}
		case west, east:
			if next, err := g.move(item.p, item.moving); err == nil {
				next_items = append(next_items, Item{p: next, moving: item.moving})
			}
		}

	case '|':
		switch item.moving {
		case west, east:
			if next, err := g.move(item.p, north); err == nil {
				next_items = append(next_items, Item{p: next, moving: north})
			}
			if next, err := g.move(item.p, south); err == nil {
				next_items = append(next_items, Item{p: next, moving: south})
			}
		case north, south:
			if next, err := g.move(item.p, item.moving); err == nil {
				next_items = append(next_items, Item{p: next, moving: item.moving})
			}
		}
	case '.':
		if next, err := g.move(item.p, item.moving); err == nil {
			next_items = append(next_items, Item{p: next, moving: item.moving})
		}

	}

	log.Println("next items", next_items)
	return next_items
}

func part1(input utils.Input) (answer interface{}) {
	queue := []Item{{p: Position{row: 0, col: 0}, moving: east}}
	processed := map[Item]bool{}
	log.Println(queue)
	energized := map[Position]bool{
		{row: 0, col: 0}: true,
	}
	n_energized := 1
	var grid Grid = input.RunesSlice()
	log.Println(queue)

	for len(queue) > 0 {
		log.Println("queue", queue)
		item := queue[0]
		queue = queue[1:]

		if _, ok := energized[item.p]; !ok {
			energized[item.p] = true
			n_energized++
		}

		for _, next_item := range grid.next(item) {
			if _, ok := processed[next_item]; !ok {
				log.Println("adding", next_item)
				processed[next_item] = true
				queue = append(queue, next_item)
			}
		}

	}
	return n_energized
}

func part2(input utils.Input) (answer interface{}) {
	var grid Grid = input.RunesSlice()
	starts := []Item{}
	starts = append(starts, Item{Position{0, 0}, south})
	starts = append(starts, Item{Position{0, 0}, east})
	starts = append(starts, Item{Position{0, len(grid[0]) - 1}, south})
	starts = append(starts, Item{Position{0, len(grid[0]) - 1}, west})
	starts = append(starts, Item{Position{len(grid) - 1, 0}, south})
	starts = append(starts, Item{Position{len(grid) - 1, 0}, east})
	starts = append(starts, Item{Position{len(grid) - 1, len(grid[0]) - 1}, south})
	starts = append(starts, Item{Position{len(grid) - 1, len(grid[0]) - 1}, west})
	for row := 1; row < len(grid)-1; row++ {
		starts = append(starts, Item{Position{row: row, col: 0}, east})
		starts = append(starts, Item{Position{row: row, col: len(grid[0]) - 1}, west})
	}
	for col := 1; col < len(grid[0])-1; col++ {
		starts = append(starts, Item{Position{row: 0, col: col}, south})
		starts = append(starts, Item{Position{row: len(grid) - 1, col: col}, north})
	}

	max_energized := 0

	for _, start := range starts {
		queue := []Item{start}
		processed := map[Item]bool{}
		energized := map[Position]bool{start.p: true}
		n_energized := 1

		for len(queue) > 0 {
			log.Println("queue", queue)
			item := queue[0]
			queue = queue[1:]

			if _, ok := energized[item.p]; !ok {
				energized[item.p] = true
				n_energized++
			}

			for _, next_item := range grid.next(item) {
				if _, ok := processed[next_item]; !ok {
					log.Println("adding", next_item)
					processed[next_item] = true
					queue = append(queue, next_item)
				}
			}

		}

		max_energized = max(max_energized, n_energized)

	}
	return max_energized
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
