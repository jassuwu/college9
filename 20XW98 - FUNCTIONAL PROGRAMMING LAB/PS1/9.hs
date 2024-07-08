lastlast :: [a] -> a
lastlast [x] = x
lastlast (_:xs) = lastlast xs
lastlast [] = error "Empty list"
