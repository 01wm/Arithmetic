package likou

// 最长回文子串

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0 // 记录最长回文子串的起始和结束索引

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // 以当前字符为中心的回文子串长度（奇数长度）
		len2 := expandAroundCenter(s, i, i+1) // 以当前字符及下一个字符之间的位置为中心的回文子串长度（偶数长度）
		length := len1                        // 默认长度为 len1

		if len2 > len1 {
			length = len2 // 如果 len2 更大，则更新为 len2
		}

		if length > end-start { // end-start 可以后面加1，那样会变成只获取第一个最长长度的回文字符串，后面同样长度便不再获取
			// 如果当前回文子串的长度大于已知的最长回文子串长度，则更新最长回文子串的起始和结束索引
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1] // 根据最长回文子串的起始和结束索引，返回最长回文子串
}

func expandAroundCenter(s string, left int, right int) int {
	// 从中心向两边扩展，寻找回文子串的长度
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--  // 向左移动左边界
		right++ // 向右移动右边界
	}
	return right - left - 1 // 返回回文子串的长度
}
