package utils

import "math/rand"

func RandomArrString(maxLen int) string {
	len := int(rand.Float64() * float64(maxLen))
	b := make([]byte, len)
	t := 0
	for i := 0; i < len; i++ {
		t = rand.Intn(24)
		b[i] = byte(t + 'a')
	}
	return string(b)
}
