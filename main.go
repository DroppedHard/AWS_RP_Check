package main

import (
	"fmt"
	"os"
)

func errorHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		aws  AwsRolePolicyChecker
		err  error
		out  bool
		path string
	)

	if len(os.Args) > 1 && os.Args[1] != "" {
		path = os.Args[1]
	} else {
		path = "./source.json"
	}
	err = aws.loadFile(path)
	errorHandle(err)

	out, err = aws.verifyResource()
	errorHandle(err)

	fmt.Println("Outcome: ", out)
}
