package main

import (
	"fmt"
	"strings"
)

func main() {
	var broken string = "G# R#cks !"
	var replacer *strings.Replacer = strings.NewReplacer("#", "o")
	var fixed string = replacer.Replace(broken)
	fmt.Println(fixed)
}
