module Main where

import AdventOfCode.Utils (getInput)
import qualified AdventOfCode.Year2021.Day01 as Y21
import qualified AdventOfCode.Year2021.Day02 as Y21
import qualified AdventOfCode.Year2021.Day03 as Y21
import qualified AdventOfCode.Year2021.Day04 as Y21
import qualified AdventOfCode.Year2021.Day05 as Y21
import qualified AdventOfCode.Year2021.Day06 as Y21
import qualified AdventOfCode.Year2021.Day07 as Y21
import qualified AdventOfCode.Year2021.Day08 as Y21
import qualified AdventOfCode.Year2021.Day09 as Y21
import qualified AdventOfCode.Year2021.Day10 as Y21
import qualified AdventOfCode.Year2021.Day11 as Y21
import qualified AdventOfCode.Year2021.Day12 as Y21
import qualified AdventOfCode.Year2021.Day13 as Y21
import qualified AdventOfCode.Year2021.Day14 as Y21
import qualified AdventOfCode.Year2021.Day15 as Y21
import qualified AdventOfCode.Year2021.Day16 as Y21
import qualified AdventOfCode.Year2021.Day17 as Y21
import qualified AdventOfCode.Year2021.Day18 as Y21
import qualified AdventOfCode.Year2021.Day19 as Y21
import qualified AdventOfCode.Year2021.Day20 as Y21
import qualified AdventOfCode.Year2021.Day21 as Y21
import qualified AdventOfCode.Year2021.Day22 as Y21
import qualified AdventOfCode.Year2021.Day23 as Y21
import qualified AdventOfCode.Year2021.Day24 as Y21
import qualified AdventOfCode.Year2021.Day25 as Y21
import System.Environment
import System.TimeIt
import Text.Printf (printf)

main :: IO ()
main = do
  args <- getArgs
  case args of
    [] -> usage
    [_] -> usage
    (year : xs) -> case year of
      "2021" -> mapM_ run2021 xs
      _ -> printf "Year '%s' does not exist.\n" year

usage :: IO ()
usage = putStrLn "Usage: adventofcode <year> {1a | 1b | ... | 25b}"

