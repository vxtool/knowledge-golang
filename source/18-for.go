package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}

// For
// Go tem apenas uma estrutura de laço, o for.
// O for básico parece com o que o de C ou Java fazem, exceto que as ( ) se foram (não são nem mesmo opcionais) e os { } são obrigatórios.
