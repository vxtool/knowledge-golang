package main

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i         // point to i
  fmt.Println(*p) // read i through the pointer
  *p = 21         // set i through the pointer
  fmt.Println(i)  // see the new value of i

  p = &j         // point to j
  *p = *p / 37   // divide j through the pointer
  fmt.Println(j) // see the new value of j
}

// Ponteiros
// Go tem ponteiros. Um ponteiro guarda na memória o endereço de uma variável.
// O tipo *T é um ponteiro para um valor T. Seu valor zero é nil.
// var p *int
// O operador & gera um ponteiro para seu operando.
// i := 42
// p = &i
// O operador * indica valor subjacente do ponteiro.
// fmt.Println(*p) // lê i através do ponteiro p
// *p = 21         // defina i através do ponteiro p
// Isto é conhecido como "desreferenciamento" ou "indirecionamento".
// Diferente de C, Go não faz aritimética de ponteiros.
