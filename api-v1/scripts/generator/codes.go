package generator

import "math/rand"

func RandomLetters(iterations int) string {
	validLetters := "abcdefghjkmnopqrstuvwxyz"
	result := ""
	for range iterations {
		index := rand.Intn(len(validLetters))
		result += string(validLetters[index])
	}
	return result
}

func RandomNumbers(iterations int) string {
	validNumbers := "0123456789"
	result := ""
	for range iterations {
		index := rand.Intn(len(validNumbers))
		result += string(validNumbers[index])
	}
	return result
}
