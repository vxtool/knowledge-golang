package main

import "fmt"

func main() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}

// For é o "while" de Go
// Nesse ponto, você pode tirar as vírgulas: o while do C é escrito com for em Go.
