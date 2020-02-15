package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.Join(os.Args[1:], " ")
	fmt.Printf("Command: %s Args: %s", os.Args[0], s)
}
