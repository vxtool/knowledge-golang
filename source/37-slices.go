package main

import "fmt"

func main() {
  p := []int{2, 3, 5, 7, 11, 13}
  fmt.Println("p ==", p)

  for i := 0; i < len(p); i++ {
    fmt.Printf("p[%d] == %d\n", i, p[i])
  }
}

// Slices
// Um slice aponta para uma matriz de valorqes que também inclui o tamanho.
// []T é um slice com elementos do tipo T.
