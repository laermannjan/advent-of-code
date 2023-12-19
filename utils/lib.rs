use std::env;
use std::{error::Error, path::Path, time::Instant};

pub struct Day {
    pub loc: &'static str,
    pub part_one: fn(&str) -> Option<i32>,
    pub part_two: fn(&str) -> Option<i32>,
}

impl Day {
    fn dir(&self) -> &Path {
        Path::new(&self.loc).parent().unwrap()
    }
    pub fn run(&self) -> Result<(), Box<dyn Error>> {
        let mut input_file = "";
        let mut part = "";

        let args: Vec<String> = env::args().collect();
        let mut args_iter = args.iter().peekable();

        while let Some(arg) = args_iter.next() {
            match arg.as_str() {
                "--input" => {
                    if let Some(next_arg) = args_iter.peek() {
                        input_file = next_arg;
                        args_iter.next();
                    }
                }
                "--part" => {
                    if let Some(next_arg) = args_iter.peek() {
                        part = next_arg;
                        args_iter.next();
                    }
                }
                _ => {} // Ignore unrecognized arguments
            }
        }
        let input_path = &self.dir().join(input_file);

        let input = std::fs::read_to_string(input_path.clone())
            .expect(&format!("could not open data file {:?}", &input_path));

        if part == "one" {
            let start = Instant::now();
            let result = (&self.part_one)(input.as_str()).unwrap();
            let elapsed = start.elapsed();
            println!("part one: {:?} (took: {:?})", result, elapsed);
        }
        if part == "two" {
            let start = Instant::now();
            let result = (&self.part_two)(input.as_str()).unwrap();
            let elapsed = start.elapsed();
            println!("part two: {:?} (took: {:?})", result, elapsed);
        }

        Ok(())
    }
}
