package main

import (
	"gowork/base"
	"gowork/foundation"
)

func main() {
	testPointer()
	testSlice()
}

func testPointer() {
	base.PointTest()
	base.FlagPara()
}

func testSlice() {
	foundation.Slice1()
}
