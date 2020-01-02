package main

import ("fmt")

func main(){

var numbers [11]uint8 = [11]uint8{0,1,2,3,4,5,6,7,8,9,10}
     
fmt.Println(numbers)
fmt.Println(numbers[1])
fmt.Println(numbers[3])
fmt.Println(numbers[5])

numbers[1] = 99
numbers[3] = 66
numbers[5] = 33

fmt.Println(numbers)
fmt.Println(numbers[1])
fmt.Println(numbers[3])
fmt.Println(numbers[5])

	   }
