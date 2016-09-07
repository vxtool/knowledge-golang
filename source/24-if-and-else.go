package main

import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though
  return lim
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}

// If e else
// Variáveis ​​declaradas dentro de uma instrução if curta também estão disponíveis dentro dos blocos else.
