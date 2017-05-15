package webhookparser

import (
	"fmt"
	. "os"
	. "bufio"
)

func ReadFile(filePath string) (*Scanner) {
	file, err := Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	return NewScanner(file)
}
