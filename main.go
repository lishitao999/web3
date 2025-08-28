package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	m := []int{1, 3, 2, 3, 2, 4, 1}
	fmt.Println("只出现一次的数字：", singleNumber(m))
	i := 123321
	fmt.Println("回文数判断：", isPalindrome(i))
	s := "({{[]}})"
	fmt.Println("字符串是否有效判断：", isValid(s))
	arr := []string{"flower", "flow", "flight"}
	fmt.Println("最长公共前缀：", longestCommonPrefix(arr))
	nums := []int{1, 2, 3}
	fmt.Println("加一：", plusOne(nums))
	nums2 := []int{1, 1, 3}
	fmt.Println("删除有序数组中的重复项：", removeDuplicates(nums2))
	num3 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("合并区间：", merge(num3))
	num4 := []int{2, 6, 3, 4}
	target := 8
	fmt.Println("两数之和：", twoSum(num4, target))

}

//=============流程控制=======================================================

/*
*1.只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
*                  找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
*			       例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*
*       		 https://leetcode.cn/problems/single-number/description/
*               给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，
*    			其余每个元素均出现两次。找出那个只出现了一次的元素。
 */
func singleNumber(nums []int) int {
	times := map[int]int{}

	for _, value := range nums {
		times[value]++
	}

	for k, v := range times {
		if v == 1 {
			return k
		}
	}
	return 0

}

/*
*2.回文数：判断一个整数是否是回文数
*
*		  https://leetcode.cn/problems/palindrome-number/description/
*		  给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
*		  回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*		  例如，121 是回文，而 123 不是。
 */
func isPalindrome(i int) bool {
	var flag = false

	//负数和最后一位是0的都不是回文数
	if i < 0 || (i%10 == 0 && i != 0) {
		return flag
	}

	//先转换为字符串
	str := strconv.Itoa(i)
	//对字符串反向遍历
	var rever_str string
	for i := len(str) - 1; i >= 0; i-- {
		rever_str = rever_str + string(str[i])
	}

	rever_num, _ := strconv.Atoi(rever_str)

	return i == rever_num
}

//=============字符串=======================================================

/*3.有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*
*	            https://leetcode.cn/problems/valid-parentheses/description/
*      		   给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
*		       有效字符串需满足：
*				    1.左括号必须用相同类型的右括号闭合。
*				    2.左括号必须以正确的顺序闭合。
*					3.每个右括号都有一个对应的相同类型的左括号。
 */
func isValid(s string) bool {
	//如果字符串的个数是奇数则直接返回false
	if len(s)%2 == 1 {
		return false
	}

	//定义切片
	slice := []rune{}

	//定义校验map
	mapVlid := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	//遍历字符串进行校验
	for _, value := range s {
		switch value {
		case '(', '{', '[':
			slice = append(slice, value)
		case ')', '}', ']':
			// slice: "{[("  }  取切片的最后一位
			if len(slice) == 0 || mapVlid[value] != slice[len(slice)-1] {
				return false
			}
			//如果括号匹配，则截取切片0 ~ len(slice)-2
			slice = slice[:len(slice)-1]

		}
	}

	return len(slice) == 0

}

/*4.最长公共前缀:：查找字符串数组中的最长公共前缀。
*
*                https://leetcode.cn/problems/longest-common-prefix/description/
*			  	 编写一个函数来查找字符串数组中的最长公共前缀。
*                如果不存在公共前缀，返回空字符串 ""。
 */
func longestCommonPrefix(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}

	//取数组的第一个字符串，分别与数组中的后面字符串匹配
	strFirst := strArr[0]
	for i := 1; i < len(strArr); i++ {
		j := 0
		for ; j < len(strFirst) && j < len(strArr[i]); j++ {
			v1 := strFirst[j]
			v2 := strArr[i][j]

			if v1 != v2 {
				break
			}
		}

		//截取匹配上的子字符串
		strFirst = strFirst[:j]

		if len(strFirst) == 0 { //没有找到
			return ""
		}
	}

	return strFirst

}

//=============基本值类型=======================================================

/*5.加一：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*
*        https://leetcode.cn/problems/plus-one/description/
*		 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
*		 这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
*	    将大整数加 1，并返回结果的数字数组。
 */
func plusOne(digits []int) []int {
	//倒序遍历数组
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		//等于9，+1变成0
		digits[i] = 0
	}

	//所有元素都是9，有进位的情况
	digits = make([]int, n+1)
	digits[0] = 1
	return digits

}

//============= 引用类型：切片=======================================================

/*6.删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，
*					     使每个元素只出现一次，返回删除后数组的新长度。
*						 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*						 可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
*						 当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*
*            			https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
*                       给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
*                		返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
*
*                       考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
*								更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，
*								并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
*						    	返回 k 。
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 1 //慢指针
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}

	}

	return j
}

/*7.合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
*            请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*		     可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
*
*		   https://leetcode.cn/problems/merge-intervals/description/
*		   以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
*		   请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*
*		   示例 1：
*			输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
*			输出：[[1,6],[8,10],[15,18]]
*			解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*
*		   示例 2：
*			输入：intervals = [[1,4],[4,5]]
*			输出：[[1,5]]
*		    解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 */
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//取二维数组的第一个元素
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intervals[i][0] <= last[1] { //有重叠
			last[1] = max(last[1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}

	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//============= 基础=======================================================

/*8.两数之和：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*
*			https://leetcode.cn/problems/two-sum/description/
*		    给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，
*			并返回它们的数组下标。
*
*			你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
*			你可以按任意顺序返回答案。
*
*			示例 1：
*			输入：nums = [2,7,11,15], target = 9
*			输出：[0,1]
*			解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*
*			示例 2：
*			输入：nums = [3,2,4], target = 6
*			输出：[1,2]
*
*			示例 3：
*			输入：nums = [3,3], target = 6
*			输出：[0,1]
 */

func twoSum(nums []int, target int) []int {
	for i, value := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+value == target {
				return []int{i, j}
			}
		}
	}

	return nil

}
