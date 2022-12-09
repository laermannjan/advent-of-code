use itertools::Itertools;

type Input = (Vec<usize>, Vec<Vec<Vec<usize>>>);

fn parse_input(input: &str) -> Input {
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

    (draws, boards)
}

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

pub fn part_one(input: Input) -> Option<i32> {
    let (draws, boards) = input;
    for num_draws in 5..draws.len() {
        let winners = boards
            .iter()
            .filter(|b| has_won(b, &draws[0..num_draws]))
            .collect_vec();
        if winners.len() >= 1 {
            return Some(score(winners[0], &draws[0..num_draws]) as i32);
        }
    }
    unreachable!()
}
pub fn part_two(input: Input) -> Option<i32> {
    let (draws, boards) = input;

    // only consider boards that will win eventually
    let boards = boards.iter().filter(|b| has_won(b, &draws)).collect_vec();

    // going backwards, reducing the number of draws to consider, find the first losing board
    for num_draws in (5..draws.len()).rev() {
        let losers = boards
            .iter()
            .filter(|b| !has_won(b, &draws[0..num_draws]))
            .collect_vec();
        if losers.len() >= 1 {
            // + 1 because the first losing board would win on the next draw
            return Some(score(losers[0], &draws[0..num_draws + 1]) as i32);
        }
    }
    unreachable!()
}

utils::main!(2021, 4);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2021, 4);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 4, 1);
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2021, 4);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2021, 4, 2);
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
