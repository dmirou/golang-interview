package generics1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapCompare(t *testing.T) {
	assert.True(t, MapCompare(
		map[string]int{},
		map[string]int{}),
	)

	assert.True(t, MapCompare(
		map[string]int{"December": 12, "November": 11},
		map[string]int{"November": 11, "December": 12}),
	)

	assert.False(t, MapCompare(
		map[string]int{"December": 12, "November": 11},
		map[string]int{"November": 12, "December": 11}),
	)

	assert.False(t, MapCompare(
		map[string]int{"December": 12, "November": 11, "September": 9},
		map[string]int{"November": 11, "December": 12}),
	)

	assert.False(t, MapCompare(
		map[string]int{"December": 12, "November": 11},
		map[string]int{"November": 11, "December": 12, "September": 9}),
	)
}

func TestMapApplyFunc(t *testing.T) {
	t.Run("map[string]string", func(t *testing.T) {
		m1 := map[string]string{"Bob": "shy", "Diana": "gorgeous", "John": "sociable"}

		MapApplyFunc(
			m1,
			func(item string) string {
				return item + " (wo)man"
			})

		assert.True(t, MapCompare(
			map[string]string{
				"Bob": "shy (wo)man", "Diana": "gorgeous (wo)man", "John": "sociable (wo)man",
			},
			m1,
		))
	})

	t.Run("map[int]bool", func(t *testing.T) {
		m1 := map[int]bool{2: true, 3: false, 4: true}

		MapApplyFunc(
			m1,
			func(item bool) bool {
				return !item
			})

		assert.True(t, MapCompare(
			map[int]bool{2: false, 3: true, 4: false},
			m1,
		))
	})
}
