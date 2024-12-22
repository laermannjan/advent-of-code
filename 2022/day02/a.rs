use std::io::{self, Read};
use std::{cmp::Ordering, str::FromStr};

#[derive(Debug, PartialEq)]
enum Shape {
    Rock,
    Paper,
    Scissors,
}

impl PartialOrd for Shape {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        match (self, other) {
            (Shape::Rock, Shape::Paper) => Some(Ordering::Less),
            (Shape::Rock, Shape::Scissors) => Some(Ordering::Greater),
            (Shape::Paper, Shape::Rock) => Some(Ordering::Greater),
            (Shape::Paper, Shape::Scissors) => Some(Ordering::Less),
            (Shape::Scissors, Shape::Rock) => Some(Ordering::Less),
            (Shape::Scissors, Shape::Paper) => Some(Ordering::Greater),
            (a, b) if a == b => Some(Ordering::Equal),
            _ => None,
        }
    }
}

impl FromStr for Shape {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "A" => Ok(Shape::Rock),
            "B" => Ok(Shape::Paper),
            "C" => Ok(Shape::Scissors),
            _ => Err("Invalid shape".to_string()),
        }
    }
}

struct Round {
    our_shape: Shape,
    their_shape: Shape,
}

impl Round {
    fn points_won(&self) -> u32 {
        let round_points = match self.our_shape.partial_cmp(&self.their_shape) {
            Some(Ordering::Greater) => 6,
            Some(Ordering::Equal) => 3,
            Some(Ordering::Less) => 0,
            None => panic!("Invalid shapes"),
        };

        let shape_points = match self.our_shape {
            Shape::Rock => 1,
            Shape::Paper => 2,
            Shape::Scissors => 3,
        };

        round_points + shape_points
    }
}

impl FromStr for Round {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut shapes = s.split_whitespace();
        let our_shape = shapes.next().unwrap().parse()?;
        let their_shape = shapes.next().unwrap().parse()?;

        Ok(Round {
            our_shape,
            their_shape,
        })
    }
}

pub fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;

    let result = input
        .lines()
        .map(|line| {
            let mut shapes = line.split_whitespace();
            let their_shape = shapes.next().unwrap().parse().unwrap();
            let our_shape = match shapes.next().unwrap() {
                "X" => Shape::Rock,
                "Y" => Shape::Paper,
                "Z" => Shape::Scissors,
                _ => panic!("Invalid shape"),
            };
            Round {
                our_shape,
                their_shape,
            }
            .points_won() as i32
        })
        .sum::<i32>();

    eprintln!("{}", result as i32);
    Ok(())
}

