package main

import (
	"01wm/likou"
	"fmt"
)

func main() {
	// z变换
	test := likou.Convert("PAYPALISHIRING", 3)
	fmt.Println(test)

	// 寻找两个有序数组的中位数
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6, 7}
	index := likou.FindMedianSortedArrays(nums1, nums2)
	fmt.Println(index)

	// 最长回文子串
	str := "abba"
	longestPal := likou.LongestPalindrome(str)
	fmt.Println("Longest palindrome substring:", longestPal)
}
