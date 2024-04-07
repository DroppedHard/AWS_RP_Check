package main

import (
	"fmt"
	"os"
)

type AWS_RP_Check struct {
	data string
}

func errorHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func (aws *AWS_RP_Check) loadFile(path string) {
	bin_data, err := os.ReadFile(path)
	errorHandle(err)
	aws.data = string(bin_data)
	fmt.Println(aws.data)
}
