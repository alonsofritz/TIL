### Por que usar []rune ao invés de []byte?

Em Go, string é uma sequência de bytes, não necessariamente de caracteres. Um rune é um alias para int32, usado para representar um código Unicode (um caractere). Quando você faz []byte(s), cada posição do slice representa apenas 1 byte. Isso funciona para ASCII, mas quebra para caracteres com acentuação ou emojis, que usam múltiplos bytes (UTF-8). Já []rune(s) garante que cada posição do slice corresponde a um caractere completo (Unicode code point).

```go
s := "Olá 🚀"
fmt.Println(len(s))         // 8 bytes
fmt.Println(len([]rune(s))) // 5 runes (O, l, á, espaço, 🚀)
```

Se você fizesse []byte(s) e invertesse, os bytes do á (0xC3 0xA1) seriam embaralhados e a string ficaria corrompida.
Com []rune(s) isso não acontece, pois cada caractere é tratado como uma unidade.

### Custo Assintótico
A complexidade de tempo para inverter uma string usando runes é O(n), onde n é o número de runes na string. Isso ocorre porque precisamos iterar sobre cada rune para copiá-los em ordem reversa. A complexidade de espaço também é O(n) devido ao armazenamento do slice de runes. Embora haja um custo adicional de conversão de string para runes, isso é necessário para garantir a integridade dos caracteres Unicode, mas não altera a complexidade assintótica geral, sendo ainda O(n).