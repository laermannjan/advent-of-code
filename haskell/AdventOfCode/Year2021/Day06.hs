module AdventOfCode.Year2021.Day06 (day06a, day06b) where

import Data.Text (pack, split, unpack)

type Freqs = (Int, Int, Int, Int, Int, Int, Int, Int, Int)

step :: Freqs -> Freqs
step (a, b, c, d, e, f, g, h, i) = (b, c, d, e, f, g, h + a, i, a)

freqs :: [Int] -> Freqs
freqs timings = cast $ map (\x -> length $ filter (== x) timings) [0 .. 8]
  where
    cast [a, b, c, d, e, f, g, h, i] = (a, b, c, d, e, f, g, h, i)
    cast _ = undefined

loopN :: (a -> a) -> Int -> a -> a
loopN func 0 inp = undefined
loopN func 1 inp = func inp
loopN func n inp = loopN func (n -1) (func inp)

parse :: String -> Freqs
parse input = freqs $ map ((read :: String -> Int) . unpack) $ split (== ',') $ pack input

countFish :: Freqs -> Int
countFish (a, b, c, d, e, f, g, h, i) = a + b + c + d + e + f + g + h + i

day06a :: String -> Int
day06a input = countFish . loopN step 80 $ parse input

day06b :: String -> Int
day06b input = countFish . loopN step 256 $ parse input