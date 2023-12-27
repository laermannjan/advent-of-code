package main

import (
	"aoc-go/utils"
	"container/heap"
	"log"
)

type Position struct {
	row int
	col int
}

type QItem struct {
	heatLoss      int
	row           int
	col           int
	drow          int
	dcol          int
	stepsStraight int
}

type SeenItem struct {
	row           int
	col           int
	drow          int
	dcol          int
	stepsStraight int
}

type PriorityQueue []QItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].heatLoss != pq[j].heatLoss {
		return pq[i].heatLoss < pq[j].heatLoss
	}
	if pq[i].row != pq[j].row {
		return pq[i].row < pq[j].row
	}
	if pq[i].col != pq[j].col {
		return pq[i].col < pq[j].col
	}
	if pq[i].drow != pq[j].drow {
		return pq[i].drow < pq[j].drow
	}
	if pq[i].dcol != pq[j].dcol {
		return pq[i].dcol < pq[j].dcol
	}
	return pq[i].stepsStraight < pq[j].stepsStraight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(QItem))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	// old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func part1(input utils.Input) (answer interface{}) {
	grid := input.RunesSlice()
	q := PriorityQueue{}
	heap.Init(&q)
	heap.Push(&q, QItem{row: 0, col: 0, heatLoss: 0, drow: 0, dcol: 0, stepsStraight: 0})

	seen := map[SeenItem]bool{}

	i := 0
	for len(q) > 0 {
		log.Printf("%v, %v", q[0], len(q))
		item := heap.Pop(&q).(QItem)

		if item.row == len(grid)-1 && item.col == len(grid[0])-1 {
			print(i)
			return item.heatLoss
		}

		seenItem := SeenItem{
			row:           item.row,
			col:           item.col,
			drow:          item.drow,
			dcol:          item.dcol,
			stepsStraight: item.stepsStraight,
		}

		if _, ok := seen[seenItem]; ok {
			continue
		}

		seen[seenItem] = true

		if item.stepsStraight < 3 && !(item.drow == 0 && item.dcol == 0) {
			nrow := item.row + item.drow
			ncol := item.col + item.dcol

			if 0 <= nrow && nrow < len(grid) && 0 <= ncol && ncol < len(grid[0]) {
				// log.Println("adding straight", nrow, ncol)
				heap.Push(&q, QItem{
					row:           nrow,
					col:           ncol,
					drow:          item.drow,
					dcol:          item.dcol,
					heatLoss:      item.heatLoss + utils.Atoi(string(grid[nrow][ncol])),
					stepsStraight: item.stepsStraight + 1,
				})
			}

		}

		for _, offset := range []Position{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		} {
			// log.Println("checking offset", offset, "for item", item)
			ndrow := offset.row
			ndcol := offset.col
			if !(ndrow == item.drow && ndcol == item.dcol) && !(ndrow == -item.drow && ndcol == -item.dcol) {
				// log.Println("not straight and not opposite")

				nrow := item.row + ndrow
				ncol := item.col + ndcol
				if 0 <= nrow && nrow < len(grid) && 0 <= ncol && ncol < len(grid[0]) {
					// log.Println("adding", nrow, ncol)
					heap.Push(&q, QItem{
						row:           nrow,
						col:           ncol,
						drow:          offset.row,
						dcol:          offset.col,
						heatLoss:      item.heatLoss + utils.Atoi(string(grid[nrow][ncol])),
						stepsStraight: 1,
					})
				}
			}
		}
		i++

	}
	// for r, line := range input.RunesSlice() {
	// 	for c, heat_loss := range line {
	// 	}
	// }
	return
}

func part2(input utils.Input) (answer interface{}) {
	return
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
