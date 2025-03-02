package main

import "fmt"
import (
	"math"
	"math/rand")
/*or 
func add(x int, y int) (a int,b int) {
return x + y, 3
}*/ 
func add(x int, y int) (int, int) {
return x + y, 3
}
func main(){
	/*types are int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64
	byte = uint8
	*/ 
	a, b := add(5, 8)
	var c int
	var d int
	c, d = add(3, 8)
	fmt.Println("hello world", rand.Intn(15), "pi is equal to:", math.Pi, "5 + 8", a, "adsf", b, c, d)
}