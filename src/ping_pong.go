package main

import "fmt"
/*
  Escreva um programa que imprime na tela números de 1 a 100, um por linha.
  Existem 2 exceções:

    1. Quando o número for múltiplo de 3, imprima apenas a palavra `ping` em vez do número.
    2. Quando o número for múltiplo de 5, imprima apenas a palavra `pong` em vez do número.
*/

func main(){
  fmt.Println("The Ping-Pong program")
  fmt.Println("=====================")
  fmt.Println("\nThis program print number between 1 and 100 but, if number is multiply 3 it print word ping, if number multiply 5 print word pong...\n")
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 && i % 5 == 0 {
      fmt.Println("* Ping-Pong *")
    } else if i % 3 == 0 {
      fmt.Println("** ping **")
    } else if i % 5 == 0 {
      fmt.Println("** pong **")
    } else {
      fmt.Println(i)
    }
  }

}
