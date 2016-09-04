package main

import "fmt"

func main() {
  var i, j int = 1, 2
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}

// Declarações curtas de variáveis
// Dentro de uma função a instrução de atribuição curta := pode ser utilizada em lugar de uma declaração var com o tipo implícito.
// Fora de uma função cada estrutura começa com uma palavra-chave (var, func, e assim por diante) e não é possível usar o :=.
