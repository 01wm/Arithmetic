package likou

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var (
//		this1 = l1
//		this2 = l2
//		adds = 0
//		rel = &ListNode{}
//		relthis = &ListNode{}
//	)
//	rel = nil
//	relthis = rel
//	for{
//		br:=false
//		// 1 != nil 2 != nil
//		// 1 == nil 2 != nil
//		// 1 != nil 2 ==nil
//		nums:=0
//		if  this1 != nil {
//			nums = this1.Val
//		} else {
//			br = true
//		}
//		if this2 != nil {
//			nums += this2.Val
//			br = false
//		}
//		nums += adds
//		if nums>=10 {
//			adds = 1
//			nums = nums/1%10
//		} else {
//			adds = 0
//		}
//		if  br {
//			if nums!=0 {
//				relthis.Next = &ListNode{
//					Val: 1,
//				}
//			}
//			break
//		}
//		if  rel == nil{
//			rel = &ListNode{
//				Val: nums,
//			}
//			relthis = rel
//		} else{
//			relthis.Next = &ListNode{
//				Val: nums,
//			}
//			relthis = relthis.Next
//		}
//
//		if this1 != nil {
//			this1 = this1.Next
//		}
//		if this2 != nil {
//			this2 = this2.Next
//		}
//	}
//	return rel
//}

// 两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	suns := func(l *ListNode) int64 {
		var (
			i     = int64(0)
			this  = l
			l1num = int64(0)
		)
		for {
			// this.Val * (10 ^ i)

			l1num += int64(this.Val * int(math.Pow(10, float64(i))))
			i++
			if this.Next == nil {
				break
			} else {
				this = this.Next
			}
		}
		return l1num
	}
	sum1, sum2 := suns(l1), suns(l2)
	res := sum1 + sum2
	i := 0
	rel := &ListNode{}
	this := rel
	var fast *ListNode = nil
	for {
		//n/10^(i)%10
		pw := int64(math.Pow(10, float64(i)))
		if pw > res {
			if fast != nil {
				fast.Next = nil
			}
			break
		}
		i++
		this.Val = int(res / pw % 10)
		this.Next = &ListNode{}
		fast = this
		this = this.Next

	}
	return rel
	var (
		x     = 0
		that  = l2
		l2num = 0
	)
	for {
		l2num += that.Val * int(math.Pow(10, float64(x)))
		x++
		if that.Next == nil {
			break
		} else {
			that = that.Next
		}
	}
	return nil
}
