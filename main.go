package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Set(color.BgBlue)
	fmt.Println("Hello world. This is a test")
	fmt.Println("the next test")
	color.Set(color.BgBlack)
}
