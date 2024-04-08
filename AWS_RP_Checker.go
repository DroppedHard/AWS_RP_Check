package main

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strings"
)

type AwsIamStatement struct { // SID field is optional
	Effect    string      `json:"Effect"`    // required https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_effect.html
	Action    interface{} `json:"Action"`    // contains string or []string
	NotAction interface{} `json:"NotAction"` // contains string or []string
	Resource  interface{} `json:"Resource"`  // contains string or []string
}

type AwsIamPolicyDocument struct {
	Version   string            `json:"Version"` // required https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
	Statement []AwsIamStatement `json:"Statement"`
}

type AwsIamRolePolicy struct {
	PolicyDocument AwsIamPolicyDocument `json:"PolicyDocument"`
	PolicyName     string               `json:"PolicyName"`
}

type AwsRolePolicyChecker struct {
	JsonData AwsIamRolePolicy
}

func isInArray(target string, array []string) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}
	return false
}

func (aws *AwsRolePolicyChecker) verifyAwsIamRolePolicyFormat() error {
	// verify PolicyName
	if 1 > len(aws.JsonData.PolicyName) || len(aws.JsonData.PolicyName) > 128 {
		return errors.New("invalid length of PolicyName field, or not defined")
	}
	if match, _ := regexp.MatchString("[\\w+=,.@-]+", aws.JsonData.PolicyName); !match {
		return errors.New("PolicyName field does not match RegEx pattern")
	}
	// verify PolicyDocument.Version
	if !isInArray(aws.JsonData.PolicyDocument.Version, []string{"2012-10-17", "2008-10-17"}) {
		return errors.New("invalid PolicyDocument Version")
	}
	// verify PolicyDocument.Statement
	if len(aws.JsonData.PolicyDocument.Statement) == 0 {
		return errors.New("undefined Statements")
	}
	for _, statement := range aws.JsonData.PolicyDocument.Statement {
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

func (aws *AwsRolePolicyChecker) loadFile(path string) (err error) {
	bin_data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bin_data, &aws.JsonData); err != nil {
		return
	}
	if err = aws.verifyAwsIamRolePolicyFormat(); err != nil {
		return
	}
	return nil
}

func (aws *AwsRolePolicyChecker) verifyResource() (bool, error) {
	for _, statement := range aws.JsonData.PolicyDocument.Statement {
		switch resource_field := statement.Resource.(type) {
		case string:
			if strings.Contains(resource_field, "*") {
				return false, nil
			}
		case []interface{}:
			for _, resource := range resource_field {
				if resource, ok := resource.(string); ok {
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
