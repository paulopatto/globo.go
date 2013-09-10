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
    fmt.Printf("- Recursivamente: \n\t")
    fmt.Println(Factorial(n))
    fmt.Printf("- Iterativamente: \n\t")
    fmt.Println(Factorial2(n))
  }
}

func Factorial(number int) int {
  if number < 2 { return 1 }
  return number * Factorial(number - 1)
}

func Factorial2(number int) int{
  fat := 1
  for ; number > 1 ; {
    fat *= number
    number--
  }
  return fat
}
