package main

import (
	"fmt"
	"time"
)
func main(){
day := int(time.Now().Weekday()) 

if(day== 0 ){
day=7
}
fmt.Println("Day number:", day)
switch day {
case 1 : fmt.Println("Monday")
case 2 : fmt.Println("Tuesday")
case 3: fmt.Println("Wednesday")
case 4: fmt.Println("Thursday")
case 5: fmt.Println("Friday")
case 6: fmt.Println("Saturday")
case 7 : fmt.Println("Sunday")
default : fmt.Println("Enter Any number between 1 to 7")
}
}
