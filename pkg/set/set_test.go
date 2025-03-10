package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testInput struct {
	cache map[string]string
	key   string
	value string
}

func TestSet(t *testing.T) {
	testCases := map[string]struct {
		input  testInput
		expect string
	}{
		"set key": {
			input: testInput{
				cache: make(map[string]string),
				key:   "testKey",
				value: "testValue",
			},
			expect: "testValue",
		},
		"update key": {
			input: testInput{
				cache: map[string]string{
					"testKey": "testValue",
				},
				key:   "testKey",
				value: "updatedValue",
			},
			expect: "updatedValue",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			value := Set(tc.input.cache, tc.input.key, tc.input.value)

			assert.Equal(t, tc.expect, value)
		})
	}
}
