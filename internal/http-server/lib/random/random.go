package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	charts := []rune("ABSDEFGHIJKLMOPQRSTUVWXYZ" +
		"absdefghijklmopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, size)
	for i := range b {
		b[i] = charts[rnd.Intn(len(charts))]
	}

	return string(b)
}
