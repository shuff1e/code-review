package main

func longestPalindrome(s string) string {

	dp := make([][]int,len(s))
	for i :=0 ;i<len(dp);i++ {
		dp[i] = make([]int,len(s))
	}

	for i:=0;i<len(s);i++ {
		dp[i][i] = 1
	}

	for  i := 0;i<len(s)-1;i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 0
		}
	}

	for length:=2;length<len(s);length++ {
		for i := 0;i<len(s);i++ {
			if i + length < len(s) {
				if s[i] == s[i+length] && dp[i+1][i+length-1] > 0 {
					dp[i][i+length] = dp[i+1][i+length-1] + 2
				}
			}
		}
	}

	max := 0
	maxStr := ""
	for i:=0;i<len(s);i++ {
		for j := i;j<len(s);j++ {
			if dp[i][j] > max {
				max = dp[i][j]
				maxStr = s[i:j+1]
			}
		}
	}

	return  maxStr

}
