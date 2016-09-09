package main

import "fmt"

func main() {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)
}

// Matrizes
// O tipo [n]T é uma matriz de n valor do tipo T.
// A expressão
// var a [10]int
// declara a variável a com uma matriz de dez inteiros.
// O tamanho da matriz é parte de seu tipo, logo, matrizes não podem ser redimencionadas. Isso parece uma limitação, mas não se preocupe; Go dispões de uma maneira conveniente para trabalhar com matrizes.
