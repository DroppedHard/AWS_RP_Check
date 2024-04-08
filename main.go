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
	path := "./data/false/1stat_1res_1.json"
	err := aws.loadFile(path)
	fmt.Printf("err: %T\n", err)
	errorHandle(err)
	out, err2 := aws.verifyResource()
	errorHandle(err2)
	fmt.Println("Outcome: ", out)
}
