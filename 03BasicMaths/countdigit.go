package main

import (
	"fmt"
    // "strconv"
)
func countdigit(n int){
    number :=n
    if n ==0 {
        fmt.Println("The number ", n, " has 1 digit ")
        return
    }
    count := 0; 
    for n!=0{
        n= n/10
        count++;
    }
   
    fmt.Println( number ,"has", count,"digits")
}
