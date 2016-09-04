package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}


// Resultados Múltiplos
// Uma função pode retornar qualquer número de resultados.
// Esta função retorna duas strings.
