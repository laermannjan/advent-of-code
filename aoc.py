import shutil
import subprocess
from datetime import date
from enum import StrEnum
from pathlib import Path
from typing import Annotated

import rich
import typer

app = typer.Typer()
root_path = Path(__file__).resolve().parent
default_input = Path("./input.txt")


class Language(StrEnum):
    go = "go"
    rust = "rust"
    python = "python"


class PuzzlePart(StrEnum):
    one = "one"
    two = "two"


Lang = Annotated[Language, typer.Option(envvar="AOC_LANG")]

Year = Annotated[
    int,
    typer.Option(
        envvar="AOC_YEAR",
        min=2015,
        max=date.today().year,
        show_default=False,
    ),
]

Day = Annotated[
    int,
    typer.Option(
        min=1,
        max=25,
        show_default=False,
    ),
]

Part = Annotated[PuzzlePart, typer.Option()]

Input = Annotated[
    Path,
    typer.Option(
        help="input file path. takes absolute or path relative to cwd. [default: ./<year>/day<day>/input.txt (relative to this file's parent directory)]",
        rich_help_panel="Input",
        show_default=False,
    ),
]

Example = Annotated[
    int,
    typer.Option(
        "--example",
        "-e",
        help="use example input. example: -eee selects example3.txt. (replaces 'input' with 'example' in --input)",
        rich_help_panel="Input",
        count=True,
        show_default=False,
    ),
]

Verbose = Annotated[
    bool,
    typer.Option("--verbose", "-v", help="show verbose output"),
]


@app.command()
def solve(
    lang: Lang,
    year: Year,
    day: Day,
    part: Part,
    input: Input = default_input,
    example: Example = 0,
    verbose: Verbose = False,
    release: Annotated[
        bool, typer.Option(help="enable compiler/runtime optimizations")
    ] = False,
):
    """
    Run the selected language's solver for the given year, day, and part.

    Optionally, use (one of) the example or another specific input file.
    If available, enable compiler optimization, which reduce runtime,
    but might decrease developer experience.
    """
    if not input.is_absolute():
        if input == default_input:
            input = root_path / str(year) / f"day{day:02d}" / input
        else:
            input = input.resolve()

    if example:
        print("replacing with example<n>")
        input = input.with_stem(f"example{example}")
    if not input.is_file():
        print("replacing with example")
        input = input.with_stem("example")

    command = ""
    day_dir = root_path / str(year) / f"day{day:02d}"
    match lang:
        case Language.go:
            bin = day_dir / "main.go"
            command = f"go run {bin}"
        case Language.rust:
            cmd = "cargo run"
            if release:
                cmd = f"{cmd} --release"
            command = f"{cmd} --bin {year}-{day:02d} --"
        case Language.python:
            print("running python")

    command = f"{command} --input {str(input)} --part {str(part)}"

    if verbose:
        command = f"{command} --verbose"

    rich.print(f"[italic]{command=}[/italic]")
    subprocess.run(command, shell=True)


@app.command()
def test(
    lang: Lang,
    year: Year,
    day: Day,
):
    """
    Run tests (if any) for the specified year and day.
    """
    command = ""
    day_dir = root_path / str(year) / f"day{day:02d}"
    match lang:
        case Language.go:
            bin = day_dir / "main.go"
            command = f"go run {bin}"
        case Language.rust:
            command = f"cargo test --bin {year}-{day:02d} --"
        case Language.python:
            print("running python")
    print(f"running {command=}")

    subprocess.run(command, shell=True)


@app.command()
def scaffold(
    lang: Lang,
    year: Year,
    day: Day,
):
    """
    Scaffold the day's folder from the language's template file. Downloads input & description if not present.
    """
    day_dir = root_path / str(year) / f"day{day:02d}"
    match lang:
        case Language.go:
            template = root_path / "template.go"
            bin = day_dir / "main.go"
        case Language.rust:
            template = root_path / "template.rs"
            bin = day_dir / "main.rs"
            rich.print("[gray italic]add [[bin]] entry to Cargo.toml")
            with open((root_path / "Cargo.toml"), "a") as f:
                f.write(
                    f"""\
[[bin]]
name="{year}-{day:02d}"
path="{year}/{day:02d}/main.rs"
"""
                )
        case Language.python:
            template = root_path / "template.py"
            bin = day_dir / "main.py"

    day_dir.mkdir(exist_ok=True)

    if bin.exists():
        rich.print(f"[red]{bin} already exists, aborting.")
        return

    shutil.copy2(template, bin)
    rich.print(f"[gray italic]copied {template} to {bin}")

    if not (day_dir / "input.txt").exists():
        download(year=year, day=day)
    if (
        not (day_dir / "example.txt").exists()
        and not (day_dir / "example1.txt").exists()
    ):
        subprocess.call(f"$EDITOR {day_dir/'example.txt'}", shell=True)


@app.command()
def download(
    year: Year,
    day: Day,
):
    """
    Download the puzzle input and description (converted to markdown).

    Depends on aoc-cli.
    """
    day_dir = root_path / str(year) / f"day{day:02d}"
    subprocess.run(
        f"aoc download --{year=} --{day=} --input-file {day_dir}/input.txt --puzzle-file {day_dir}/puzzle.md",
        shell=True,
    )


if __name__ == "__main__":
    app()
