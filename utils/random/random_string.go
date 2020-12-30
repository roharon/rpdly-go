package randomutils

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	letters := []rune("abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	if length == 0 {
		return ""
	}

	randArray := make([]rune, length)

	for i := range randArray {
		randArray[i] = letters[rand.Intn(len(letters))]
	}

	return string(randArray)
}
