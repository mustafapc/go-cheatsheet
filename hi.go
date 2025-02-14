package main

import "fmt"
import (
	"math"
	"math/rand")

func add(x int, y int) int {
return x + y
}
func main(){
	fmt.Println("hello world", rand.Intn(15), "pi is equal to:", math.Pi, "5 + 8", add(5, 8))
}