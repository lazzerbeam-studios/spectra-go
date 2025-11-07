package generator

import "math/rand"

func GenerateRandomLetters() string {
	validLetters := "abcdefghjkmnopqrstuvwxyz"
	result := ""
	for range 3 {
		index := rand.Intn(len(validLetters))
		result += string(validLetters[index])
	}
	return result
}
