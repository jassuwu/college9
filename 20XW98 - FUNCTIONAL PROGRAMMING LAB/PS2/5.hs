reverseList :: [a] -> [a]
reverseList = foldl (flip (:)) []
