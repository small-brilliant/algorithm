package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		sum := 0
		data := strings.Split(in.Text(), " ")
		for _, v := range data {
			num, _ := strconv.Atoi(v)
			sum += num
		}
		fmt.Println(sum)
	}
}
