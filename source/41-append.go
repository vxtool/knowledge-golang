package main

import "fmt"

func main() {
  var a []int
  printSlice("a", a)

  // append works on nil slices.
  a = append(a, 0)
  printSlice("a", a)

  // the slice grows as needed.
  a = append(a, 1)
  printSlice("a", a)

  // we can add more than one element at a time.
  a = append(a, 2, 3, 4)
  printSlice("a", a)
}

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n",
    s, len(x), cap(x), x)
}

// Adicionando elementos em um slice
// É commum acrescentar novos elementos em um slice, então Go dispões de uma função built-in append para isso. A documentação do pacote built-in descreve melhor o append.
// func append(s []T, vs ...T) []T
// O primeiro parâmetro s do append é um slice do tipo T, e o resto são valores de T para acrescentar no slice.
// O valor resultante do append é um slice contendo todos os elementos do slice original mais os valores providos.
// Se a matriz anterior de s fpr muito pequena para caber todos os valores uma matriz gigante será alocada. O slice retornado apontara para a nova matriz alocada.
// (Para aprendar mais sobre slices, leia o artigo Slices: usage and internals.)
