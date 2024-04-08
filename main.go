package main

import (
	"fmt"
)

func errorHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var aws AWS_RP_Check
	// path := "./data/valid_true4.json"
	path := "./data/invalid2.json"
	err := aws.loadFile(path)
	fmt.Printf("err: %T\n", err)
	errorHandle(err)
	out, err2 := aws.verifyResource()
	errorHandle(err2)
	fmt.Println("Outcome: ", out)
}
