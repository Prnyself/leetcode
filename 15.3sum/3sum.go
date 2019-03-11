package main

import "fmt"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sumMap := make(map[int][][]int)
	for k, v := range nums {
		if _, ok := sumMap[v]; ok {
			continue
		}
		if _, ok := sumMap[0-v]; ok {
			continue
		}
		if k+3 > len(nums) {
			break
		}
		sumMap[v] = twoSum(nums[k+1:], 0-v, sumMap)
	}
	for k, v := range sumMap {
		if len(v) == 0 {
			continue
		}
		tmp := [][]int{}
		for _, vv := range v {
			vv = append(vv, k)
			tmp = append(tmp, vv)
		}
		res = append(res, tmp...)
	}
	return res
}

func twoSum(nums []int, sum int, sumMap map[int][][]int) (res [][]int) {
	if len(nums) < 2 {
		return
	}
	tmp := make(map[int]int, len(nums))
	resMap := make(map[int]int, len(nums))
	for _, v := range nums {
		if mapV, ok := sumMap[v]; ok && len(mapV) > 0 {
			continue
		}
		if _, ok := tmp[v]; !ok {
			tmp[sum-v] = v
		} else {
			resMap[v] = sum - v
		}
	}
	for k, v := range resMap {
		res = append(res, []int{k, v})
	}
	return
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	//target := 0
	fmt.Println(threeSum(nums))
}
