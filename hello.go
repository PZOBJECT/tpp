package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s! 这是 Hermes + Claude Code 用 Go 写的代码。", name)
}

func main() {
	fmt.Println(greet("World"))
}
