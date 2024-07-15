lastButOne :: [a] -> Maybe a
lastButOne [] = Nothing
lastButOne [_] = Nothing
lastButOne [x,_] = Just x
lastButOne (_:xs) = lastButOne xs


main = do
     let arr = [1, 2, 3]
     putStrLn $ "The original array: " ++ show arr
     let secondToLast = lastButOne arr
     putStrLn $ "Last but one of the array: " ++ show secondToLast
