package main

import (
	"fmt"
	"lets-go/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH")) // test-comment 1
	fmt.Println(cmp.Diff("Hello, World", "Hello Go"))   // test-comment 2
}
