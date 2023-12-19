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
        let mut input_file = String::from("input.txt");
        let mut example = 0;
        let mut only_one = false;
        let mut only_two = false;

        let args: Vec<String> = env::args().collect();
        let mut args_iter = args.iter().peekable();

        while let Some(arg) = args_iter.next() {
            match arg.as_str() {
                "--input" => {
                    if let Some(next_arg) = args_iter.peek() {
                        input_file = next_arg.to_string();
                        args_iter.next(); // Advance the iterator since we've used this argument
                    }
                }
                "--one" => only_one = true,
                "--two" => only_two = true,

                "--example" => {
                    example = if let Some(next_arg) = args_iter.peek() {
                        // Check if the next argument is a number
                        if next_arg.parse::<i32>().is_ok() {
                            args_iter.next().and_then(|n| n.parse().ok()).unwrap_or(1)
                        } else {
                            1 // Default to 1 if next argument is not a number
                        }
                    } else {
                        1 // Default to 1 if no argument follows
                    };
                }

                _ => {} // Ignore unrecognized arguments or handle them as needed
            }
        }

        if example != 0 {
            input_file = format!("example{}.txt", example).to_string()
        };

        if !&self.dir().join(input_file.clone()).exists() {
            input_file = "example.txt".to_string();
        }
        let input_path = &self.dir().join(input_file);

        let input = std::fs::read_to_string(input_path.clone())
            .expect(&format!("could not open data file {:?}", &input_path));

        if !only_two {
            let start = Instant::now();
            let result = (&self.part_one)(input.as_str()).unwrap();
            let elapsed = start.elapsed();
            println!("part one: {:?} (took: {:?})", result, elapsed);
        }
        if !only_one {
            let start = Instant::now();
            let result = (&self.part_two)(input.as_str()).unwrap();
            let elapsed = start.elapsed();
            println!("part two: {:?} (took: {:?})", result, elapsed);
        }

        Ok(())
    }
}
