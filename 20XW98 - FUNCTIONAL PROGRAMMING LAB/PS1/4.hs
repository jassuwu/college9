-- vanilla quicksort algorithm implementation
quicksort :: Ord a => [a] -> [a]
quicksort [] = []
quicksort (x:xs) =
    let smallerSorted = quicksort [a | a <- xs, a <= x]
        biggerSorted = quicksort [a | a <- xs, a > x]
    in smallerSorted ++ [x] ++ biggerSorted

-- modified quicksort algorithm implementation for array reversal
rqsort :: Ord a => [a] -> [a]
rqsort [] = []
rqsort (x:xs) =
    let smallerSorted = rqsort [a | a <- xs, a > x]
        biggerSorted = rqsort [a | a <- xs, a <= x]
    in smallerSorted ++ [x] ++ biggerSorted

main :: IO ()
main = do
    let unsortedList = [2, 2, 3, 1, 1]
    putStrLn $ "Unsorted list: " ++ show unsortedList
    let sortedList = quicksort unsortedList
    putStrLn $ "Sorted list: " ++ show sortedList
    let rsortedList = rqsort unsortedList
    putStrLn $ "Reversed Sorted list: " ++ show rsortedList 
