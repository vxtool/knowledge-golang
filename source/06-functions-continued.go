package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}

// Funções continuação
// Quando dois ou mais consecutivos parâmetros da função nomeados compartilhar um tipo, você pode omitir o tipo de todos, menos o último.
// Neste exemplo, vamos encurtar
// x int, y int
// para
// x, y int
