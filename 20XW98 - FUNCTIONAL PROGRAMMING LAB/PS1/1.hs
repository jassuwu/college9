-- Define the double function
double :: Int -> Int
double x = x + x

result1 :: Int
result1 = double (double 2)  -- 8

main :: IO ()
main = do
    putStrLn $ "Result of double (double 2) is: " ++ show result1
