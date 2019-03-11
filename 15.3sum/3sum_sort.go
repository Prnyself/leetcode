package main

import "sort"

func threeSumSort(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums) // 先排序，默认快排nlogn
	n := len(nums) - 1
	for k, v := range nums {
		if k+2 == len(nums) { // 防止下标超长
			break
		}
		if k > 0 && v == nums[k-1] {
			continue // 如果v和之前一个不变，则跳过下一个
		}
		start, end := k+1, n // 初始化两头下标
		for start < end {
			switch tmp := nums[start] + nums[end] + v; {
			case tmp == 0:
				res = append(res, []int{v, nums[start], nums[end]})
				for nums[start-1] == nums[start] && start < end {
					start++
				}
				// 可优化如下：如果相等则继续挪。两头都挪，因为如果只一边变了另一边没变那肯定不可能继续相等了。
				//start++
				//for nums[start-1] == nums[start] && start < end {
				//	start++
				//}
				//end--
				//for nums[end+1] == nums[end] && start < end {
				//	end--
				//}
			case tmp < 0:
				start++
				// 可优化如下：如果相等继续挪。
				//for nums[start-1] == nums[start] && start < end {
				//	start++
				//}
			case tmp > 0:
				end--
				// 可优化如下：如果相等继续挪。
				//for nums[end+1] == nums[end] && start < end {
				//	end--
				//}
			}
		}
	}
	return res
}

func main() {

}
