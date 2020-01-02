package main

import ("fmt")
import ("os")

func main(){

var name string
var age uint8
var id uint32

fmt.Println("Enter your name:")
fmt.Fscan(os.Stdin, &name)

fmt.Println("Enter your age:")
fmt.Fscan(os.Stdin, &age)

fmt.Println("Enter your id:")
fmt.Fscan(os.Stdin, &id)
     
fmt.Println("Your data has been saved as:", name, age, id)

	   }
