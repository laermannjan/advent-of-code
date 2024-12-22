package main

import (
	"errors"
	"fmt"
	"lj/utils"
	"os"
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

	fmt.Println("checking", item)

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

	fmt.Println("next items", next_items)
	return next_items
}

func main() {
	input := utils.NewStdinInput()
	queue := []Item{{p: Position{row: 0, col: 0}, moving: east}}
	processed := map[Item]bool{}
	fmt.Println(queue)
	energized := map[Position]bool{
		{row: 0, col: 0}: true,
	}
	n_energized := 1
	var grid Grid = input.RunesSlice()
	fmt.Println(queue)

	for len(queue) > 0 {
		fmt.Println("queue", queue)
		item := queue[0]
		queue = queue[1:]

		if _, ok := energized[item.p]; !ok {
			energized[item.p] = true
			n_energized++
		}

		for _, next_item := range grid.next(item) {
			if _, ok := processed[next_item]; !ok {
				fmt.Println("adding", next_item)
				processed[next_item] = true
				queue = append(queue, next_item)
			}
		}

	}
	fmt.Fprintln(os.Stderr, n_energized)
}
