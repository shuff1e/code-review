package array

// 给定2个字符串，返回两个字符串的最长公共子串
// 例如给定 123456 和 2345
// 返回2345

// dp[i][j]表示以str1[i]，str2[j]为结尾的最长字串的长度

//              dp[i-1][j-1] + 1 ,str1[i] == str2[j]
// dp[i][j] =   0

func getDP3(str1,str2 string) ([][]int,int,int,int) {
	dp := make([][]int,len(str1))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(str2))
	}

	for i := 0;i<len(str1);i++ {
		if str1[i] == str2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}

	for j := 0;j<len(str2);j++ {
		if str2[j] == str1[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	maxDP := 0
	maxI := 0
	maxJ := 0

	for i := 1;i<len(str1);i++ {
		for j:=1;j<len(str2);j++ {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxDP {
					maxDP = dp[i][j]
					maxI = i
					maxJ = j
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	//fmt.Printf("%#v\n",dp)

	return dp,maxDP,maxI,maxJ
}

func Lcss(str1,str2 string) string {
	_,maxDP,maxI,_ := getDP3(str1,str2)
	return str1[maxI-maxDP+1:maxI+1]
}