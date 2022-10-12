package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	for i := 0; i < n; i++ {
		in.Scan()
		data := strings.Split(in.Text(), " ")
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		fmt.Println(strings.Join(data, " "))
	}
}
