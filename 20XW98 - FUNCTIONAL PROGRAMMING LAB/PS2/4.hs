numOfElems :: [a] -> Int
numOfElems = foldr (\x -> (+) 1) 0
