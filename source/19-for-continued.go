package main

import "fmt"

func main() {
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum)
}

// For (continuação)
// Como em C ou Java, você pode deixar as declarações do começo e do final vazias.
