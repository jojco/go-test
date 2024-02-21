package main

import (
	"fmt"
	"github.com/jojco/go-test/jq"
)

func main() {
	fmt.Println("hello")
	jq.GetEnvWithDefault("wef", "wef")
}
