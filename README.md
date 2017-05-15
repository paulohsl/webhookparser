# Simple WebHook Log Parser

This is a simple WebHook Log Parser. It process and summarizes Url's and HTTP Status.

## Design Considerations
- Use bufio.Scanner to read file line by line avoiding load all file into memory.
- Use Golang to best performance parsing file
- Use cmd package structure to reuse parser as a library

## Setup
Instructions:

1. Download Go.
Download the latest version of Go for your platform here: https://golang.org/dl/.

2. Install Go.
Follow the instructions for your platform to install the Go tools: https://golang.org/doc/install#install. 
It is recommended to use the default installation settings.

3. Test your Go installation.

Create and run the hello.go application described here: https://golang.org/doc/install#testing.

If you set up your Go environment correctly, you should be able to run “hello” from any directory and see the program execute successfully.

### Installing
Follow the steps bellow to Install Package Application on your local machine
```
go get github.com/paulohsl/webhookparser/
cd $GOPATH/src/github.com/paulohsl/webhookparser/cmd/webhookparser
go install
```

## Running the tests

Explain how to run the automated tests for this system


## Authors

* **Paulo Henrique S. Lopes** - *Initial work* -(https://github.com/paulohsl)
