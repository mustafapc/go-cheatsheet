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
	a, b := add(5, 8);
	fmt.Println("hello world", rand.Intn(15), "pi is equal to:", math.Pi, "5 + 8", a, "adsf", b)
}