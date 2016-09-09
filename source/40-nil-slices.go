package main

import "fmt"

func main() {
  var z []int
  fmt.Println(z, len(z), cap(z))
  if z == nil {
    fmt.Println("nil!")
  }
}

// Slices vazios
// O valor zero de uma slice é nil.
// Uma slice nil tem um comprimento e capacidade de 0.
// (Para saber mais sobre as slices, leia o artigo Slices: usage and internals .)
