package delete

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testInput struct {
	cache map[string]string
	key   string
}

func TestDelete(t *testing.T) {
	testCases := map[string]struct {
		input       testInput
		expectedErr bool
		err         error
	}{
		"key not found returns error": {
			input: testInput{
				cache: make(map[string]string),
				key:   "testNotFoundKey",
			},
			expectedErr: true,
			err:         ErrKeyNotFound,
		},
		"key found deletes key": {
			input: testInput{
				cache: map[string]string{
					"testKey": "testValue",
				},
				key: "testKey",
			},
			expectedErr: false,
			err:         nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := Delete(tc.input.cache, tc.input.key)

			if tc.expectedErr {
				assert.Error(t, err)
				assert.Equal(t, tc.err, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
