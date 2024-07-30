-- use of list comprehension
squeeze :: [Char] -> [Char]
squeeze l = [x | (x, y) <- zip l $ tail l ++ " ", x /= y]
