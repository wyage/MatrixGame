package main

import (
	"matrix"
)

func main() {
	testGenerate()
}

func testGenerate() {
	m1 := matrix.New()
	m1.GeneratePlay()
	m1.PrintMe()
}
