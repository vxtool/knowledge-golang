package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(math.pi)
}


// Nomes exportados
// Depois de importar um pacote, você pode consultar os nomes que exporta.
// Em Go, um nome é exportado se começa com uma letra maiúscula.
// Foo é um nome exportado, assim como FOO. O nome foo não é exportado.
// Execute o código. Em seguida, renomeie math.pi para math.Pi e tente novamente.
