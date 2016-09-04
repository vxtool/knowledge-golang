package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}

// Importações
// Este grupo de códigos em parênteses da importação, é a declaração de importação "consignada". Você também pode escrever várias declarações de importação assim:
// import "fmt"
// import "math"
// Mas é um bom estilo usar a declaração de importação consignada.
