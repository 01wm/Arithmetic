package main

import (
	"01wm/likou"
	"fmt"
	"log"
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

	// 两数之和
	fmt.Println(likou.TwoSum([]int{2, 7, 11, 15}, 9))

	// 两数相加
	l1 := &likou.ListNode{
		Val: 9,
		Next: &likou.ListNode{
			Val: 9,
			Next: &likou.ListNode{
				Val: 9,
				Next: &likou.ListNode{
					Val: 9,
					Next: &likou.ListNode{
						Val: 9,
						Next: &likou.ListNode{
							Val: 9,
							Next: &likou.ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}
	l2 := &likou.ListNode{
		Val: 9,
		Next: &likou.ListNode{
			Val: 9,
			Next: &likou.ListNode{
				Val: 9,
				Next: &likou.ListNode{
					Val: 9,
				},
			},
		},
	}
	req := likou.AddTwoNumbers(l1, l2)
	fmt.Println(req)

	// 无重复字符最长子串
	fmt.Println(likou.LengthOfLongestSubstring("ajdadjkbncl"))

	// 分块校验
	filename := "example.txt" // 要计算校验和的文件名
	hashes, err := likou.ComputeBlockHashes(filename)
	if err != nil {
		log.Fatalf("Failed to compute block hashes: %v", err)
	}
	fmt.Println("Block Hashes:")
	for _, hash := range hashes {
		fmt.Println(hash)
	}

}
