use nom::{
    branch::alt,
    bytes::complete::{tag, take_while1},
    combinator::map,
    sequence::{preceded, separated_pair},
    IResult,
};
type Input = Node;

#[derive(Debug, Clone)]
pub enum Line {
    Command(Command),
    Entry(Entry),
}

#[derive(Debug, Clone)]
pub enum Command {
    Cd(String),
    Ls,
}

#[derive(Debug, Clone)]
pub enum Entry {
    Dir { name: String },
    File { size: u32, name: String },
}

// #[derive(Debug, Clone)]
// pub enum Path {
//     Root,
//     Parent,
//     Child(String),
// }

#[derive(Debug, Clone)]
pub struct Node {
    size: u32,
    children: Vec<Node>,
}

impl Node {
    fn total_size(&self) -> u32 {
        self.size + self.children.iter().map(|c| c.total_size()).sum::<u32>()
    }

    fn all_dirs(&self) -> Box<dyn Iterator<Item = &Node> + '_> {
        Box::new(
            std::iter::once(self).chain(
                self.children
                    .iter()
                    .filter(|c| !c.children.is_empty())
                    .flat_map(|c| c.all_dirs()),
            ),
        )
    }

    fn from(mut self, it: &mut dyn Iterator<Item = Line>) -> Self {
        while let Some(line) = it.next() {
            match line {
                Line::Command(Command::Cd(path)) => match path.as_str() {
                    "/" => {}
                    ".." => break,
                    _ => self.children.push(
                        Node {
                            size: 0,
                            children: vec![],
                        }
                        .from(it),
                    ),
                },
                Line::Entry(Entry::File { size, name: _ }) => {
                    self.children.push(Node {
                        size,
                        children: vec![],
                    });
                }
                _ => {}
            }
        }
        self
    }
}

fn parse_path(input: &str) -> IResult<&str, &str> {
    take_while1(|c: char| "abcdefghijklmnopqrstuvwxyz./".contains(c))(input)
}

impl Line {
    fn parse(input: &str) -> IResult<&str, Self> {
        alt((
            map(Command::parse, |c| Line::Command(c)),
            map(Entry::parse, |e| Line::Entry(e)),
        ))(input)
    }
}

impl Command {
    fn parse(input: &str) -> IResult<&str, Self> {
        let (input, _) = tag("$ ")(input)?;
        alt((
            map(preceded(tag("cd "), parse_path), |path| {
                Command::Cd(path.to_owned())
            }),
            map(tag("ls"), |_| Command::Ls),
        ))(input)
    }
}

impl Entry {
    fn parse(input: &str) -> IResult<&str, Self> {
        let parse_file = map(
            separated_pair(nom::character::complete::u32, tag(" "), parse_path),
            |(size, path)| Entry::File {
                size,
                name: path.to_owned(),
            },
        );
        let parse_dir = map(preceded(tag("dir "), parse_path), |path| Entry::Dir {
            name: path.to_owned(),
        });

        alt((parse_file, parse_dir))(input)
    }
}

pub fn parse_input(input: &str) -> Input {
    let mut lines = input.lines().map(|line| Line::parse(line).unwrap().1);

    let root = Node {
        size: 0,
        children: vec![],
    }
    .from(&mut lines);

    root
}
pub fn part_one(input: Input) -> Option<i32> {
    Some(
        input
            .all_dirs()
            .map(|d| d.total_size())
            .filter(|&s| s < 100_000)
            .sum::<u32>() as i32,
    )
}

pub fn part_two(input: Input) -> Option<i32> {
    let total_space = 70_000_000;
    let needed_space = 30_000_000;

    let available_space = total_space - input.total_size();
    let min_freeup_space = needed_space - available_space;

    Some(
        input
            .all_dirs()
            .map(|d| d.total_size())
            .filter(|&d| d >= min_freeup_space)
            .min()
            .unwrap() as i32,
    )
}

utils::main!(2022, 7);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 7);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 7, 1).parse().unwrap();
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 7);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 7, 2).parse().unwrap();
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
