package main

import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(17))
}


// Valores nomeados de retorno
// Valores de retorno de Go podem ser nomeados e agirem apenas como variáveis.
// Esses nomes devem ser usados para documentar o significado dos valores de retorno.
// A declaração return sem argumentos retorna os valores atuais dos resultados. Isto é conhecido como um retorno "limpo".
// Instruções de retorno limpas devem ser usadas apenas em funções curtas, como no exemplo mostrado aqui. Elas podem prejudicar a legibilidade em funções mais longas.
