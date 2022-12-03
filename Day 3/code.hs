import Prelude
import Data.Char ( ord, isUpper )

splitEvery _ [] = []
splitEvery n list = first : splitEvery n rest
  where
    (first,rest) = splitAt n list

priority :: Char -> Int
priority ch = if isUpper ch then ord ch - 64 + 26 else ord ch - 96

getMisplacedItemPriority :: String -> Int
getMisplacedItemPriority s = priority $ head (filter (`elem` secondHalf) firstHalf)
    where firstHalf = take (div (length s) 2) s
          secondHalf = drop (div (length s) 2) s

getCommonItemPriority:: String -> String -> String -> Int
getCommonItemPriority a b c = priority $ head $ filter (\ch -> elem ch b && elem ch c) a

main1 :: IO ()
main1 = do
    str <- readFile "./input.txt"
    let input = lines str

    let ans = sum $ map getMisplacedItemPriority input

    putStrLn $ "Answer to part 1 is " ++ show ans
    -- Answer to part 1 is 7903

main2 :: IO ()
main2 = do
    str <- readFile "./input.txt"
    let input = splitEvery 3 (lines str)

    let ans = sum $ map (\item -> getCommonItemPriority (head item) (item !! 1) (item !! 2)) input


    putStrLn $ "Answer to part 2 is " ++ show ans
    -- Answer to part 2 is 2548


