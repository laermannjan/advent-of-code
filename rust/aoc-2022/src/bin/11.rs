use std::collections::HashMap;

use itertools::Itertools;
use nom::{bytes::complete::tag, character::complete::newline, sequence::preceded};

type Input = HashMap<usize, Monkey>;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Operation {
    Plus(i64),
    Times(i64),
    Square,
}

impl Operation {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        nom::branch::alt((
            nom::combinator::map(
                preceded(
                    tag("  Operation: new = old + "),
                    nom::character::complete::i64,
                ),
                |x| Operation::Plus(x),
            ),
            nom::combinator::map(
                preceded(
                    tag("  Operation: new = old * "),
                    nom::character::complete::i64,
                ),
                |x| Operation::Times(x),
            ),
            nom::combinator::map(tag("  Operation: new = old * old"), |_| Operation::Square),
        ))(input)
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Test {
    divisor: i64,
    true_target: usize,
    false_target: usize,
}

impl Test {
    fn parse(input: &str) -> nom::IResult<&str, Self> {
        let (input, divisor) =
            preceded(tag("  Test: divisible by "), nom::character::complete::i64)(input)?;
        let (input, _) = newline(input)?;
        let (input, true_target) = preceded(
            tag("    If true: throw to monkey "),
            nom::character::complete::u8,
        )(input)?;
        let (input, _) = newline(input)?;
        let (input, false_target) = preceded(
            tag("    If false: throw to monkey "),
            nom::character::complete::u8,
        )(input)?;

        Ok((
            input,
            Self {
                divisor,
                true_target: true_target as usize,
                false_target: false_target as usize,
            },
        ))
    }
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Monkey {
    items: Vec<i64>,
    operation: Operation,
    test: Test,
    inspected: u64,
}

impl Monkey {
    fn new(items: Vec<i64>, operation: Operation, test: Test) -> Self {
        Self {
            items,
            operation,
            test,
            inspected: 0,
        }
    }

    fn parse(input: &str) -> nom::IResult<&str, Self> {
        let (input, _monkey_no) = preceded(tag("Monkey "), nom::character::complete::u8)(input)?;
        let (input, _) = tag(":")(input)?;
        let (input, _) = newline(input)?;

        let (input, items) = preceded(
            tag("  Starting items: "),
            nom::multi::separated_list0(tag(", "), nom::character::complete::i64),
        )(input)?;
        let (input, _) = newline(input)?;
        let (input, operation) = Operation::parse(input)?;
        let (input, _) = newline(input)?;
        let (input, test) = Test::parse(input)?;

        Ok((
            input,
            Self::new(items.into_iter().rev().collect(), operation, test),
        ))
    }

    fn inspect_item(&mut self, worry_squasher: impl Fn(i64) -> i64) -> (usize, i64) {
        self.inspected += 1;
        let mut worry_level = self.items.pop().unwrap();
        match self.operation {
            Operation::Plus(x) => worry_level += x,
            Operation::Times(x) => worry_level *= x,
            Operation::Square => worry_level *= worry_level,
        }
        worry_level = worry_squasher(worry_level);

        let monkey_no = if worry_level % self.test.divisor == 0 {
            self.test.true_target
        } else {
            self.test.false_target
        };

        (monkey_no, worry_level)
    }
}

pub fn parse_input(input: &str) -> Input {
    input
        .split("\n\n")
        .enumerate()
        .map(|(i, m)| (i, Monkey::parse(m).unwrap().1))
        .collect::<HashMap<usize, Monkey>>()
}

pub fn part_one(input: Input) -> Option<i64> {
    let mut monkeys = input.clone();
    let n_monkeys = monkeys.len();

    for _round in 0..20 {
        for this_monkey in 0..n_monkeys {
            for _ in 0..monkeys[&this_monkey].items.len() {
                let (target_monkey, item) = monkeys
                    .get_mut(&this_monkey)
                    .unwrap()
                    .inspect_item(|x| x / 3);
                monkeys.get_mut(&target_monkey).unwrap().items.push(item);
            }
        }
    }

    let monkey_business = monkeys
        .iter()
        .map(|(_, m)| m.inspected)
        .sorted()
        .rev()
        .take(2)
        .collect_vec();

    Some(monkey_business.iter().product::<u64>() as i64)
}

pub fn part_two(input: Input) -> Option<i64> {
    let mut monkeys = input.clone();
    let n_monkeys = monkeys.len();
    let divisor_product = monkeys.iter().map(|(_, m)| m.test.divisor).product::<i64>();

    for _round in 0..10_000 {
        for this_monkey in 0..n_monkeys {
            for _ in 0..monkeys[&this_monkey].items.len() {
                let (target_monkey, item) = monkeys
                    .get_mut(&this_monkey)
                    .unwrap()
                    .inspect_item(|x| x % divisor_product);
                monkeys.get_mut(&target_monkey).unwrap().items.push(item);
            }
        }
    }

    let monkey_business = monkeys
        .iter()
        .map(|(_, m)| m.inspected)
        .sorted()
        .rev()
        .take(2)
        .collect_vec();

    Some(monkey_business.iter().product::<u64>() as i64)
}

utils::main!(2022, 11);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 11, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 11, 1, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let test_no = 1;
        let input = utils::get_test_input(2022, 11, test_no);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 11, 2, test_no)
            .parse()
            .unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
