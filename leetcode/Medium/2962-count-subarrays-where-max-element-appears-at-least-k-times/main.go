package main

import (
	"fmt"
)

func countSubarrays(nums []int, k int) int64 {
	// Passo 1: Encontrar o valor máximo no array
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	var result int64 = 0
	count := 0 // Contagem de maxNum na janela atual
	left := 0  // Ponteiro esquerdo da janela

	// Passo 2: Percorrer o array com janela deslizante
	for right := 0; right < len(nums); right++ {
		if nums[right] == maxNum {
			count++
		}

		// Passo 3: Enquanto a janela tem pelo menos k ocorrências do máximo
		for count >= k {
			// Todos os elementos daqui até o final do array formam subarrays válidos
			result += int64(len(nums) - right)

			// Mover o ponteiro left e ajustar a contagem
			if nums[left] == maxNum {
				count--
			}
			left++
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, 2, 3, 3}
	k := 2
	fmt.Println(countSubarrays(nums, k))
	nums = []int{1, 4, 2, 1}
	k = 3
	fmt.Println(countSubarrays(nums, k))
}
