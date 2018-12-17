package dice

import (
	"math/rand"
)

// RollN takes a count of dice and the number of sides per dice and rolls them.
func RollN(count, sides int) (rolls []int) {
	for i := 0; i < count; i++ {
		rolls = append(rolls, Roll(sides))
	}
	return rolls
}

func Roll(sides int) int {
	return int(rand.Int31n(int32(sides))+1)
}
