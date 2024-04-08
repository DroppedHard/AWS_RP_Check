package main

import (
	"testing"
)

// Helpers

var invalid_path = "./data/invalid/"
var false_path = "./data/false/"
var true_path = "./data/true/"

func helperInvalid(t *testing.T, file string, expectedError string, failText string) {
	// given
	var aws AWS_RP_Check
	// when
	err := aws.loadFile(invalid_path + file)
	// then
	if err == nil {
		t.Error("Error expected, but got nil")
	} else if err.Error() != expectedError {
		t.Error(failText)
	}
}

func helperValid(t *testing.T, path string, expectedOut bool) {
	// given
	var aws AWS_RP_Check
	// when
	err := aws.loadFile(path)
	if err != nil {
		t.Error("Unexpected error during loading file: " + err.Error())
	}
	// then
	out, err2 := aws.verifyResource()
	if err2 != nil {
		t.Error("Unexpected error during resource verification: " + err.Error())
	} else if out != expectedOut {
		t.Errorf("Incorrect output: expected: %v received: %v", expectedOut, out)
	}
}

// Tests for invalid json
func TestInvalidJsonParse(t *testing.T) {
	helperInvalid(t, "invalid_json_parse.txt", "invalid character 'a' looking for beginning of value", "JSON invalid character as beginning of value expected")
}
func TestInvalidActionNotAction1(t *testing.T) {
	helperInvalid(t, "invalid_action_notaction1.json", "invalid Action or NotAction definition", "invalid Action or NotAction definition expected")
}
func TestInvalidActionNotAction2(t *testing.T) {
	helperInvalid(t, "invalid_action_notaction2.json", "invalid Action or NotAction definition", "invalid Action or NotAction definition expected")
}
func TestInvalidNameRegex(t *testing.T) {
	helperInvalid(t, "invalid_name_regex.json", "PolicyName field does not match RegEx pattern", "PolicyName field does not match RegEx pattern expected")
}
func TestInvalidPolicyVersion1(t *testing.T) {
	helperInvalid(t, "invalid_policy_version1.json", "invalid PolicyDocument Version", "invalid PolicyDocument Version expected")
}
func TestInvalidPolicyVersion2(t *testing.T) {
	helperInvalid(t, "invalid_policy_version2.json", "invalid PolicyDocument Version", "invalid PolicyDocument Version expected")
}
func TestInvalidStatementEffect1(t *testing.T) {
	helperInvalid(t, "invalid_statement_effect1.json", "invalid Effect in Statements", "invalid Effect in Statements expected")
}
func TestInvalidStatementEffect2(t *testing.T) {
	helperInvalid(t, "invalid_statement_effect2.json", "invalid Effect in Statements", "invalid Effect in Statements expected")
}
func TestUndefinedPolicyName(t *testing.T) {
	helperInvalid(t, "undefined_policy_name.json", "invalid length of PolicyName field, or not defined", "invalid length of PolicyName field, or not defined expected")
}
func TestUndefinedResource1(t *testing.T) {
	helperInvalid(t, "undefined_resource1.json", "resource field not defined", "resource field not defined expected")
}
func TestUndefinedResource2(t *testing.T) {
	helperInvalid(t, "undefined_resource2.json", "resource field not defined", "resource field not defined expected")
}
func TestUndefinedStatements(t *testing.T) {
	helperInvalid(t, "undefined_statements.json", "undefined Statements", "undefined Statements expected")
}

// Tests for false output

func TestFalse1Statement1Resource1(t *testing.T) {
	helperValid(t, false_path+"1stat_1res_1.json", false)
}
func TestFalse1Statement1Resource2(t *testing.T) {
	helperValid(t, false_path+"1stat_1res_2.json", false)
}
func TestFalseManyStatements1Resource1(t *testing.T) {
	helperValid(t, false_path+"many_stat_1res_1.json", false)
}
func TestFalseManyStatements1Resource2(t *testing.T) {
	helperValid(t, false_path+"many_stat_1res_2.json", false)
}
func TestFalseManyStatements1Resource3(t *testing.T) {
	helperValid(t, false_path+"many_stat_1res_3.json", false)
}
func TestFalseManyStatementsManyResource1(t *testing.T) {
	helperValid(t, false_path+"many_stat_many_res_1.json", false)
}
func TestFalseManyStatementsManyResource2(t *testing.T) {
	helperValid(t, false_path+"many_stat_many_res_2.json", false)
}

// Tests for true output

func TestTrue1Statement1Resource1(t *testing.T) {
	helperValid(t, true_path+"1stat_1res_1.json", true)
}
func TestTrue1Statement1Resource2(t *testing.T) {
	helperValid(t, true_path+"1stat_1res_2.json", true)
}
func TestTrueManyStatement1Resource1(t *testing.T) {
	helperValid(t, true_path+"many_stat_1res_1.json", true)
}
func TestTrueManyStatement1Resource2(t *testing.T) {
	helperValid(t, true_path+"many_stat_1res_2.json", true)
}
func TestTrueManyStatementManyResource1(t *testing.T) {
	helperValid(t, true_path+"many_stat_many_res_1.json", true)
}
func TestTrueManyStatementManyResource2(t *testing.T) {
	helperValid(t, true_path+"many_stat_many_res_2.json", true)
}
