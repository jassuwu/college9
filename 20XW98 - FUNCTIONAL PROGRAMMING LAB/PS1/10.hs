init1 :: [a] -> [a]
init1 xs = reverse $ tail $ reverse $ xs

init2 :: [a] -> [a]
init2 [_] = []
init2 (x:xs) = x : init2 xs
init2 [] = error "Empty list"
