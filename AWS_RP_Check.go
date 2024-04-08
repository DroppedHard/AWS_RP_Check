package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type AWS_IAM_Statement struct { // SID field is optional
	Effect    string      `json:"Effect"`    // required https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_effect.html
	Action    interface{} `json:"Action"`    // contains string or []string
	NotAction interface{} `json:"NotAction"` // contains string or []string
	Resource  interface{} `json:"Resource"`  // contains string or []string
}

type AWS_IAM_PolicyDocument struct {
	Version   string              `json:"Version"` // required https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
	Statement []AWS_IAM_Statement `json:"Statement"`
}

type AWS_IAM_Role struct {
	PolicyDocument AWS_IAM_PolicyDocument `json:"PolicyDocument"`
	PolicyName     string                 `json:"PolicyName"`
}

type AWS_RP_Check struct {
	jsonData AWS_IAM_Role
}

func isInArray(target string, array []string) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}
	return false
}

func (aws *AWS_RP_Check) verifyAWS_IAM_RP() error {
	// verify PolicyName
	if 1 >= len(aws.jsonData.PolicyName) || len(aws.jsonData.PolicyName) >= 128 {
		return errors.New("invalid length of PolicyName field")
	}
	match, _ := regexp.MatchString("[\\w+=,.@-]+", aws.jsonData.PolicyName)
	if !match {
		return errors.New("PolicyName field does not match RegEx pattern")
	}
	// verify PolicyDocument.Version
	if !isInArray(aws.jsonData.PolicyDocument.Version, []string{"2012-10-17", "2008-10-17"}) {
		return errors.New("invalid PolicyDocument Version")
	}
	// verify PolicyDocument.Statement
	for _, statement := range aws.jsonData.PolicyDocument.Statement {
		// verify Effect
		if !isInArray(statement.Effect, []string{"Allow", "Deny"}) {
			return errors.New("invalid Effect in Statements")
		}
		// verify "XOR" on Action and NotAction field (either one must be defined)
		if (statement.Action == nil) == (statement.NotAction == nil) {
			return errors.New("invalid Action or NotAction definition")
		}
		// verify if Resources are defined
		if statement.Resource == nil {
			return errors.New("resource field not defined")
		}
	}

	return nil
}

func (aws *AWS_RP_Check) loadFile(path string) error {
	bin_data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(string(bin_data))
	err = json.Unmarshal(bin_data, &aws.jsonData)
	if err != nil {
		return err
	}
	err = aws.verifyAWS_IAM_RP()
	if err != nil {
		return err
	}
	return nil
}

func (aws *AWS_RP_Check) verifyResource() (bool, error) {
	for _, statement := range aws.jsonData.PolicyDocument.Statement {
		switch resource_field := statement.Resource.(type) {
		case string:
			if strings.Contains(resource_field, "*") {
				return false, nil
			}
		case []interface{}: // resource can be either string, or slice of strings (as documentation says)
			for _, resource := range resource_field {
				// using type assertions
				resource, ok := resource.(string)
				if ok {
					if strings.Contains(resource, "*") {
						return false, nil
					}
				} else {
					return true, errors.New("invalid resource field type")
				}
			}
		default:
			return true, errors.New("invalid resource field type")
		}
	}
	return true, nil
}
