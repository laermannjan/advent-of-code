use std::{
    cmp::Ordering,
    collections::{BinaryHeap, HashMap},
};

use crate::grid::{Coord, Direction, Grid};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct Candidate {
    coord: Coord,
    cost: usize,
}

impl Ord for Candidate {
    fn cmp(&self, other: &Self) -> Ordering {
        other
            .cost
            .cmp(&self.cost)
            .then_with(|| self.coord.y.cmp(&other.coord.y))
            .then_with(|| self.coord.x.cmp(&other.coord.x))
    }
}

impl PartialOrd for Candidate {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

static DIRECTIONS: [Direction; 4] = [
    Direction::North,
    Direction::South,
    Direction::East,
    Direction::West,
];

pub fn shortest_path<T>(
    grid: &Grid<T>,
    starts: &[Coord],
    end: &Coord,
    filter_neighbor: impl Fn(&Coord, &Coord) -> bool,
    compute_cost: fn(&Coord) -> usize,
) -> Option<usize>
where
    T: Default + Clone,
{
    let mut candidates = BinaryHeap::new();
    let mut costs = grid
        .coords()
        .map(|c| (c, usize::MAX))
        .collect::<HashMap<Coord, usize>>();

    for start in starts {
        candidates.push(Candidate {
            coord: *start,
            cost: 0,
        });
        costs.insert(*start, 0);
    }

    while let Some(Candidate { coord, cost }) = candidates.pop() {
        if coord == *end {
            return Some(cost);
        }

        // already found a shorter path
        if cost > *costs.get(&coord).unwrap() {
            continue;
        }

        let neighbors: Vec<Coord> = DIRECTIONS
            .iter()
            .filter_map(|d| {
                if let Some(new_coord) = grid.step(&coord, d) {
                    if filter_neighbor(&coord, &new_coord) {
                        return Some(new_coord);
                    }
                }
                None
            })
            .collect();

        for neighbor in neighbors {
            let next = Candidate {
                coord: neighbor,
                cost: cost + compute_cost(&neighbor),
            };

            if next.cost < *costs.get(&next.coord).unwrap() {
                candidates.push(next);
                costs.insert(next.coord, next.cost);
            }
        }
    }

    None
}
