package main

import (
	"../datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
	    matched <= false
	    for i, name := range names{
	        if name == 
	    }
	}
	fmt.Println(lines)
}
