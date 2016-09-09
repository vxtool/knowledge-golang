package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2}  // has type Vertex
  v2 = Vertex{X: 1}  // Y:0 is implicit
  v3 = Vertex{}      // X:0 and Y:0
  p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
  fmt.Println(v1, p, v2, v3)
}

// Struct literais
// A struct literal denota um valor struct recém-alocado, ao enumerar os valores de seus campos.
// Você pode listar apenas um subconjunto de campos usando o Name: sintaxe. (E a ordem dos campos nomeados é irrelevante.)
// O prefixo especial & constrói um ponteiro para uma struct literal.
