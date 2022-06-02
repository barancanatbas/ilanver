package helpers

import "math/rand"

func RandNumber(start int, end int) int {
	code := rand.Intn(end-start) + start
	return code
}
