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
	var (
		aws AwsRolePolicyChecker
		err error
		out bool
	)

	path := "./source.json"
	err = aws.loadFile(path)
	errorHandle(err)

	out, err = aws.verifyResource()
	errorHandle(err)

	fmt.Println("Outcome: ", out)
}
