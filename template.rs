pub fn part_one(input: String) -> Option<i32> {
    None
}

pub fn part_two(input: String) -> Option<i32> {
    None
}

pub fn main() {
    let day = Day {
        loc: Path::new(file!()).into(),
        part_one,
        part_two,
    };
    let _ = day.run();
}
