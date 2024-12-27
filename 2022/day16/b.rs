use std::io::{self, Read};
use std::{
    cmp::Ordering,
    collections::{BTreeSet, BinaryHeap, HashMap, HashSet, VecDeque},
};

use itertools::Itertools;

type Input<'a> = HashMap<&'a str, Valve<'a>>;

#[derive(Debug, PartialEq, Eq)]
pub struct Valve<'a> {
    flow_rate: u32,
    connections: HashSet<&'a str>,
}

#[derive(PartialEq, Eq)]
struct Node<'a> {
    cost: u32,
    id: &'a str,
}

impl<'a> Ord for Node<'a> {
    fn cmp(&self, other: &Self) -> Ordering {
        other.cost.cmp(&self.cost)
    }
}

impl<'a> PartialOrd for Node<'a> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(Debug, Hash, PartialEq, Eq, Clone)]
struct State<'a> {
    opened: BTreeSet<&'a str>,
    curr: &'a str,
    elapsed: u32,
    released: u32,
}

/// return lowest cost to move from a valve to an other valve
fn min_cost(from: &str, to: &str, map: &HashMap<&str, Valve>) -> u32 {
    // shortest path:
    // Dijkstra's algorithm
    // nodes in the priority queue are sorted so the lowest cost gets popped first
    let mut pq = BinaryHeap::new();
    // prevent backtracking by keeping track of valve rooms we already saw
    let mut seen = HashSet::new();

    pq.push(Node { cost: 0, id: from });
    seen.insert(from);

    while let Some(Node { cost, id }) = pq.pop() {
        if id == to {
            return cost;
        }

        for neighbour in map[id].connections.iter() {
            // only insert into the pq if we did not already see the neighbour valve
            if seen.insert(neighbour) {
                pq.push(Node {
                    cost: cost + 1,
                    id: neighbour,
                });
            }
        }
    }

    u32::MAX
}

/// map shortest distance from "AA" to any flowing valve
/// map shortest distance from any flowing valve to an other
fn min_distances<'a>(map: &'a HashMap<&str, Valve>) -> HashMap<(&'a str, &'a str), u32> {
    map.iter()
        // only keep flowing valves
        .filter(|(_, valve)| valve.flow_rate > 0)
        // get the name of flowing valves
        .map(|(&name, _)| name)
        // iterate over every combination of 2 flowing valves
        .tuple_combinations()
        // record shortest distance between those 2
        // (and the dist from "AA" to a flowing valve because we start there)
        .fold(HashMap::new(), |mut acc, (name1, name2)| {
            // from AA to name1
            acc.entry(("AA", name1))
                .or_insert_with(|| min_cost("AA", name1, map));
            // from AA to name2
            acc.entry(("AA", name2))
                .or_insert_with(|| min_cost("AA", name2, map));

            let dist = min_cost(name1, name2, map);
            // from name1 to name2
            acc.insert((name1, name2), dist);
            // from name2 to name1
            acc.insert((name2, name1), dist);

            acc
        })
}

fn wait_until_ending(
    max_time: u32,
    elapsed: u32,
    released: u32,
    opened: &BTreeSet<&str>,
    map: &HashMap<&str, Valve>,
) -> u32 {
    let time_left = max_time - elapsed;
    let released_per_min: u32 = opened.iter().map(|name| &map[name].flow_rate).sum();
    released + (released_per_min * time_left)
}

pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    let map = input
        .lines()
        .map(|line| {
            let (valve, connections) = line.split_once("; ").unwrap();
            let value = valve.strip_prefix("Valve ").unwrap();
            let (name, flow_rate) = value.split_once(" has flow rate=").unwrap();
            let flow_rate = flow_rate.parse().unwrap();
            let connections = connections
                .strip_prefix("tunnels lead to valves ")
                .or_else(|| connections.strip_prefix("tunnel leads to valve "))
                .unwrap();
            let connections = connections.split(", ").collect();

            (
                name,
                Valve {
                    flow_rate,
                    connections,
                },
            )
        })
        .collect();

    let dist_map = min_distances(&map); // key: (from, to), value: move_cost
    let flowing: HashSet<_> = map
        .iter()
        .filter(|(_, valve)| valve.flow_rate > 0)
        .map(|(&name, _)| name)
        .collect();

    // key: opened, val: released_at_end
    let mut max_released_states: HashMap<BTreeSet<&str>, u32> = HashMap::new();

    let mut q = VecDeque::new();
    q.push_back(State {
        curr: "AA",
        opened: BTreeSet::new(),
        elapsed: 0,
        released: 0,
    });

    while let Some(State {
        opened,
        curr,
        elapsed,
        released,
    }) = q.pop_front()
    {
        let released_at_end = wait_until_ending(26, elapsed, released, &opened, &map);
        // record state. only update state if it beats the `released_at_end` number
        max_released_states
            .entry(opened.clone())
            .and_modify(|val| *val = released_at_end.max(*val))
            .or_insert(released_at_end);

        // if all flowing valves are opened or the timelimit was reached, skip
        if opened.len() == flowing.len() || elapsed >= 26 {
            continue;
        }
        // for every unopened valve, run simulation
        let unopened = flowing.iter().filter(|name| !opened.contains(*name));

        for dest in unopened {
            // how long would moving to dest take? +1 to open the valve
            let cost = dist_map[&(curr, *dest)] + 1;
            let new_elapsed = elapsed + cost;
            // if opening the dest valve would exceed the time limit, skip
            if new_elapsed >= 26 {
                continue;
            }

            // relieve pressure of opened valves while we move to dest and open it
            let released_per_min: u32 = opened.iter().map(|name| &map[name].flow_rate).sum();
            let new_released = released + (released_per_min * cost);

            // add opened valve to opened valves
            let mut new_opened = opened.clone();
            new_opened.insert(dest);

            q.push_back(State {
                opened: new_opened,
                curr: dest,
                elapsed: new_elapsed,
                released: new_released,
            });
        }
    }

    eprintln!(
        "{}",
        max_released_states
            .iter()
            .tuple_combinations()
            .filter(|(human, elephant)| human.0.is_disjoint(elephant.0))
            .map(|(human, elephant)| human.1 + elephant.1)
            .max()
            .unwrap(),
    );
    Ok(())
}
