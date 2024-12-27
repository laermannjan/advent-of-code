use std::io::{self, Read};

use id_tree::InsertBehavior::*;
use id_tree::*;
use nom::{
    branch::alt,
    bytes::complete::{tag, take_while1},
    combinator::map,
    sequence::{preceded, separated_pair},
    IResult,
};
type Input = Tree<FsNode>;

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
pub struct FsNode {
    name: String,
    size: u32,
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

pub fn parse_input() -> Input {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();

    let lines = input.lines().map(|line| Line::parse(line).unwrap().1);

    let mut tree = Tree::<FsNode>::new();
    let root_id = tree
        .insert(
            Node::new(FsNode {
                name: "/".to_owned(),
                size: 0,
            }),
            AsRoot,
        )
        .unwrap();
    let mut curr_id = root_id.clone();

    for line in lines {
        match line {
            Line::Command(Command::Cd(path)) => match path.as_str() {
                "/" => {
                    curr_id = root_id.clone();
                }
                ".." => {
                    curr_id = tree.get(&curr_id).unwrap().parent().unwrap().clone();
                }
                dir => {
                    let node = FsNode {
                        name: dir.to_owned(),
                        size: 0,
                    };
                    if let Some(child_id) = tree
                        .get(&curr_id)
                        .unwrap()
                        .children()
                        .iter()
                        .find(|child| tree.get(child).unwrap().data().name == node.name)
                    {
                        curr_id = child_id.clone();
                    } else {
                        curr_id = tree.insert(Node::new(node), UnderNode(&curr_id)).unwrap();
                    }
                }
            },
            Line::Entry(Entry::File { size, name }) => {
                let node = FsNode {
                    name: name.to_owned(),
                    size,
                };
                tree.insert(Node::new(node), UnderNode(&curr_id)).unwrap();
            }
            _ => {}
        }
    }

    for node_id in &mut tree
        .traverse_post_order_ids(tree.root_node_id().unwrap())
        .unwrap()
    {
        tree.get_mut(&node_id).unwrap().data_mut().size += tree
            .children(&node_id)
            .unwrap()
            .map(|child| child.data().size)
            .sum::<u32>();
    }

    // let mut s = String::new();
    // tree.write_formatted(&mut s).unwrap();
    // println!("{s}");

    tree
}

pub fn main() -> io::Result<()> {
    let input = parse_input();

    let total_space = 70_000_000;
    let needed_space = 30_000_000;

    let available_space = total_space
        - input
            .get(input.root_node_id().unwrap())
            .unwrap()
            .data()
            .size;
    let min_freeup_space = needed_space - available_space;

    eprintln!(
        "{}",
        input
            .traverse_pre_order(input.root_node_id().unwrap())
            .unwrap()
            .filter(|&n| n.children().len() > 0 && n.data().size >= min_freeup_space)
            // .inspect(|node| {
            //     dbg!(node);
            // })
            .map(|node| node.data().size)
            .min()
            .unwrap() as i32,
    );
    Ok(())
}
