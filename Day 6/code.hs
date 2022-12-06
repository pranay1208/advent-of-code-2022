import Data.List

isUniqueList :: [Char] -> Bool 
isUniqueList list = length (nub list) == length list 

findStartMarker :: String -> Int
findStartMarker [] = 3
findStartMarker (x:xs) = 1 + findStartMarker (if isUniqueList (x:take 3 xs) then [] else xs)

findMessageMarker :: String -> Int
findMessageMarker [] = 13
findMessageMarker (x:xs) = 1 + findMessageMarker (if isUniqueList (x:take 13 xs) then [] else xs)

main :: IO ()
main = do
    str <- readFile "./input.txt"

    putStrLn $ "Answer to part 1 is " ++ show (findStartMarker str)
    -- Answer to part 1 is 1287

    putStrLn $ "Answer to part 2 is " ++ show (findMessageMarker str)
    -- Answer to part 2 is 3716