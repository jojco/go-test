package main

import (
	"fmt"
	"github.com/jojco/go-test/utils"
)

func main() {
	fmt.Println("hello")
	jojcoutils.GetEnvWithDefault("wef", "wef")
}
