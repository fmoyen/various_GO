package main

import ("fmt")

// main function
func main() {
  fmt.Print("Enter Your First Name: ")
  var first string
  fmt.Scanln(&first)
  fmt.Print("Enter Second Last Name: ")
  var second string
  fmt.Scanln(&second)
  fmt.Print("Your Full Name is: ")
  fmt.Println(first + " " + second)
}
