use itertools::Itertools;
use std::io::{self, Read};
use std::collections::HashMap;

fn point_map(lines: impl Iterator<Item = ((i32, i32), (i32, i32))>) -> HashMap<(i32, i32), u32> {
    let mut points = HashMap::new();
    for ((x1, y1), (x2, y2)) in lines {
        let dx = (x2 - x1).signum();
        let dy = (y2 - y1).signum();
        let (mut x, mut y) = (x1, y1);
        // don't forget (x2, y2)
        while (x, y) != (x2 + dx, y2 + dy) {
            *points.entry((x, y)).or_insert(0) += 1;
            x += dx;
            y += dy;
        }
    }
    return points;
}

fn count_overlaps(points: HashMap<(i32, i32), u32>, t: u32) -> usize {
    points.values().filter(|&&val| val >= t).count()
}

fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let lines = input.lines()
        .map(|l| {
            l.split(" -> ")
                .map(|p| {
                    p.split(',')
                        .map(|n| n.parse::<i32>().unwrap())
                        .collect_tuple::<(i32, i32)>()
                        .unwrap()
                })
                .collect_tuple::<((i32, i32), (i32, i32))>()
                .unwrap()
        })
        .collect_vec();
    let points = point_map(
        lines
            .iter()
            .copied()
            .filter(|((x1, y1), (x2, y2))| x1 == x2 || y1 == y2),
    );

    eprintln!("{}", count_overlaps(points, 2) as i32);
    Ok(())
}
