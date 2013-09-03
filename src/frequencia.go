package main
import "fmt"

/*
  Escreva um programa que recebe como parâmetro uma string qualquer e retorna um dicionário/hashtable
  no qual as chaves são todos os caracteres que ocorrem na string.
  Não deve ser utilizada nenhuma função pronta da linguagem para a contagem de caracteres.
*/

func main(){
  usage()
}
func usage(){
  fmt.Printf("Usage command: frequencia --options $string\n")
  fmt.Printf("Options: \n")
  fmt.Printf("\t--ic \tIgnore case string\n")
}
