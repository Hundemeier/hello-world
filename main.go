package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Set(color.BgBlue)
	fmt.Println("Hello world. This is a test")
	color.Set(color.BgBlack)
}
