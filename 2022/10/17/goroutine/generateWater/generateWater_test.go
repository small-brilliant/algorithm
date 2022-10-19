package generatewater

import (
	"fmt"
	"testing"
)

func TestGenerateWater(t *testing.T) {
	N := 2100093
	ohh, hoh, hho := GenerateWater(N)
	N /= 3
	fmt.Println(float64(ohh) / float64(N))
	fmt.Println(float64(hoh) / float64(N))
	fmt.Println(float64(hho) / float64(N))
}
