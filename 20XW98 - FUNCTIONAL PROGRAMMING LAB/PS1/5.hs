-- vanilla quicksort algorithm implementation
quicksort :: Ord a => [a] -> [a]
quicksort [] = []
quicksort (x:xs) =
    let smallerSorted = quicksort [a | a <- xs, a < x]
        biggerSorted = quicksort [a | a <- xs, a > x]
    in smallerSorted ++ [x] ++ biggerSorted

main = do
     let unsortedList = [2, 2, 3, 1, 1]
     putStrLn $ "Unsorted list" ++ show unsortedList
     let sortedList = quicksort unsortedList
     putStrLn $ "Sorted list" ++ show sortedList --[1, 2, 3]: Changing the <= to <, ignored the duplicates
	 
