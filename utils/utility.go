package utility

import (
	"fmt"
	"math"
)

func IsArmstrong(num int) bool {
	// Convert the number to a string to easily count the digits
	digits := fmt.Sprintf("%d", num)
	numDigits := len(digits)

	// Calculate the sum of digits raised to the power of the number of digits
	sum := 0
	for _, digit := range digits {
		digitValue := int(digit - '0') // Convert the character back to an integer
		sum += int(math.Pow(float64(digitValue), float64(numDigits)))
	}

	// Check if the sum is equal to the original number
	return sum == num
}


func DigitSum(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}


func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPerfect(n int) bool {
	sum := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum == n
}