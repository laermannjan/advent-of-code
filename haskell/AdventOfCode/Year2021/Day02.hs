{-# LANGUAGE ViewPatterns #-}

module AdventOfCode.Year2021.Day02 (day02a, day02b) where

import Data.List (stripPrefix)

parseInst :: String -> (Int, Int)
parseInst (stripPrefix "forward" -> Just x) = (read x, 0)
parseInst (stripPrefix "down" -> Just x) = (0, read x)
parseInst (stripPrefix "up" -> Just x) = (0, negate $ read x)
parseInst x = (0, 0)

sumCoords :: [(Int, Int)] -> (Int, Int)
sumCoords = foldl (\(a1, b1) (a2, b2) -> (a1 + a2, b1 + b2)) (0, 0)

sumCoordsAim :: [(Int, Int)] -> (Int, Int, Int)
sumCoordsAim = foldl accumulator (0, 0, 0)
  where
    accumulator (horizontal, depth, aim) =
      \(h, a) -> (horizontal + h, depth + (h * (aim + a)), aim + a)

day02a :: String -> Int
day02a = convert . sumCoords . map parseInst . lines
  where
    convert = uncurry (*)

day02b :: String -> Int
day02b = convert . sumCoordsAim . map parseInst . lines
  where
    convert = \(h, d, a) -> h * d