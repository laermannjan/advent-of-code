module AdventOfCode.Utils where

import Data.Char (isSpace)

getInput :: FilePath -> IO String
getInput = fmap trim . readFile

trim :: String -> String
trim = trimRight . trimLeft

trimLeft :: String -> String
trimLeft = dropWhile isSpace

trimRight :: String -> String
trimRight = reverse . trimLeft . reverse