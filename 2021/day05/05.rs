use itertools::Itertools;
use std::collections::HashMap;

type Input = Vec<((i32, i32), (i32, i32))>;

fn parse_input(input: &str) -> Input {
    input
        .lines()
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
        .collect_vec()
}

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

pub fn part_one(input: Input) -> Option<i32> {
    let lines = input;
    let points = point_map(
        lines
            .iter()
            .copied()
            .filter(|((x1, y1), (x2, y2))| x1 == x2 || y1 == y2),
    );
    Some(count_overlaps(points, 2) as i32)
}
pub fn part_two(input: Input) -> Option<i32> {
    let lines = input;
    let points = point_map(lines.iter().copied());
    Some(count_overlaps(points, 2) as i32)
}

utils::main!(2021, 5);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2021, 5, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 5, 1, 1).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2021, 5, 1);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 5, 2, 1).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
