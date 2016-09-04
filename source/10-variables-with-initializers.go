package main

import "fmt"

var i, j int = 1, 2

func main() {
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}


// Variáveis com inicializadores
// A declaração var pode incluir inicializadores, uma por variável.
// Se um inicializador está presente, o tipo pode ser omitido; a variável terá o tipo do inicializador.
