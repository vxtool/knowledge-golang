package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
  pic.Show(Pic)
}

// Exercício: Slices
// Implemente Pic. Ela deve retornar uma slice de comprimento dy, cada elemento do qual é uma slice de dx 8-bit inteiros sem sinal. Quando você executa o programa, ele irá exibir a sua imagem, interpretando os números inteiros como escala dos valores de cinza (bem, bluescale).
// A escolha da imagem é com você. Funções interessantes incluem x^y, (x+y)/2, e x*y.
// (Você precisa usar um loop para alocar cada []uint8 dentro do [][]uint8.)
// (Utilize uint8(intValue) para converter entre os tipos.)
