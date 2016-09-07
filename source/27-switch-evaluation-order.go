package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In two days.")
  default:
    fmt.Println("Too far away.")
  }
}

// Switch com ordem de avaliação
// Switch cases avaliam casos de cima para baixo, parando quando um caso for bem-sucedido.
// (Por exemplo,
// switch i {
// case 0:
// case f():
// }
// não chamar f se i==0.)
// Nota: Tempo no playground Go sempre aparece para começar às 2009-11-10 23:00:00 UTC, um valor cujo significado é deixado como um exercício para o leitor.
