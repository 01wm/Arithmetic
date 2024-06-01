package likou

// 两数之和

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//fmt.Printf("%d + %d = %d\n",nums[i],nums[j],nums[i]+nums[j])
			if nums[i]+nums[j] == target {
				/*ret [0]=i
				ret [1]=j*/
				return []int{i, j}
			}
		}
	}
	return nil
}
