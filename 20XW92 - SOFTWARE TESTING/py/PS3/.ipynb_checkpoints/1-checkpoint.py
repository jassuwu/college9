import unittest
from dataclasses import dataclass

def password_validator(password: str) -> bool:
    length = len(password)
    return 6 <= length <= 12

class TestPasswordValidator(unittest.TestCase):
    def test_password_validator(self):
        @dataclass
        class TestCase:
            name: str
            input: str
            want: bool

        testcases = [
            TestCase(name="len < 6", input="pass", want=False),
            TestCase(name="len == 6", input="passwd", want=True),
            TestCase(name="7 <= len <= 11", input="password", want=True),
            TestCase(name="len == 12", input="passwordishy", want=True),
            TestCase(name="len > 12", input="passwordyword", want=False),
        ]

        for case in testcases:
            got = password_validator(case.input)
            self.assertEqual(
                case.want,
                got,
                f"failed test {case.name} want {case.want} got {got}",
            )

if __name__ == "__main__":
    unittest.main()
