use itertools::Itertools;

fn parse_input(input: &str) -> Vec<(u8, u32)> {
    input
        .split_whitespace()
        .tuples()
        .map(|(inst, val)| (inst.as_bytes()[0], val.parse().unwrap()))
        .collect::<Vec<_>>()
}

pub fn day02a(input: &str) -> usize {
    let (horizontal, depth) =
        parse_input(input)
            .iter()
            .fold((0, 0), |(horizontal, depth), (inst, val)| match inst {
                b'f' => (horizontal + val, depth),
                b'u' => (horizontal, depth - val),
                b'd' => (horizontal, depth + val),
                _ => unreachable!(),
            });
    return (horizontal * depth) as usize;
}

pub fn day02b(input: &str) -> usize {
    let (horizontal, depth, _) = parse_input(input).iter().fold(
        (0, 0, 0),
        |(horizontal, depth, aim), (inst, val)| match inst {
            b'f' => (horizontal + val, depth + aim * val, aim),
            b'u' => (horizontal, depth, aim - val),
            b'd' => (horizontal, depth, aim + val),
            _ => unreachable!(),
        },
    );
    return (horizontal * depth) as usize;
}
