package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
}

func main() {
  fmt.Println(Sqrt(2))
}

// Exercício: Laços e Funções
// Como uma forma simples de brincar com as funções e laços, implemente a função de raiz quadrada usando o método de Newton.
// Neste caso, o método de Newton é aproximar Sqrt(x), escolhendo um ponto de partida z e depois repetindo:
// Para começar, basta repetir esse cálculo 10 vezes e ver o quão perto você chega à resposta para vários valores (1, 2, 3, ...).
// Em seguida, altere a condição do laço para parar uma vez que o valor deixou de mudar (ou apenas mudanças por um delta muito pequeno). Veja se isso são muitas ou poucas iterações. Como ainda falta para o math.Sqrt?
// Dica: para declarar e inicializar um valor de ponto flutuante, dê-lhe a sintaxe de ponto flutuante ou use uma conversão:
// z := float64(1)
// z := 1.0

