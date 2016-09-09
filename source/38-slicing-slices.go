package main

import "fmt"

func main() {
  p := []int{2, 3, 5, 7, 11, 13}
  fmt.Println("p ==", p)
  fmt.Println("p[1:4] ==", p[1:4])

  // missing low index implies 0
  fmt.Println("p[:3] ==", p[:3])

  // missing high index implies len(s)
  fmt.Println("p[4:] ==", p[4:])
}

// Fatiando slices
// Slices podem ser re-fatiados, criando um valor de uma nova slice que aponta para a mesma matriz.
// A expressão
// s[lo:hi]
// avalia para uma slice de elementos de lo através de hi-1, inclusive. Assim
// s[lo:lo]
// é vazia e
// s[lo:lo+1]
// tem um elemento.
