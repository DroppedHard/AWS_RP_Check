package main

import (
	"fmt"
)

func main() {
	fmt.Println("Verifying the test json ")
	var aws AWS_RP_Check
	path := "./data/valid_false1"
	aws.loadFile(path)
}
