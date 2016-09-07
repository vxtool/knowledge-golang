package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.", os)
  }
}

// Switch
// Você provavelmente sabia que o switch estava chegando.
// O bloco case falhará automaticamente, a menos que termine com uma declaração default.
