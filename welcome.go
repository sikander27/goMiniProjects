package main

import "fmt"

func main(){
	var name string;
	fmt.Println("Please Enter your name")
	fmt.Scanf("%s", &name)
	fmt.Printf("Welcome, %v", name)
}