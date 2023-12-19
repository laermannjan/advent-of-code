# [🎅 Advent of Code 💻](http://adventofcode.com)

## Requirements
* [aoc-cli](https://github.com/scarvalhojr/aoc-cli) (for the puzzle and input download)
* python 3.12+
* requirements in `Cargo.toml`, `go.mod`, `Pipfile`, etc.

## Usage
```
 $ python aoc.py --help

 Usage: aoc.py [OPTIONS] COMMAND [ARGS]...

╭─ Options ───────────────────────────────────────────────────────────────────────╮
│ --install-completion          Install completion for the current shell.         │
│ --show-completion             Show completion for the current shell, to copy it │
│                               or customize the installation.                    │
│ --help                        Show this message and exit.                       │
╰─────────────────────────────────────────────────────────────────────────────────╯
╭─ Commands ──────────────────────────────────────────────────────────────────────╮
│ download  Download the puzzle input and description (converted to markdown).    │
│ scaffold  Scaffold the day's folder from the language's template file.          │
│           Downloads input & description if not present.                         │
│ solve     Run the selected language's solver for the given year, day, and part. │
│ test      Run tests (if any) for the specified year and day.                    │
╰─────────────────────────────────────────────────────────────────────────────────╯
```

## TODO
Update previous years and languages to new structure
- [ ] python
- [ ] rust
- [ ] haskell
