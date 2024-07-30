three :: Int -> Int
three = (!!) [1..]

main = do
  print $ three 2
