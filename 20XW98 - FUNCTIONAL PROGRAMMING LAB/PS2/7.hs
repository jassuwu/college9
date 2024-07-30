-- Since the question is not practical Haskell, we've to create our own nestedList datatype.
-- The question might from lisp or clojure idk.
data NestedList a = Elem a | List [NestedList a] deriving (Show)

nestedList :: NestedList Char
nestedList = List [Elem 'a', List [Elem 'b', List [Elem 'c', Elem 'd'], Elem 'e']]

flatten :: NestedList a -> [a]
flatten (Elem x) = [x]
flatten (List xs) = concatMap flatten xs

main :: IO ()
main = do
  let flattened = flatten nestedList
  print flattened  -- Outputs: "abcde"

