package main

import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println("My favorite number is", rand.Intn(10))
}

// Pacotes
// Cada programa Go é composto de pacotes.
// Programas começam rodando pelo pacote main.
// Este programa está usando os pacotes com caminhos de importação "fmt" e "math".
// Por convenção, o nome do pacote é o mesmo que o último elemento do caminho de importação.
// Nota: o ambiente em que esses programas são executados é determinístico, então rand.Intn sempre retornará o mesmo número. (Para ver um número diferente, a semente do gerador de números; veja rand.Seed.)
