import unittest

def determine_grade(mark1: int, mark2: int, mark3: int) -> str:
    average = (mark1 + mark2 + mark3) / 3
    if 90 <= average <= 100:
        return "First Class Distinction"
    elif 75 <= average < 90:
        return "First Class"
    elif 60 <= average < 75:
        return "Second Class"
    elif 50 <= average < 60:
        return "Third Class"
    else:
        return "Fail"

# Tests with Boundary Value Analysis
class TestDetermineGrade(unittest.TestCase):
    def test_below_0(self):
        got, want = determine_grade(-1, -1, -1), "Fail"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_exactly_0(self):
        got, want = determine_grade(0, 0, 0), "Fail"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_above_0(self):
        got, want = determine_grade(1, 1, 1), "Fail"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_below_50(self):
        got, want = determine_grade(49, 49, 49), "Fail"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_exactly_50(self):
        got, want = determine_grade(50, 50, 50), "Third Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_above_50(self):
        got, want = determine_grade(51, 51, 51), "Third Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_below_60(self):
        got, want = determine_grade(55, 55, 55), "Third Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_exactly_60(self):
        got, want = determine_grade(60, 60, 60), "Second Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_above_60(self):
        got, want = determine_grade(61, 61, 61), "Second Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_below_75(self):
        got, want = determine_grade(70, 70, 70), "Second Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_exactly_75(self):
        got, want = determine_grade(75, 75, 75), "First Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_above_75(self):
        got, want = determine_grade(76, 76, 76), "First Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_below_90(self):
        got, want = determine_grade(85, 85, 85), "First Class"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_exactly_90(self):
        got, want = determine_grade(90, 90, 90), "First Class Distinction"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_above_90(self):
        got, want = determine_grade(91, 91, 91), "First Class Distinction"
        self.assertEqual(got, want, f"got={got}, want={want}")

    def test_above_100(self):
        got, want = determine_grade(101, 101, 101), "First Class Distinction"
        self.assertEqual(got, want, f"got={got}, want={want}")

if __name__ == "__main__":
    unittest.main(verbosity=2)
