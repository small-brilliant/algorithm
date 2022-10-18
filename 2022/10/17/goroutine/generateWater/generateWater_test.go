package generatewater

import (
	"fmt"
	"testing"
)

func TestGenerateWater(t *testing.T) {
	ohh, hoh, hho := GenerateWater(210009)
	fmt.Println(ohh, hoh, hho)
}
