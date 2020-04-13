package two_sum

import "testing"
import "github.com/stretchr/testify/assert"

func Test_twoSum(t *testing.T) {
	test := []int{2,7,11,15}
	target := 9
	ret := twoSum(test, target)
	assert.Contains(t, ret, 0, 1)

	test = []int{15,0,1,-1}
	target = 0
	ret = twoSum(test, target)
	assert.Contains(t, ret, 2, 3)

	test = []int{0,3,4,0}
	target = 0
	ret = twoSum(test, target)
	assert.Contains(t, ret, 0, 3)
}
