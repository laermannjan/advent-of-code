use std::cmp::Ordering;
use std::io::{self, Read};

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

pub fn parse_input() -> Input {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    input.split("\n\n").map(Pair::parse).collect()
}
pub fn main() -> io::Result<()> {
    let input = parse_input();
    eprintln!(
        "{}",
        input
            .iter()
            .enumerate()
            .map(|(i, p)| if p.left < p.right { i + 1 } else { 0 })
            .sum::<usize>(),
    );
    Ok(())
}
