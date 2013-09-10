package main

import "fmt"
import "os"
import "strconv"

/*
  Escreva uma função que recebe como parâmetro um
  número inteiro positivo e retorna o fatorial
  desse número.
*/

func main(){
  if len(os.Args) < 2 {
    fmt.Println("Usage: `$ go run src/fatorial.go $NUMBER`")
    return
  }
  n, err := strconv.Atoi(os.Args[1])
  if err == nil{
    fmt.Printf("Fatorial do número %q\n\t", os.Args[1])
    fmt.Println(factorial(n))
  }
}

func factorial(number int) int {
  if number < 2 { return number }
  return number * factorial(number - 1)
}
