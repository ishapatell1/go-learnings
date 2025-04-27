
package main
import "fmt"
func reversedigit(n int){
fmt.Println("Given an integer N return the reverse of the given number")
if n < 0 {
	n = -n // Convert negative number to positive
}
reverse:=0
for n!=0{
	
	lastdigit := n%10
	reverse = reverse*10+lastdigit

	n= n/10
}
fmt.Println("reverse :",reverse)
}