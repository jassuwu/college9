-- Calculate sum [x] = x for any x
result2 :: Int
result2 = sum [5]  -- 5

main :: IO ()
main = do
    putStrLn $ "Sum of [5] is: " ++ show result2
