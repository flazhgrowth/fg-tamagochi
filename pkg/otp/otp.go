package otp

import "github.com/flazhgrowth/fg-tamagopkg/random"

func GenerateNumbers(length int) uint64 {
	return random.GenerateRandomNumber(length)
}

func GenerateStrings(length int) string {
	return random.GenerateRandomString(length)
}
