package main

func process() {
	// defer func() {
	// 	recover()
	// }()
	panic("出现错误")
}
func main() {
	process()
}
