import unittest

def get_discount_rate(price: int) -> int:
    if 1 <= price <= 50:
        return 0
    elif 50 < price <= 200:
        return 5
    elif 201 <= price <= 500:
        return 10
    elif price >= 501:
        return 15
    else:
        return -1

# tests w/ BVA
class TestGetDiscountRate(unittest.TestCase):
    def test_less_than_1(self):
        got, want = get_discount_rate(0), -1
        self.assertEqual(got, want, f"got={got}, want={want}")
        
    def test_equal_to_1(self):
        got, want = get_discount_rate(1), 0
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_greater_than_1(self):
        got, want = get_discount_rate(2), 0
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_less_than_50(self):
        got, want = get_discount_rate(49), 0
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_equal_to_50(self):
        got, want = get_discount_rate(50), 0
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_greater_than_50(self):
        got, want = get_discount_rate(51), 5
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_less_than_200(self):
        got, want = get_discount_rate(199), 5
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_equal_to_200(self):
        got, want = get_discount_rate(200), 5
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_greater_than_200(self):
        got, want = get_discount_rate(201), 10
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_less_than_500(self):
        got, want = get_discount_rate(499), 10
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_equal_to_500(self):
        got, want = get_discount_rate(500), 10
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_greater_than_500(self):
        got, want = get_discount_rate(501), 15
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_greater_than_501(self):
        got, want = get_discount_rate(502), 15
        self.assertEqual(got, want, f"got={got}, want={want}")

if __name__ == "__main__":
    unittest.main()