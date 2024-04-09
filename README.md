# AWS_RP_Check
## Description
Go program for verifying AWS::IAM::Role Policy as the Remitly Internship Homework. The task was to write a method, which verifies input JSON data (of AWS::IAM::Role Policy type) and returns a boolean value:
- `false` if an input JSON *Resource* field contains a single asterisk
- `true` in any other case
Disclaimer: If JSON data is incorrect (not a JSON AWS::IAM::Role Policy data) only error is thrown due to invalid input data.

## How to start?
### Requirements
- [Go Compiler](https://go.dev/dl/) for your Operating System to run and build the module (remember to add go binary to PATH - we'll need `go` command)
- [Git](https://git-scm.com/downloads) to clone this repository

### Running
1. Clone this repository to your desired folder
```cmd
cd your/desired/folder
git clone https://github.com/DroppedHard/AWS_RP_Check.git
```
2. Navigate to freshly cloned repository
```cmd
cd AWS_RP_Check
```
3. Run the module
```cmd
go run main.go AWS_RP_Checker.go [JSON data path]
```
- With no parameters module loads sample data given in [source.json](https://github.com/DroppedHard/AWS_RP_Check/blob/main/source.json) file.
- Module accepts relative and absolute path to JSON data file.
#### Optional
You can build the project with `go build` command, which will create a binary file for your OS. For windows it'll be called `ResourceChecker.exe`. With this approach you can use the module a bit differently:
```cmd
ResourceChecker.exe [JSON data path]
```
## Change default JSON data
You can either change [source.json](https://github.com/DroppedHard/AWS_RP_Check/blob/main/source.json) file locally, or give path to other JSON data you wish to verify.

You can use examples in [data](https://github.com/DroppedHard/AWS_RP_Check/tree/main/data) folder, which are grouped by expected output:
- [invalid](https://github.com/DroppedHard/AWS_RP_Check/tree/main/data/invalid) data should result with error explaining reason why given JSON data is not a JSON AWS::IAM::Role Policy data.
- [false](https://github.com/DroppedHard/AWS_RP_Check/tree/main/data/false) and [true](https://github.com/DroppedHard/AWS_RP_Check/tree/main/data/true) group JSON data which should load properly, and represent expected output from `AwsRolePolicyChecker.verifyResource()` method.

## Testing
To run all module tests (located in [AWS_RP_Checker_test.go](https://github.com/DroppedHard/AWS_RP_Check/blob/main/AWS_RP_Checker_test.go) file) use:
```cmd
go test [-v]
```
This command will run all tests and print summarized output. For verbose output (outcome of each test separately) add `-v` flag at the end of the command.