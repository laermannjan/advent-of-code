use itertools::Itertools;

type Input = Vec<(u8, u32)>;

fn parse_input(input: &str) -> Input {
    input
        .split_whitespace()
        .tuples()
        .map(|(inst, val)| (inst.as_bytes()[0], val.parse().unwrap()))
        .collect::<Vec<_>>()
}

pub fn part_one(input: Input) -> Option<i32> {
    let (horizontal, depth) = input
        .iter()
        .fold((0, 0), |(horizontal, depth), (inst, val)| match inst {
            b'f' => (horizontal + val, depth),
            b'u' => (horizontal, depth - val),
            b'd' => (horizontal, depth + val),
            _ => unreachable!(),
        });
    Some((horizontal * depth) as i32)
}

pub fn part_two(input: Input) -> Option<i32> {
    let (horizontal, depth, _) = input.iter().fold(
        (0, 0, 0),
        |(horizontal, depth, aim), (inst, val)| match inst {
            b'f' => (horizontal + val, depth + aim * val, aim),
            b'u' => (horizontal, depth, aim - val),
            b'd' => (horizontal, depth, aim + val),
            _ => unreachable!(),
        },
    );
    Some((horizontal * depth) as i32)
}

utils::main!(2021, 2);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2021, 2);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 2, 1);
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2021, 2);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 2, 2);
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
