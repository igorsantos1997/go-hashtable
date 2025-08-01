package main

import "fmt"

func main() {
	table := HashTable[uint32]{}

	fmt.Println("=== Testando Hash Table ===")

	// Teste de inserção
	table.Insert("abc", 123)
	table.Insert("acb", 456)
	table.Insert("cba", 789)
	table.Insert("bca", 1011)
	table.Insert("nãoColide", 10)

	// Teste de busca
	fmt.Println("\n--- Teste de Busca ---")
	table.Print("abc")
	table.Print("acb")
	table.Print("cba")
	table.Print("bca")
	table.Print("nãoColide")
	table.Print("teste") // Deve retornar erro

	// Teste de atualização
	fmt.Println("\n--- Teste de Atualização ---")
	table.Insert("abc", 111)
	table.Insert("cba", 222)
	table.Print("abc")
	table.Print("cba")

	// Teste de estatísticas
	fmt.Println("\n--- Estatísticas ---")
	table.PrintStats()

	// Teste de remoção
	fmt.Println("\n--- Teste de Remoção ---")
	fmt.Println("Removendo 'abc':")
	err := table.Delete("abc")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Println("'abc' removido com sucesso")
	}

	table.Print("abc") // Deve retornar erro

	// Estatísticas após remoção
	fmt.Println("\n--- Estatísticas Após Remoção ---")
	table.PrintStats()

	// Teste de chaves que causam colisão
	fmt.Println("\n--- Teste de Colisões ---")
	table.Insert("ab", 100)
	table.Insert("ba", 200)
	table.Print("ab")
	table.Print("ba")

	// Verificar se estão no mesmo bucket
	hashAb := hashFunction("ab")
	hashBa := hashFunction("ba")
	fmt.Printf("Hash de 'ab': %d\n", hashAb)
	fmt.Printf("Hash de 'ba': %d\n", hashBa)
	fmt.Printf("Colisão: %t\n", hashAb == hashBa)
}
