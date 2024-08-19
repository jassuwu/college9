package two

import (
	"testing"
)

var testCases = []struct {
	Name     string
	Fname    string
	Address  string
	Email    string
	IsValid  bool
	Expected string
}{
	{
		Name:     "All valid input",
		Fname:    "John",
		Address:  "123 Street",
		Email:    "john.doe@example.com",
		IsValid:  true,
		Expected: "First name: John \t Address: 123 Street \t Email: john.doe@example.com",
	},
	{
		Name:     "Invalid first name",
		Fname:    "Jon",
		Address:  "123 Street",
		Email:    "john.doe@example.com",
		IsValid:  false,
		Expected: "Invalid first name\n",
	},
	{
		Name:     "Invalid address length",
		Fname:    "John",
		Address:  "123",
		Email:    "john.doe@example.com",
		IsValid:  false,
		Expected: "Invalid address length\n",
	},
	{
		Name:     "Invalid email length",
		Fname:    "John",
		Address:  "123 Street",
		Email:    "jon@d.c",
		IsValid:  false,
		Expected: "Invalid Email length\n",
	},
	{
		Name:     "Email missing @",
		Fname:    "John",
		Address:  "123 Street",
		Email:    "john.doeexample.com",
		IsValid:  false,
		Expected: "Email must contain . and @ characters\n",
	},
	{
		Name:     "Email missing .",
		Fname:    "John",
		Address:  "123 Street",
		Email:    "johndoe@examplecom",
		IsValid:  false,
		Expected: "Email must contain . and @ characters\n",
	},
	{
		Name:     "Invalid first name and email",
		Fname:    "Jon",
		Address:  "123 Street",
		Email:    "jon@d.c",
		IsValid:  false,
		Expected: "Invalid first name\nInvalid Email length\n",
	},
}

func TestValidateInput(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			gotValid, gotOutput := ValidateInput(tc.Fname, tc.Address, tc.Email)
			if gotValid != tc.IsValid || gotOutput != tc.Expected {
				t.Errorf(
					"ValidateInput(%q, %q, %q) = (%v, %q); want (%v, %q)",
					tc.Fname,
					tc.Address,
					tc.Email,
					gotValid,
					gotOutput,
					tc.IsValid,
					tc.Expected,
				)
			}
		})
	}
}
