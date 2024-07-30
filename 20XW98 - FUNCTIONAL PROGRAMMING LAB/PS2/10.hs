pack :: Eq a => [a] -> [[a]]
pack [] = []
pack (x:xs) = let (first, rest) = span (==x) xs
              in (x:first) : pack rest

rle :: Eq a => [a] -> [(Int, a)]
rle = map (\x -> (length x, head x)) . pack

