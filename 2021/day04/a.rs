use itertools::Itertools;
use std::io::{self, Read};

fn score(board: &Vec<Vec<usize>>, draws: &[usize]) -> usize {
    let unmarked_sum = board
        .iter()
        .flatten()
        .filter(|x| !draws.contains(x))
        .sum::<usize>();
    let last_draw = draws.last().unwrap();
    unmarked_sum * last_draw
}

fn has_won(board: &Vec<Vec<usize>>, draws: &[usize]) -> bool {
    for row in 0..board.len() {
        if (0..board[0].len()).all(|col| draws.contains(&board[row][col])) {
            return true;
        }
    }

    for col in 0..board[0].len() {
        if (0..board.len()).all(|row| draws.contains(&board[row][col])) {
            return true;
        }
    }
    false
}

fn main() -> io::Result<()> {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input)?;


    let mut blocks = input.split("\n\n");
    let draws = blocks
        .next()
        .unwrap()
        .split(',')
        .map(|d| d.parse::<usize>().unwrap())
        .collect_vec();

    let boards = blocks
        .map(|b| {
            b.lines()
                .map(|l| {
                    l.split_whitespace()
                        .map(|d| d.parse::<usize>().unwrap())
                        .collect_vec()
                })
                .collect_vec()
        })
        .collect_vec();

    for num_draws in 5..draws.len() {
        let winners = boards
            .iter()
            .filter(|b| has_won(b, &draws[0..num_draws]))
            .collect_vec();
        if winners.len() >= 1 {
            let result = score(winners[0], &draws[0..num_draws]) as i32;
            eprintln!("{}", result);
            return Ok(())
        }
    }
    unreachable!()

}
