package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		// Calcula o complemento necessário para atingir o alvo
		complement := target - num
		// Verifica se o complemento já foi visto
		if index, ok := seen[complement]; ok {
			return []int{index, i}
		}
		// Adiciona o número atual e seu índice ao mapa
		seen[num] = i
	}
	// Se nenhum par foi encontrado, retorna um slice vazio
	return []int{}
}
