# Implementação simples de uma Hash Table em Go

## Descrição

Esta implementação demonstra uma Hash Table básica em Go utilizando Linked Lists para tratamento de colisões.

## Funcionamento

### Função Hash
Itera sobre cada caractere da chave, obtém seu valor numérico (Unicode), soma esses valores, e determina o índice do array usando o resto da divisão da soma pelo tamanho total do array.

### Armazenamento dos Valores:
Cada posição do array aponta para uma Linked List. Caso ocorra uma colisão, o novo valor é adicionado ao final da lista através do ponteiro Next.
