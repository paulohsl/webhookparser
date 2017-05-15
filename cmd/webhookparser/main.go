package main

import (
	"os"
	"fmt"
	. "github.com/paulohsl/webhookparser"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Filepath is missing usage: webhookparser <filepath>")
		os.Exit(1)
	}

	file := ReadFile(os.Args[1])

	urls, statuses, err := Parse(file)
	if err != nil {
		fmt.Println(err)
	}

	PrintTopRanked(urls, 3)
	PrintRanked(statuses)

}
