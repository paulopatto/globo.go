package main

import "fmt"
// import "strings"
import "regexp"
import "os"

/*
  Escreva um programa que recebe como parâmetro uma string qualquer e retorna um dicionário/hashtable
  no qual as chaves são todos os caracteres que ocorrem na string.
  Não deve ser utilizada nenhuma função pronta da linguagem para a contagem de caracteres.
*/

func main(){
  var input_string string
  input_string = "Programa escrito em Go!"
  //fmt.Printf("Entre com uma frase: ")
  //fmt.Scanln(&input_string)
  /*
    a: 2, e: 2, o: 3, i: 1
    P: 1, r: 3, g: 1, m: 2, s: 1, c: 1, t: 1, G: 1
    ----------------------------------------------
    Regex case-sensitive:   /[a-z]/g | /[A-Z]/g
    Regex case-insensitive: /[A-Za-z]/g
    Regex marotinha:        /[\w]/g
    ----------------------------------------------

  */
  // m := false
  // usage()
  /*
  if m {
    s := regexp.MustCompile("[a-z]").FindAllString(input_string, -1)
    fmt.Printf("%q\n", s)
  } else {
    s := regexp.MustCompile("[a-zA-Z]").FindAllString(input_string, -1)
    fmt.Printf("%q\n", s)
  }
  */
  // fmt.Printf("%q\n", strings.Split(msg, ""))
  s := regexp.MustCompile("[a-zA-Z]").FindAllString(input_string, -1)
  fmt.Printf("%q\n", s)
  fmt.Println(os.Args[0])
}
func usage(){
  fmt.Printf("Usage command: frequencia --options $string\n")
  fmt.Printf("Options: \n")
  fmt.Printf("\t--ic \tIgnore case string\n")
}
