# [🎅 Advent of Code 💻](http://adventofcode.com)

## Quick start

```bash
aoc load                    # scaffold next day
aoc run 2b > /dev/null      # test solution day 2 part b; hide debug prints
aoc run                     # test latest solution
hyperfine 'aoc run input'   # benchmark latest solution

```

## Project Structure

```
.
├── 20XX/
│   ├── day01/
│   │   ├── input.txt
│   │   ├── example.txt
│   │   ├── example2.txt
│   │   ├── a.py
│   │   └── b.go
│   └── day02/
│       ├── input.txt
│       ├── example.txt
│       ├── a.py
│       └── b.py
└── aoc
```

## Solution File Requirements

Solution files should:
1. Accept input via stdin
2. Print the answer as the last line of stderr (to be able to use prints for debugging and run the files with `> /dev/null` to skip those)
3. Be executable from any directory (see "Solution Independence" below)

## Usage

### Environment Variables

- `AOC_YEAR`: The year to work with (defaults to current year)
- `AOC_LANG`: The language suffix to use for solutions (defaults to "py")
- `AOC_COOKIE`: Your session cookie from adventofcode.com (required for downloading inputs)

To get your session cookie:
1. Log into [Advent of Code](https://adventofcode.com)
2. Open your browser's developer tools (usually F12)
3. Go to the Storage or Application tab
4. Look for Cookies and find the 'session' value
5. Export it: `export AOC_COOKIE='your-cookie-value'` (e.g., in `.envrc`)

### Commands

```bash
# Load input for a day
aoc load [DAY]          # Load input for specified day (or next day if omitted) and create solution files from template"

# Run a solution
aoc run [DAY_PART] [INPUT_TYPE]
# Examples:
aoc run                    # Run latest solution with example.txt
aoc run input              # Run latest solution with input.txt
aoc run 5a                 # Run day 5 part a with example.txt
aoc run 5a input           # Run day 5 part a with input.txt
aoc run 13b example2       # Run day 13 part b with example2.txt
aoc run < custom.txt       # Run latest solution with an arbitary input file
echo '1 2 3' | aoc run 5a  # Pipe input to day 5 part a
```

## Solution Independence

There are two recommended approaches for maintaining solution independence:

1. **Self-contained Solutions**: Each solution file contains all necessary code without external dependencies. This is the simplest approach but may lead to code duplication.

2. **Relative Imports**: Solutions can use utility functions by importing them relative to the solution file's location. Example:

```python
# In 2023/day1/a.py
import os
import sys
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
from utils import helper_function
```

Both approaches ensure solutions can be run from any directory while maintaining code organization flexibility.

## Why This Design?

1. **Minimal Lock-in**: Solutions remain runnable without the CLI:
   ```bash
   python3 2023/day1/a.py < 2023/day1/input.txt
   ```

2. **Language Agnostic**: Supports any language that can read from stdin and write to stdout.

3. **Flexible Testing**: Easy to test with example inputs or custom test cases.

4. **Progressive Enhancement**: The CLI adds convenience without imposing requirements on solution structure.
