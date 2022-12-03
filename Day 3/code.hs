import Prelude
import Data.Char ( ord, isUpper )

priority :: Char -> Int
priority ch = if isUpper ch then ord ch - 64 + 26 else ord ch - 96

getPriority :: String -> Int
getPriority s = priority $ head (filter (`elem` secondHalf) firstHalf)
    where firstHalf = take (div (length s) 2) s
          secondHalf = drop (div (length s) 2) s

main1 :: IO ()
main1 = do
    str <- readFile "./input.txt"
    let input = lines str

    let ans = sum $ map getPriority input

    putStrLn $ "Answer to part 1 is " ++ show ans


