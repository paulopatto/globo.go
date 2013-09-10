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
    fmt.Printf("Fatorial do número %q\n", os.Args[1])
    fmt.Printf("recursivo\n\t")
    fmt.Println(factorial(n))
    fmt.Printf("Agora com loop for while\n\t")
    fmt.Println(factorial2(n))
  }
}

func factorial(number int) int {
  if number < 2 { return number }
  return number * factorial(number - 1)
}

func factorial2(number int) int{
  fat := 1
  for ; number > 1 ; {
    fat *= number
    number--
  }
  return fat
}
