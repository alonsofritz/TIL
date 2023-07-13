# TIL: Go
Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Gera build para qualquer sistema operacional de qualquer sistema operacional. Atualmente o compilador do Go é feito em Go, logo é uma linguagem que se auto compila.


## Package Imports

    import ( "fmt" )
    
Nota-se que "fmt" refere-se a um package nativo do Go, e o mesmo possui a maioria das funcionalidades necessarias para implementação, evitando a necessidade de instalar pacotes de terceiros. A maioria dos outros pacotes para Go estão disponiveis no próprio Github, diferente do Docker por exemplo que possui um Hub próprio.

## Exportações

O Go trabalha com métodos, variáveis e funções que podem ser exportadas ou não exportadas, isso significa que todas as vezes ao acessar um pacote externo, só terei acesso aos métodos que forem exportados. Mas o que define se um método é ou não exportado? Para isso a primeira letra utilizada na declaração desses métodos devem ser em letra maiúscula (UPPERCASE), para definir então algo "privado", basta declarar iniciando com letra minúscula (lowercase).

## Tratamento de Erros
Em Go, o erro sempre será algo explícito no momento em que trabalharmos com funções. Podemos ter multiplos valores no retorno de uma função, em sua maioria um desses retornos é um "erro", que trás consigo informações sobre o erros que tenham ocorrido na função em questão.

## Bons hábitos
- Sempre que exportar métodos, variáveis e/ou funções, é interessante adicionar um comentário/documentação sobre o método a ser exportado, de forma a prover uma documentação sobre as responsabilidades do alvo exportado. Deve-se sempre inicar o comentário com o nome do método, dessa forma ele reconhece o formato de documentação.