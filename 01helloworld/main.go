package main

import "fmt"

func main() {
    i := 42
    f := float64(i)
    u := uint(f)
fmt.Printf(" f: %T, i: %T, u: %T\n",  f, i, u)
}