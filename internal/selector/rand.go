package selector

import "math/rand/v2"

// WeightedRand selector picks elements randomly based on their weights
type WeightedRand[T any] struct {
	values      []T
	weights     []int
	totalWeight int
}

func NewWeightedRandSelector[T any](values []T, weights []int) *WeightedRand[T] {
	wrs := &WeightedRand[T]{
		values:      make([]T, len(values)),
		weights:     make([]int, len(weights)),
		totalWeight: 0,
	}
	// copy the underlying array values as we're going to modify content of slices
	copy(wrs.values, values)
	copy(wrs.weights, weights)

	for _, w := range weights {
		wrs.totalWeight += w
	}

	return wrs
}

// Pick returns randomly chose element from values based on its weight if any exists
func (wrs *WeightedRand[T]) Pick() T {
	var defaultVal T
	if len(wrs.values) == 0 {
		return defaultVal
	}

	rNum := rand.IntN(wrs.totalWeight) + 1

	sum := 0
	for i := range len(wrs.values) {
		sum += wrs.weights[i]
		if sum >= rNum {
			wrs.totalWeight -= wrs.weights[i]
			result := wrs.values[i]

			// remove picked element and its weight
			wrs.values[i] = wrs.values[len(wrs.values)-1]
			wrs.values = wrs.values[:len(wrs.values)-1]
			wrs.weights[i] = wrs.weights[len(wrs.weights)-1]
			wrs.weights = wrs.weights[:len(wrs.weights)-1]
			return result
		}
	}

	return defaultVal
}
