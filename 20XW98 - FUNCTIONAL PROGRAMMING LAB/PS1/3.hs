-- Define the product function
product :: [Int] -> Int
product [] = 1
product (x:xs) = x * product xs

result3 :: Int
result3 = product [2, 3, 4] -- 24

main = do
     putStrLn $ "Product of [2, 3, 4] is: " ++ show result3
