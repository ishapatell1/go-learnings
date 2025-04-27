package main

import "fmt"

func main(){
	fmt.Println("This is Rectangle Pattern")
	printRectangle()
	fmt.Println("This is Triangle Pattern")
	printTriangle()
	printNumberAsTriangle()
}
//Print Rectangle
func printRectangle(){
	N:=3
	for i:=0;i<N; i++ {
		for j:=0; j<N;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
//Print Triangle
func printTriangle(){
	N:=5
	for i:=0; i<N;i++{
		for j:=0; j<=i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
//Print Numbers as Triangle
func printNumberAsTriangle(){
	N:=5
	for i:=1; i<=N; i++{
		for j:=1; j<=i;j++{
			fmt.Print(j)
		}
		fmt.Println()
	}
}