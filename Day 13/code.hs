import Data.List ( elemIndex, sortBy )
import Data.Maybe (fromMaybe)

data Packet = Number Int | List [Packet]

split :: String -> [String]
split s = filter (/="") (splitRecursor "" s)

splitEvery :: Int -> [a] -> [[a]]
splitEvery _ [] = []
splitEvery n list = first : splitEvery n rest
  where
    (first,rest) = splitAt n list

splitRecursor :: String -> String -> [String]
splitRecursor acc "" = [acc]
splitRecursor acc ('[':s) = brkTxt : splitRecursor "" (drop (length brkTxt) ('[':s))
  where brkTxt = getFullBracketText ('[':s)
splitRecursor acc (',':s) = acc : splitRecursor "" s
splitRecursor acc (c: s) = splitRecursor (acc ++ [c]) s

getFullBracketText :: String -> String
getFullBracketText str = recursor str 0 ""

getBracketText :: String -> String
getBracketText "" = ""
getBracketText (c:"") = ""
getBracketText str = tail $ init (getFullBracketText str)

recursor :: String -> Int -> String -> String
recursor [] i res = error "something has gone terribly wrong"
recursor (']' : cs) 1 res = res ++ "]"
recursor (']' : cs) i res = recursor cs (i-1) (res++"]")
recursor ('[' : cs) i res = recursor cs (i+1) (res++"[")
recursor (c : cs) i res = recursor cs i (res ++ [c])

parsePacket :: String -> Packet
parsePacket ('[': xs) = parseList $ getBracketText ('[':xs)
parsePacket s = Number (read s :: Int)

parseList :: String -> Packet
parseList str = List (map parsePacket $ split str)

getOrderNumber :: Packet -> Packet -> Int
getOrderNumber (Number a) (Number b)
  | a == b = 0
  | a < b = 1
  | otherwise = -1
getOrderNumber (Number a) b = getOrderNumber (List [Number a]) b
getOrderNumber a (Number b) = getOrderNumber a (List [Number b])
getOrderNumber (List []) (List []) = 0
getOrderNumber (List []) (List es) = 1
getOrderNumber (List es) (List []) = -1
getOrderNumber (List (a:as)) (List (b:bs))
  | val == 0 = getOrderNumber (List as) (List bs)
  | otherwise = val
  where val = getOrderNumber a b

isCorrectOrder :: Packet -> Packet -> Bool
isCorrectOrder p1 p2
  | orderNum == 1  = True
  | orderNum == -1 = False
  | otherwise      = error "something has gone wrong"
  where orderNum = getOrderNumber p1 p2

main1 :: IO()
main1 = do
  str <- readFile "./input.txt"
  let fileLines = lines str
  let pairs = splitEvery 3 fileLines

  let orderedPairs = map (\(s1:s2:s) -> isCorrectOrder (parsePacket s1) (parsePacket s2)) pairs

  let ans = sum $ map fst $ filter snd $ zip [1..length pairs] orderedPairs

  putStrLn $ "Answer to part 1 is " ++ show ans
  -- Answer to part 1 is 5882

sortFn :: String -> String  -> Ordering
sortFn s1 s2
  | orderNum == 1  = LT
  | orderNum == -1 = GT
  | otherwise      = error "something has gone wrong"
  where orderNum = getOrderNumber (parsePacket s1) (parsePacket s2)

main2 :: IO()
main2 = do
  str <- readFile "./input.txt"
  let fileLines = lines str
  let marker1 = "[[2]]"
  let marker2 = "[[6]]"
  let packets = filter (/="") fileLines ++ [marker1, marker2]

  let sortedPackets = sortBy sortFn packets

  let marker1Index = fromMaybe (-1) (elemIndex marker1 sortedPackets) + 1
  let marker2Index = fromMaybe (-1) (elemIndex marker2 sortedPackets) + 1

  putStrLn $ "Answer to part 2 is " ++ show (marker1Index * marker2Index)