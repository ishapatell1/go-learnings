package main

import "fmt"
func main(){
	fmt.Println("This is list of patterns printed using go")
	rectangle(5)
	pattern02(6)
	pattern03(5)
	pattern04(5)
	pattern05(5)
	pattern06(5)
	pattern07(5)
	pattern08(5)
	pattern09()
	pattern10(5)
	pattern11(5)
	pattern12(5)
	pattern13(5)
	pattern14(5)
	pattern15(5)
	pattern16(5)
	pattern17(2)
	pattern18(8)
	printNumberAsTriangle()
	printRectangle()
	printTriangle()

}
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
func rectangle(n int){
	for i:=0; i<n; i++{
		for j:=0;j<n; j++{
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}
func pattern02(n int){
	for i:=0;i<n;i++{
		for j:=0; j<i;j++{
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}
func pattern04(n int){
	for i:=1; i<=n; i++{
		for j:=0; j<i; j++{
			fmt.Print( i )
		}
		fmt.Println()
	}
	
}
func pattern03(n int){
	for i:=1; i<n; i++{
		for j:=1; j<=i; j++{
			fmt.Print(j)
		}
		fmt.Println()
	}
	
}
func pattern05(n int){
	for i:=n; i>0; i--{
		for j:=0; j<i; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func pattern06(n int){
	for i:=n; i>0; i--{
		for j:=1; j<=i; j++{
			fmt.Print(j)
		}
		fmt.Println()
	}
}
func pattern07(n int){
	for i:=0; i<n; i++{
		for j:=0; j<n-i-1; j++{
			fmt.Print("-")
		}
			for j:=0; j<2*i+1; j++{
				fmt.Print("*")
			}
			for j:=0; j<n-i-1; j++{
			fmt.Print("-")
			}
			fmt.Println()
		}
		
		
	}
func pattern08(n int){
		for i:=0; i<n; i++{
			for j:=1; j<=i; j++{
				fmt.Print("-")
			}
				for j:=1; j<=2*n-(2*i+1); j++{
					fmt.Print("*")
				}
				for j:=1; j<=i; j++{
				fmt.Print("-")
				}
				fmt.Println()
			}
			
			
		}
func pattern09(){
	pattern07(5)
	pattern08(5)
}
func pattern10(n int){
	for i:=0; i<=n; i++{
		for j:=0;j<i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n-1; i>=0; i--{
		for j:=1; j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func pattern11(n int){
	start := 1
	//define start = 1 and do i = 1-start
	for i:=0 ; i<n ; i++{
		if i%2==0{start =1}else {start = 0}
		for j:=0; j<i; j++{
			fmt.Print(start)
			start = 1-start
		}
		fmt.Println()
	}

	
}
func pattern12(n int){
for i:=1; i<=n; i++{
	for j :=1; j<=i; j++{
		fmt.Print(j)
	}
	for k:=1; k<=2*(n-i); k++{
		fmt.Print(" ")
	}
	for j:=i; j>=1; j--{
		fmt.Print(j)
	}
	fmt.Println()
}
}
func pattern13(n int){
	num :=1
	for i:=0; i<n; i++{
		for j:=0; j<=i; j++{
			fmt.Print(num)
			num++;
		}
		fmt.Println()
	}
}
func pattern14(n int){
	for i:=0; i<n; i++{
		for j:=0; j<=i;j++{
			fmt.Printf("%c",'A'+j)
		}
		fmt.Println()
	}	
} 
//in go, "" double string is for string but single quote is for rune(a fancy name for integer that represets a unicode character)
func pattern15(n int){
	for i:=n; i>=1; i--{
		for j:=0; j<i;j++{
			fmt.Printf("%c",'A'+j)
		}
		fmt.Println()
	}
}
func pattern16(n int){
	for i:=0;i<n;i++{
		char := 'A'+ rune(i)
		for j:=0;j<=i;j++{
			
			fmt.Printf( "%c",char )
		}
		fmt.Println()
	}
}
func pattern17(n int){
	for i:=0; i<n; i++{
		for j:=0; j<=i;j++{

		}
		for j:= i-1;j>=0;j--{
			fmt.Println( )
		}
	}
}
func pattern18(n int){
	for i:=n; i<=1;i++{
		char := 'E'-1
		for j:=i;j<i;j++{
			fmt.Printf("%c", char)
		}
	}	
}