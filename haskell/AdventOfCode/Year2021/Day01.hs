module AdventOfCode.Year2021.Day01 (day01a, day01b) where

import Data.List (tails)

strToIntList :: String -> [Int]
strToIntList string = map (read :: String -> Int) (lines string)

hasIncreased :: [Int] -> [Bool]
hasIncreased lst = zipWith (<) lst (tail lst)

numTrue :: [Bool] -> Int
numTrue lst = length (filter (== True) lst)

movingSum k lst = map sumup $ tails lst
  where
    sumup = sum . take k

day01a :: String -> Int
day01a = numTrue . hasIncreased . strToIntList

day01b :: String -> Int
day01b = numTrue . hasIncreased . movingSum 3 . strToIntList