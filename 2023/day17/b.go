age main

import (
	"container/heap"
	"fmt"
	"lj/utils"
	"os"
)

type Position struct {
	row int
	col int
}

func (p *Position) Equal(other Position) bool {
	return p.row == other.row && p.col == other.col
}

func (p *Position) Move(off Offset) Position {
	return Position{
		row: p.row + off.row,
		col: p.col + off.col,
	}
}

type Offset struct {
	row int
	col int
}

func (p *Offset) Equal(other Offset) bool {
	return p.row == other.row && p.col == other.col
}

func (p *Offset) Opposite(other Offset) bool {
	return p.row == -other.row && p.col == -other.col
}

var (
	north Offset = Offset{-1, 0}
	south Offset = Offset{1, 0}
	west  Offset = Offset{0, -1}
	east  Offset = Offset{0, 1}
	null  Offset = Offset{0, 0}
)

type QItem struct {
	heatLoss  int
	pos       Position
	moved     Offset
	nStraight int
}

type SeenItem struct {
	pos           Position
	moved         Offset
	stepsStraight int
}

type PriorityQueue []QItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(QItem))
}

func (pq *PriorityQueue) Pop() any {
	n := pq.Len()
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

type Grid [][]rune

func (grid *Grid) contains(p Position) bool {
	g := *grid
	return 0 <= p.row && p.row < len(g) && 0 <= p.col && p.col < len(g[0])
}

func (grid *Grid) getHeatLoss(p Position) int {
	g := *grid
	return utils.Atoi(string(g[p.row][p.col]))
}

func main() {
	input := utils.NewStdinInput()
	var grid Grid = input.RunesSlice()
	start := Position{0, 0}
	end := Position{len(grid) - 1, len(grid[0]) - 1}

	q := PriorityQueue{}
	heap.Init(&q)
	heap.Push(&q, QItem{heatLoss: 0, pos: start, moved: null, nStraight: 0})

	seen := map[SeenItem]bool{}

	i := 0
	for len(q) > 0 {
		fmt.Printf("%v, %v", q[0], len(q))
		item := heap.Pop(&q).(QItem)

		if item.pos.Equal(end) && item.nStraight >= 4 {
			print(i)
			fmt.Fprintln(os.Stderr, item.heatLoss)
		}

		seenItem := SeenItem{
			pos:           item.pos,
			moved:         item.moved,
			stepsStraight: item.nStraight,
		}

		if _, ok := seen[seenItem]; ok {
			continue
		}

		seen[seenItem] = true

		for _, offset := range []Offset{north, south, west, east} {
			next_pos := item.pos.Move(offset)
			nStraight := item.nStraight
			if !grid.contains(next_pos) {
				continue
			} else if offset.Opposite(item.moved) {
				continue
			} else if offset.Equal(item.moved) {
				if nStraight >= 10 {
					continue
				}
				nStraight += 1
			} else {
				if nStraight < 4 && !item.moved.Equal(null) {
					continue
				}
				nStraight = 1
			}

			heap.Push(&q, QItem{
				pos:       next_pos,
				moved:     offset,
				heatLoss:  item.heatLoss + grid.getHeatLoss(next_pos),
				nStraight: nStraight,
			})

		}

	}
	fmt.Fprintln(os.Stderr, "<error>")
}
