package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}

// Ponteiros para structs
// Campos de estruturas podem ser acessados através de um ponteiro de estrutura.
// A indireção através do ponteiro é transparente.
