package main

import (
	"fmt"
)

func Hello(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(Hello("Yadomin"))
}
