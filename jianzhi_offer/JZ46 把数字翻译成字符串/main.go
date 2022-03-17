package main

// 当前字符不等于0的时候,dp[i] = dp[i-1],此时将当前位置的一个字符译码
// 当前字符+前一个字符，记为num, 如果 10<=num<=26 此时符合两个合并一起译码的条件；
// 若此时i等于1，直接dp[i]++;
// 大于1, 则dp[i] += dp[i-2];
//
// 举个例子： nums = "324"
// 此时dp[0] = 1, dp[1]呢？ dp[2]呢？
// 很明显nums[1] != '0'，所以dp[1] = dp[0],num = 32，
// 此时不满足两个一起译码的条件则循环往下执行，此时 nums[2] != '0',
// 则 dp[2] = dp[1] = 1, num = 24，此时满足两个一起译码的条件,因为i==2大于1，
// 所以dp[2] += dp[2-2] ,dp[2] = 1+1 = 2

/**
 * 解码
 * @param nums string字符串 数字串
 * @return int整型
 */
func solve(nums string) int {
	// write code here
	if len(nums) == 0 || nums[0] == '0' {
		return 0
	}

	//   表示字符串nums中 以i个位置结尾的前缀字符串的解码种数
	dp := make([]int, len(nums))
	// 第一个字符不为0
	//  由于只有一个值（1-9之间），所以，只有一种编码
	//  在第0个字符的时候只有一个字母所以只有一种翻译办法
	dp[0] = 1

	// 填充剩余dp数组
	for i := 1; i < len(dp); i++ {
		// 当前字符不等于0的时候,dp[i] = dp[i-1]
		// 此时将当前位置的一个字符译码
		if nums[i] != '0' {
			dp[i] = dp[i-1]
		}

		// 当前字符与前面一个字符组成的数值
		num := (nums[i-1]-'0')*10 + nums[i] - '0'
		if num >= 10 && num <= 26 {
			// 第二个字符，特殊处理
			if i == 1 {
				dp[i] += 1
			} else {
				// 由于前面  dp[i] = dp[i-1]
				// 所以此时相当于 dp[i] = dp[i-1] + dp[i-2]
				dp[i] += dp[i-2]
			}
		}
	}

	return dp[len(nums)-1]
}
