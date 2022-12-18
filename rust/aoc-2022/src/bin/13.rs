use std::cmp::Ordering;

use nom::{branch, IResult};

type Input = Vec<Pair>;

#[derive(Clone, PartialEq, Eq)]
pub enum Element {
    Integer(i32),
    List(Vec<Element>),
}

impl Ord for Element {
    fn cmp(&self, other: &Self) -> Ordering {
        self.partial_cmp(other).unwrap()
    }
}

impl PartialOrd for Element {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        match (self, other) {
            (Element::Integer(a), Element::Integer(b)) => a.partial_cmp(b),
            (Element::List(a), Element::List(b)) => a.partial_cmp(b),
            (Element::Integer(_), Element::List(b)) => vec![self.clone()].partial_cmp(b),
            (Element::List(a), Element::Integer(_)) => a.partial_cmp(&vec![other.clone()]),
        }
    }
}

impl Element {
    fn parse(input: &str) -> IResult<&str, Self> {
        //parse vec of i32

        // parse lit of lists or integers
        branch::alt((
            nom::combinator::map(nom::character::complete::i32, Element::Integer),
            nom::combinator::map(
                nom::sequence::delimited(
                    nom::character::complete::char('['),
                    nom::multi::separated_list0(
                        nom::character::complete::char(','),
                        Element::parse,
                    ),
                    nom::character::complete::char(']'),
                ),
                Element::List,
            ),
        ))(input)
    }
}
impl std::fmt::Debug for Element {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Element::Integer(i) => write!(f, "{}", i),
            Element::List(l) => write!(f, "{:?}", l),
        }
    }
}

#[derive(Clone)]
pub struct Pair {
    left: Element,
    right: Element,
}

impl Pair {
    fn parse(input: &str) -> Self {
        let (left, right) = input.split_once("\n").unwrap();
        let left = Element::parse(left).unwrap().1;
        let right = Element::parse(right).unwrap().1;

        Self { left, right }
    }
}

impl std::fmt::Debug for Pair {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}\n{:?}", self.left, self.right)
    }
}

pub fn parse_input(input: &str) -> Input {
    input.split("\n\n").map(Pair::parse).collect()
}
pub fn part_one(input: Input) -> Option<usize> {
    Some(
        input
            .iter()
            .enumerate()
            .map(|(i, p)| if p.left < p.right { i + 1 } else { 0 })
            .sum(),
    )
}

pub fn part_two(input: Input) -> Option<i32> {
    let mut packets = input
        .iter()
        .map(|p| vec![p.left.clone(), p.right.clone()])
        .flatten()
        .collect::<Vec<_>>();

    packets.push(Element::parse("[[6]]").unwrap().1);
    packets.push(Element::parse("[[2]]").unwrap().1);
    packets.sort();

    let pos_6 = packets
        .iter()
        .position(|p| p == &Element::parse("[[6]]").unwrap().1)
        .map(|i| i as i32 + 1);
    let pos_2 = packets
        .iter()
        .position(|p| p == &Element::parse("[[2]]").unwrap().1)
        .map(|i| i as i32 + 1);

    Some(pos_6.unwrap() * pos_2.unwrap())
}

utils::main!(2022, 13);

#[cfg(test)]
mod tests {
    use std::cmp::Ordering;

    use super::*;

    #[test]
    fn test() {
        let a = vec![1, 1, 3, 1, 1];
        let b = vec![1, 1, 3, 1, 1];

        assert_eq!(a.cmp(&b), Ordering::Less);
    }

    #[test]
    fn test_part_one() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 13, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 13, 1, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 13, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 13, 2, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
