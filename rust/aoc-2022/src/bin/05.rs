use itertools::Itertools;

type Stack = Vec<char>;
type Instruction = (i32, i32, i32);
type Input = (Vec<Stack>, Vec<Instruction>);

pub fn parse_input(input: &str) -> Input {
    let (config, instructions) = input.split_once("\n\n").unwrap();

    let stacks = parse_stack_configuration(config);
    let instructions = parse_instructions(instructions);

    (stacks, instructions)
}

fn parse_instructions(input: &str) -> Vec<Instruction> {
    input
        .lines()
        .map(|line| {
            let (n, origin, target) = line
                .split(" ")
                .skip(1)
                .step_by(2)
                .map(|n| n.parse::<i32>().unwrap())
                .collect_tuple()
                .unwrap();
            (n, origin - 1, target - 1)
        })
        .collect::<Vec<(i32, i32, i32)>>()
}

fn parse_stack_configuration(input: &str) -> Vec<Stack> {
    // 3 cols for a stack and 1 between them:
    // n_stacks * 3 + n_stacks - 1 = n_stacks * 4 - 1
    let n_stacks = (input.lines().last().unwrap().len() + 1) / 4;
    let mut stacks = vec![vec![]; n_stacks];
    input.lines().rev().skip(1).for_each(|line| {
        line.chars()
            .skip(1)
            .step_by(4)
            .enumerate()
            .filter(|(_i, c)| *c != ' ')
            .for_each(|(i, c)| stacks[i].push(c))
    });
    stacks
}

fn read_top_create(stacks: Vec<Stack>) -> String {
    stacks.iter().map(|s| s.last().unwrap()).collect::<String>()
}
pub fn part_one(input: Input) -> Option<String> {
    let (mut stacks, instructions) = input;
    instructions.iter().for_each(|(n, origin, target)| {
        (0..*n).for_each(|_| {
            let c = stacks[*origin as usize].pop().unwrap();
            stacks[*target as usize].push(c);
        });
    });

    Some(read_top_create(stacks))
}

pub fn part_two(input: Input) -> Option<String> {
    let (mut stacks, instructions) = input;
    instructions.iter().for_each(|(n, origin, target)| {
        let new_len = stacks[*origin as usize].len() - *n as usize;
        let mut to_be_moved = stacks[*origin as usize].split_off(new_len);
        stacks[*target as usize].append(&mut to_be_moved);
    });

    Some(read_top_create(stacks))
}

utils::main!(2022, 5);

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let input = utils::get_test_input(2022, 5);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 5, 1);
        assert_eq!(part_one(parsed_input), Some(expected));
    }

    #[test]
    fn test_part_two() {
        let input = utils::get_test_input(2022, 5);
        let parsed_input = parse_input(&input);
        let expected = utils::get_test_result(2022, 5, 2);
        assert_eq!(part_two(parsed_input), Some(expected));
    }
}
