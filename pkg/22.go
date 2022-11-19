package pkg

func generateParenthesis(n int) []string {
	var ans []string
	backtrack(&ans, "", 0, 0, n)

	return ans
}

func backtrack(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}

	if open < max {
		backtrack(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		backtrack(ans, cur+")", open, close+1, max)
	}
}
