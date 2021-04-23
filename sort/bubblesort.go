package sort

import (
	"fmt"
)

//氣泡排序法
func bubblesort(nums []int) []int {
	//穩定
	//交換
	//時間複雜度 最差 Ο(n2)
	//空間複雜度 Ο(1)
	//有flag方式

	for i := 0; i < len(nums) - 1; i++ {
		flag := false
		for j := 0; j < len(nums) - i - 1; j++ {
			fmt.Println(nums, i, j)
			if nums[j] > nums[j+1] { //大於換到前面 小於換到後面
				flag = true
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
		if !flag {
			break;
		}
	}
	return nums
}
