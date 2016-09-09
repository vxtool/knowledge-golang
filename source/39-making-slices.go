package main

import "fmt"

func main() {
  a := make([]int, 5)
  printSlice("a", a)
  b := make([]int, 0, 5)
  printSlice("b", b)
  c := b[:2]
  printSlice("c", c)
  d := c[2:5]
  printSlice("d", d)
}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n",
    s, len(x), cap(x), x)
}

// Fazendo slices
// Slices são criados com a função make. Ele funciona através da atribuição de uma matriz zerada e retornando uma slice que se refere a essa matriz:
// a := make([]int, 5)  // len(a)=5
// Para especificar uma capacidade, passe um terceiro argumento para make:
// b := make([]int, 0, 5) // len(b)=0, cap(b)=5
// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:]      // len(b)=4, cap(b)=4