run2021 :: String -> IO ()
run2021 "1a" = timeIt $ getInput "../data/2021/day01/day01_input_puzzle.txt" >>= print . Y21.day01a
run2021 "1b" = timeIt $ getInput "../data/2021/day01/day01_input_puzzle.txt" >>= print . Y21.day01b
run2021 "2a" = timeIt $ getInput "../data/2021/day02/day02_input_puzzle.txt" >>= print . Y21.day02a
run2021 "2b" = timeIt $ getInput "../data/2021/day02/day02_input_puzzle.txt" >>= print . Y21.day02b
run2021 "3a" = timeIt $ getInput "../data/2021/day03/day03_input_puzzle.txt" >>= print . Y21.day03a
run2021 "3b" = timeIt $ getInput "../data/2021/day03/day03_input_puzzle.txt" >>= print . Y21.day03b
run2021 "4a" = timeIt $ getInput "../data/2021/day04/day04_input_puzzle.txt" >>= print . Y21.day04a
run2021 "4b" = timeIt $ getInput "../data/2021/day04/day04_input_puzzle.txt" >>= print . Y21.day04b
run2021 "5a" = timeIt $ getInput "../data/2021/day05/day05_input_puzzle.txt" >>= print . Y21.day05a
run2021 "5b" = timeIt $ getInput "../data/2021/day05/day05_input_puzzle.txt" >>= print . Y21.day05b
run2021 "6a" = timeIt $ getInput "../data/2021/day06/day06_input_puzzle.txt" >>= print . Y21.day06a
run2021 "6b" = timeIt $ getInput "../data/2021/day06/day06_input_puzzle.txt" >>= print . Y21.day06b
run2021 "7a" = timeIt $ getInput "../data/2021/day07/day07_input_puzzle.txt" >>= print . Y21.day07a
run2021 "7b" = timeIt $ getInput "../data/2021/day07/day07_input_puzzle.txt" >>= print . Y21.day07b
run2021 "8a" = timeIt $ getInput "../data/2021/day08/day08_input_puzzle.txt" >>= print . Y21.day08a
run2021 "8b" = timeIt $ getInput "../data/2021/day08/day08_input_puzzle.txt" >>= print . Y21.day08b
run2021 "9a" = timeIt $ getInput "../data/2021/day09/day09_input_puzzle.txt" >>= print . Y21.day09a
run2021 "9b" = timeIt $ getInput "../data/2021/day09/day09_input_puzzle.txt" >>= print . Y21.day09b
run2021 "10a" = timeIt $ getInput "../data/2021/day10/day10_input_puzzle.txt" >>= print . Y21.day10a
run2021 "10b" = timeIt $ getInput "../data/2021/day10/day10_input_puzzle.txt" >>= print . Y21.day10b
run2021 "11a" = timeIt $ getInput "../data/2021/day11/day11_input_puzzle.txt" >>= print . Y21.day11a
run2021 "11b" = timeIt $ getInput "../data/2021/day11/day11_input_puzzle.txt" >>= print . Y21.day11b
run2021 "12a" = timeIt $ getInput "../data/2021/day12/day12_input_puzzle.txt" >>= print . Y21.day12a
run2021 "12b" = timeIt $ getInput "../data/2021/day12/day12_input_puzzle.txt" >>= print . Y21.day12b
run2021 "13a" = timeIt $ getInput "../data/2021/day13/day13_input_puzzle.txt" >>= print . Y21.day13a
run2021 "13b" = timeIt $ getInput "../data/2021/day13/day13_input_puzzle.txt" >>= print . Y21.day13b
run2021 "14a" = timeIt $ getInput "../data/2021/day14/day14_input_puzzle.txt" >>= print . Y21.day14a
run2021 "14b" = timeIt $ getInput "../data/2021/day14/day14_input_puzzle.txt" >>= print . Y21.day14b
run2021 "15a" = timeIt $ getInput "../data/2021/day15/day15_input_puzzle.txt" >>= print . Y21.day15a
run2021 "15b" = timeIt $ getInput "../data/2021/day15/day15_input_puzzle.txt" >>= print . Y21.day15b
run2021 "16a" = timeIt $ getInput "../data/2021/day16/day16_input_puzzle.txt" >>= print . Y21.day16a
run2021 "16b" = timeIt $ getInput "../data/2021/day16/day16_input_puzzle.txt" >>= print . Y21.day16b
run2021 "17a" = timeIt $ getInput "../data/2021/day17/day17_input_puzzle.txt" >>= print . Y21.day17a
run2021 "17b" = timeIt $ getInput "../data/2021/day17/day17_input_puzzle.txt" >>= print . Y21.day17b
run2021 "18a" = timeIt $ getInput "../data/2021/day18/day18_input_puzzle.txt" >>= print . Y21.day18a
run2021 "18b" = timeIt $ getInput "../data/2021/day18/day18_input_puzzle.txt" >>= print . Y21.day18b
run2021 "19a" = timeIt $ getInput "../data/2021/day19/day19_input_puzzle.txt" >>= print . Y21.day19a
run2021 "19b" = timeIt $ getInput "../data/2021/day19/day19_input_puzzle.txt" >>= print . Y21.day19b
run2021 "20a" = timeIt $ getInput "../data/2021/day20/day20_input_puzzle.txt" >>= print . Y21.day20a
run2021 "20b" = timeIt $ getInput "../data/2021/day20/day20_input_puzzle.txt" >>= print . Y21.day20b
run2021 "21a" = timeIt $ getInput "../data/2021/day21/day21_input_puzzle.txt" >>= print . Y21.day21a
run2021 "21b" = timeIt $ getInput "../data/2021/day21/day21_input_puzzle.txt" >>= print . Y21.day21b
run2021 "22a" = timeIt $ getInput "../data/2021/day22/day22_input_puzzle.txt" >>= print . Y21.day22a
run2021 "22b" = timeIt $ getInput "../data/2021/day22/day22_input_puzzle.txt" >>= print . Y21.day22b
run2021 "23a" = timeIt $ getInput "../data/2021/day23/day23_input_puzzle.txt" >>= print . Y21.day23a
run2021 "23b" = timeIt $ getInput "../data/2021/day23/day23_input_puzzle.txt" >>= print . Y21.day23b
run2021 "24a" = timeIt $ getInput "../data/2021/day24/day24_input_puzzle.txt" >>= print . Y21.day24a
run2021 "24b" = timeIt $ getInput "../data/2021/day24/day24_input_puzzle.txt" >>= print . Y21.day24b
run2021 "25a" = timeIt $ getInput "../data/2021/day25/day25_input_puzzle.txt" >>= print . Y21.day25a
run2021 "25b" = timeIt $ getInput "../data/2021/day25/day25_input_puzzle.txt" >>= print . Y21.day25b
run2021 s = printf "Problem '%s' does not exist.\n" s
