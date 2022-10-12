package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		data := strings.Split(in.Text(), " ")
		sort.Strings(data)
		fmt.Println(strings.Join(data, ","))
	}
}
