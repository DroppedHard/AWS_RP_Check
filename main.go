package main

import (
	"fmt"
)

func main() {
	fmt.Println("Verifying the test json ")
	var aws AWS_RP_Check
	path := "./data/valid_true4.json"
	// path := "./data/invalid13.json"
	aws.loadFile(path)
	fmt.Println("Outcome: ", aws.verifyResource())
}
