# Implementação dos testes da Globo.com distribuidos na RubyConf 2013

A composição deste teste de composta de..


## Ping-Pong

Escreva um programa que imprime na tela números de 1 a 100, um por linha. Existem 2 exceções:
1. Quando o número for múltiplo de 3, imprima apenas a palavra `ping` em vez do número.
2. Quando o número for múltiplo de 5, imprima apenas a palavra `pong` em vez do número.

### Bonus point

- Existe uma terceira exceção que não foi especificada?
> Sim claro, e no caso do númeor `15` que é múltiplo tanto de `3` quanto de `5`? Eu imprimo ping pong?

- Você consegue escrever testes automatizados para garantir que a saida do programa está correta?
> Sim e no caso vamos tentar usar [go-spec](https://github.com/bmatsuo/go-spec) para esse propósito. E vamos tentar rodar no [travis-ci](http://about.travis-ci.org/docs/user/languages/go/)

## Fatorial

Escreva uma função que recebe como parâmetro um número inteiro positivo e retorna o fatorial desse número.

### Bonus point

- Como você pode proteger as implementação (tava escrito função) de parâmetros inválidos.
- Existem duas formas de implementar essa função: **Iterativa** e **Recursiva**. Você consegue implementar de ambas as formas?
- Quais as vantagens e desvantagens da implementação *iterativa* e *recursiva*?
- Você consegue escrever testes automatizados para garantir que a saida do programa está correta?

## Frequencia

Escreva um programa que recebe como parâmetro uma string qualquer e retorna um dicionário/hashtable em qual as chaves são todos os caracteres que ocorrem na string. Não deve ser utilizada nenhuma função pronta da linguagem para a contagem de caracteres (*mancada*).

### Bonus point

- Mutar a função para não distinguir a contagem entre letras maiúsculas e minúsculas.
- Sua implementação possúi complexidade `O(n)`?
- Você consegue escrever testes automatizados para garantir que a saida do programa está correta?

