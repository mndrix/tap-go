package main

import "github.com/mndrix/tap.go"

func main() {
	tap.Header(2)
	tap.Ok(1, "first test")
	tap.Ok(2, "second test")
}
