package challenge_2

import (
	"unicode"
)

func isValidPassword(password string) bool {
	countOfUpper := 0
	countOfDigit := 0
	if len(password) < 5 {
		return false
	}

	if len(password) > 12 {
		return false
	}

	for _, char := range password {
		if unicode.IsUpper(char) {
			countOfUpper += 1
		}
		if unicode.IsDigit(char) {
			countOfDigit += 1
		}
	}

	if countOfUpper > 0 && countOfDigit > 0 {
		return true
	}

	return false

}

/*
Password Strength
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one digit.
A string is really just a read-only slice of bytes.
This means that you can use the same techniques you learned in the previous lesson to iterate over the characters of a string.

Assignment
Implement the isValidPassword function by looping through each character in the password string.
Make sure the password is long enough and includes at least one uppercase letter and one digit.


their solution:

func isValidPassword(password string) bool {
	length := len(password)
	hasUpper := false
	hasNumber := false

	for i := 0; i < length; i++ {
		char := password[i]
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}

	return length >= 5 && length <= 12 && hasUpper && hasNumber
}


*/
