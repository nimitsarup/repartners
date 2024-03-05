package handlers

import (
	"math"
	"sort"

	"github.com/nimitsarup/rep/db"
)

type Handlers struct {
	DB db.PacksInMemoryDB
}

//go:generate moq -out mock/HandlersInterface.go -pkg mock . HandlersInterface
type HandlersInterface interface {
	UpdatePacks(packs []int) error
	GetPacksForItems(items int) map[int]int
}

// returns http-status + error (if any)
func (h *Handlers) UpdatePacks(packs []int) error {
	return h.DB.UpdatePacks(packs)
}

func (h *Handlers) GetPacksForItems(items int) map[int]int {
	return calculateBestPackCombination(items, h.DB.GetPackSizes())
}

/*
fulfill the order starting with the largest packs first
ensuring that we send out least items and use as few packs as possible
*/
func calculateBestPackCombination(order int, packSizes []int) map[int]int {
	// try larger packs first
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	// map to keep track of the best pack sizes combo
	var bestCombination map[int]int
	minPacks := math.MaxInt
	minExcess := math.MaxInt

	// Recursive - to explore all combinations
	var findCombinations func(int, int, map[int]int, int)
	findCombinations = func(index, currentOrder int, currentCombination map[int]int, currentPacks int) {
		if currentOrder >= order {
			excess := currentOrder - order
			// first priority - minimum excess
			// second priority - minimum packs
			// recurse and pick the best
			if excess < minExcess || (excess == minExcess && currentPacks < minPacks) {
				// good/better candidate solution - store it till we get the best
				minExcess = excess
				minPacks = currentPacks
				bestCombination = make(map[int]int)
				for k, v := range currentCombination {
					bestCombination[k] = v
				}
			}
			return
		}

		if index == len(packSizes) {
			return
		}

		// Skip current pack size
		findCombinations(index+1, currentOrder, currentCombination, currentPacks)

		// Include current pack size
		newCombination := make(map[int]int)
		for k, v := range currentCombination {
			newCombination[k] = v
		}
		newCombination[packSizes[index]]++
		findCombinations(index, currentOrder+packSizes[index], newCombination, currentPacks+1)
	}

	// kick off with empty values
	findCombinations(0, 0, make(map[int]int), 0)

	return bestCombination
}
