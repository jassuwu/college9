isPali :: Eq a => [a] -> Bool
isPali x = x == reverse x
