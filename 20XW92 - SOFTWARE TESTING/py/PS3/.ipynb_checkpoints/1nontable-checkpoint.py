import unittest

def passwordValidator(password):
    length = len(password)
    return 6 <= length <= 12

class TestPasswordValidator(unittest.TestCase):
    def test_below_boundary(self):
        password = "abcde"
        expected = False
        actual = passwordValidator(password)
        self.assertEqual(expected, actual, f"got {actual} want {expected}")

    def test_equalto_lower_boundary(self):
        password = "abcdef"
        expected = True
        actual = passwordValidator(password)
        self.assertEqual(expected, actual, f"got {actual} want {expected}")

    def test_within_boundary(self):
        password = "abcdefghi"
        expected = True
        actual = passwordValidator(password)
        self.assertEqual(expected, actual, f"got {actual} want {expected}")

    def test_equalto_upper_boundary(self):
        password = "abcdefghijkl"
        expected = True
        actual = passwordValidator(password)
        self.assertEqual(expected, actual, f"got {actual} want {expected}")

    def test_above_boundary(self):
        password = "abcdefghijklmnop"
        expected = True
        actual = passwordValidator(password)
        self.assertEqual(expected, actual, f"got {actual} want {expected}")

if __name__ == '__main__':
    unittest.main(verbosity=10)
