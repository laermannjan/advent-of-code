use clap::{Parser, Subcommand};

// parser that reads the year as u32 and day as u8
#[derive(Parser, Debug)]
#[clap(author, version, about)]
pub struct Scaffold {
    #[clap(subcommand)]
    commands: Command,
}

#[derive(Subcommand, Debug)]
enum Command {
    #[clap(name = "year", about = "Create a new year")]
    Year {
        #[clap(help = "The year to create")]
        year: u32,
    },
    #[clap(name = "day", about = "Create a new day")]
    Day {
        #[clap(help = "The day to create")]
        day: u8,
    },
}

fn scaffold_year(year: u32) {
    println!("Creating year {}", year);
}

fn scaffold_day(day: u8) {
    println!("Creating day {}", day);
}

fn main() {
    let scaffold = Scaffold::parse();
    println!("{:?}", scaffold);

    match scaffold.commands {
        Command::Year { year } => scaffold_year(year),
        Command::Day { day } => scaffold_day(day),
    }
}
