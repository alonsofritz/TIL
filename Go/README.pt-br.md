# TIL: Go
Go é expressivo, conciso, limpo e eficiente. Seus mecanismos de simultaneidade facilitam a escrita de programas que tiram o máximo proveito de máquinas multicore e em rede, enquanto seu novo tipo de sistema permite a construção flexível e modular de programas. O Go compila rapidamente para o código de máquina, mas tem a conveniência da coleta de lixo e o poder da reflexão em tempo de execução. É uma linguagem compilada, rápida e estaticamente digitada que parece uma linguagem interpretada e digitada dinamicamente.

Gera build para qualquer sistema operacional de qualquer sistema operacional. Atualmente o compilador do Go é feito em Go, logo é uma linguagem que se auto compila.


## Package Imports

    import ( "fmt" )
    
Nota-se que "fmt" refere-se a um package nativo do Go, e o mesmo possui a maioria das funcionalidades necessarias para implementação, evitando a necessidade de instalar pacotes de terceiros. A maioria dos outros pacotes para Go estão disponiveis no próprio Github, diferente do Docker por exemplo que possui um Hub próprio.

## Exportações

O Go trabalha com métodos, variáveis e funções que podem ser exportadas ou não exportadas, isso significa que todas as vezes ao acessar um pacote externo, só terei acesso aos métodos que forem exportados. Mas o que define se um método é ou não exportado? Para isso a primeira letra utilizada na declaração desses métodos devem ser em letra maiúscula (UPPERCASE), para definir então algo "privado", basta declarar iniciando com letra minúscula (lowercase).

## Tratamento de Erros
Em Go, o erro sempre será algo explícito no momento em que trabalharmos com funções. Podemos ter multiplos valores no retorno de uma função, em sua maioria um desses retornos é um "erro", que trás consigo informações sobre o erros que tenham ocorrido na função em questão.

## Funções e Métodos
Ao declarar funções utiliza-se a palavra reservada *func* como podemos notar no exemplo:

    func  main() {
	    //code here
    }

Perceba que a função anterior não necessita e portanto não possui nenhum tipo de retorno. Para que uma função possua retorno é necessário declarar previamente no escopo da função, como no caso a seguir:

    func  sum(a int, b int) int {
	    return a + b
    }

Caso contrário, mesmo que exista um *return* ao compilar teremos um erro, como ocorre no exemplo:

    func  sumError(a int, b int) {
	    return a + b
    }

Perceba que a diferença entre as funções acontece logo após a declaração dos parâmetros da função e antes de abrir as chaves, esse é o local onde declaramos o tipo do retorno da função.

Agora que sabemos como declarar uma função com retorno, podemos entender que em *Go* nossas funções podem ter múltiplos retornos, como a seguir:

    func  sumAndReturnString(a int, b int) (int, string) {
	    // Retorna respectivamente um int e uma string
	    return a + b, "Added!"
    }

Em Go temos também o que chamamos de *Retorno Nomeado*, que nada mais é do que definir previamente um *nome* para o seu retorno. Isso pode ser feito, definindo um nome a esquerda da declaração de tipo de retorno, dessa forma, tendo um *Retorno Nomeado* não é necessário explicitar o que é retornado na função, apenas a palavra reservada *return* basta, pois o próprio Go entende o que se espera do retorno daquela função, como no exemplo:

    func  sumNamedReturn(a int, b int) (result int) {
	    result = a + b
	    return
    }

Outra variação que temos são as *Funções Variáticas* onde podemos passar uma quantidade indefinida de parâmetros para a função, porém atente-se que **todas devem ser do tipo definido na declaração da função,** veja abaixo: 

    // FUNCTION
    func  sumAll(x ...int) int {
	    result :=  0
	    
	    for _, v :=  range x {
		    result += v
	    }
	    
	    return result
    }

O responsável por permitir a passagem de **n** parâmetros para uma função é o trecho `(x ...int)`, onde **x** funciona como array de valores do tipo definido, no caso *int*. O resto da função utiliza-se de um loop para caminhar pelos valores de **x** e dessa forma somar todos os valores passados para a função, nota-se que temos a utilização de um **blank identifier** o **_** que neste caso é utilizado para descartar a key/index retornado pelo loop de **x**. Para fazermos a chamada de uma função deste tipo, fazemos o seguinte: 

    resultAll :=  sumAll(10, 20, 3, 5, 6, 4)

Podemos também possuir uma variável que é implicitamente atribuída com o tipo de uma função, que tem como retorno o tipo `func() int` que por sua fez, retornará um tipo `int` como demonstrado a seguir:

    resultFunc :=  func(x ...int) func() int {
	    resultTemp :=  0
	    for _, v :=  range x {
		    resultTemp += v
	    }
	    return  func() int {
		    return resultTemp * resultTemp
	    }
    }

Para entendermos melhor o que essa função retorna podemos fazer:

    fmt.Println(resultFunc(54, 54, 54, 54))
   e
   
    fmt.Println(resultFunc(54, 54, 54, 54)())

No primeiro caso temos um retorno diferente, ou seja, ao invés de obtermos o valor final que deveria ser retornado pela função interna dentro da função, obtemos na verdade a posição da memoria deste retorno. Já no segundo caso, como estamos chamando a função, passando os parâmetros e por fim executando a função, entende-se que estamos executando a função dentro da função, e portanto o retorno que obtemos dessa forma é o esperado, ou seja, um retorno do tipo *int* com as operações dos valores passados por parâmetro realizadas.
Algo sem dúvidas muito importante em Go são a "transformação" de uma função em um métodos, atenta-se que o Go possui uma orientação objeto propriamente dita, temos o que chamamos de GoWay, ou seja, a forma Go de construir as coisas, dessa forma, para criarmos **métodos** necessitamos de duas coisas, primeiramente uma Struct e por segundo uma função atrelada a essa Struct, mas como? Veja abaixo:

    type  Car  struct {
	    Name string
    }

Temos uma struct *Car* com a propriedade Name do tipo String. Agora para criar o método, precisamos declarar uma função praticamente como normalmente fazemos, com uma única diferença de que, desta vez faremos um tipo de *Binding* para atrelhar a função a Struct anterior, e dessa forma tornar-se um método da Struct, veja:

    func (c Car) drive() {
	    fmt.Println(c.Name, "Drive")
    }

Perceba que a estrutura para declarar a função muda ligeiramente, de forma que temos a palavra reservada *func* e depois o binding propriamente dito, ao fazermos `(c Car)` antes de nomear a função.

## Bons hábitos
- Sempre que exportar métodos, variáveis e/ou funções, é interessante adicionar um comentário/documentação sobre o método a ser exportado, de forma a prover uma documentação sobre as responsabilidades do alvo exportado. Deve-se sempre inicar o comentário com o nome do método, dessa forma ele reconhece o formato de documentação.

## Códigos Fonte
- [Introdução: Hello World](https://github.com/alonsofritz/TIL/tree/master/Go/hello-go)
- [Tratamento de Erros](https://github.com/alonsofritz/TIL/tree/master/Go/error-handling-go)
- [Funções e Métodos](https://github.com/alonsofritz/TIL/tree/master/Go/functions-go)

##
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/alonsofritz/TIL/tree/master/Go/README.md) [![pt-br](https://img.shields.io/badge/lang-pt--br-green.svg)](https://github.com/alonsofritz/TIL/tree/master/Go/README.pt-br.md)