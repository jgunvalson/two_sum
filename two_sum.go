package two_sum

import log "github.com/sirupsen/logrus"

func twoSum(nums []int, target int) []int {
	valMap := make(map[int]int)
	for index, value := range nums {
		log.Infof("the left value in the sum is %d at index %d", value, index)
		right := target - value
		log.Infof("the right value needs to equal %d", right)
		if i, ok := valMap[right]; ok {
			log.Infof("the right value already exists in the map, found solution [%d, %d] at indices [%d, %d]", value, valMap[right], index, i)
			return []int{index, i}
		}
		valMap[value] = index
	}
	return nil
}
