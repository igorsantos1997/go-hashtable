package main

import "fmt"

const HASHTABLE_SIZE = 256

type hashTableValue[T any] struct {
	key   string
	value T
}
type HashTable[T any] struct {
	values [HASHTABLE_SIZE]LinkedList[hashTableValue[T]]
}

func hashFunction(key string) uint8 {
	total := 0
	for _, char := range key {
		total += int(char)
	}
	// Evita valores negativos e melhora a distribuição
	return uint8(total % HASHTABLE_SIZE)
}

func (hashTable *HashTable[T]) Insert(key string, value T) {
	hashCode := hashFunction(key)
	currentLinkedList := &hashTable.values[hashCode]

	// Verifica se a lista está vazia
	if currentLinkedList.Current.key == "" {
		hashTable.values[hashCode] = LinkedList[hashTableValue[T]]{
			Current: hashTableValue[T]{
				key:   key,
				value: value,
			},
		}
		return
	}

	// Verifica se a chave já existe
	existingLinkedList := getLinkedList(currentLinkedList, key)
	if existingLinkedList != nil {
		// Atualiza o valor existente
		existingLinkedList.Current.value = value
		return
	}

	// Insere novo valor no final da lista
	currentLinkedList.Insert(hashTableValue[T]{
		key:   key,
		value: value,
	})
}

func getLinkedList[T any](linkedList *LinkedList[hashTableValue[T]], key string) *LinkedList[hashTableValue[T]] {
	if linkedList.Current.key == key {
		return linkedList
	}

	if linkedList.Next == nil {
		return nil
	}

	return getLinkedList(linkedList.Next, key)
}

func (hashTable *HashTable[T]) Get(key string) (T, error) {
	hashCode := hashFunction(key)
	linkedList := &hashTable.values[hashCode]
	ll := getLinkedList(linkedList, key)
	if ll == nil {
		var zero T
		return zero, fmt.Errorf("valor não encontrado")
	}
	return ll.Current.value, nil
}

func (hashTable *HashTable[T]) Delete(key string) error {
	hashCode := hashFunction(key)
	linkedList := &hashTable.values[hashCode]

	// Se a lista está vazia
	if linkedList.Current.key == "" {
		return fmt.Errorf("chave não encontrada")
	}

	// Se é o primeiro elemento
	if linkedList.Current.key == key {
		if linkedList.Next == nil {
			// Lista fica vazia
			hashTable.values[hashCode] = LinkedList[hashTableValue[T]]{}
		} else {
			// Move o próximo elemento para a posição atual
			hashTable.values[hashCode] = *linkedList.Next
		}
		return nil
	}

	// Busca e remove o elemento
	success := linkedList.Delete(func(item hashTableValue[T]) bool {
		return item.key == key
	})

	if !success {
		return fmt.Errorf("chave não encontrada")
	}

	return nil
}

func (hashTable *HashTable[T]) Print(key string) {
	hashCode := hashFunction(key)
	fmt.Printf("%s(%d) - ", key, hashCode)

	if value, error := hashTable.Get(key); error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(value)
	}
}

// Size retorna o número total de elementos na hash table
func (hashTable *HashTable[T]) Size() int {
	total := 0
	for i := 0; i < HASHTABLE_SIZE; i++ {
		if hashTable.values[i].Current.key != "" {
			total += hashTable.values[i].Size()
		}
	}
	return total
}

// IsEmpty verifica se a hash table está vazia
func (hashTable *HashTable[T]) IsEmpty() bool {
	return hashTable.Size() == 0
}

// PrintStats imprime estatísticas da hash table
func (hashTable *HashTable[T]) PrintStats() {
	totalElements := hashTable.Size()
	nonEmptyBuckets := 0
	maxBucketSize := 0

	for i := 0; i < HASHTABLE_SIZE; i++ {
		if hashTable.values[i].Current.key != "" {
			nonEmptyBuckets++
			bucketSize := hashTable.values[i].Size()
			if bucketSize > maxBucketSize {
				maxBucketSize = bucketSize
			}
		}
	}

	fmt.Printf("Estatísticas da Hash Table:\n")
	fmt.Printf("  Total de elementos: %d\n", totalElements)
	fmt.Printf("  Buckets não vazios: %d\n", nonEmptyBuckets)
	fmt.Printf("  Tamanho máximo do bucket: %d\n", maxBucketSize)
	fmt.Printf("  Fator de carga: %.2f\n", float64(totalElements)/float64(HASHTABLE_SIZE))
}
