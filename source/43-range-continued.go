package main

import "fmt"

func main() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i)
  }
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}

// Range continuação
// Você pode ignorar o índice ou valor, atribuindo o _.
// Se você só quiser o índice, deixe inteiramente sem o ", value".
