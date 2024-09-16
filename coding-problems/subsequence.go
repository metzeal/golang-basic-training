package main

import "fmt"

func solution(a string, b string) int {
    count:=0
  
     i,j:=0,0

	 for i<len(a) && j <len(b){
          if (a[i]==b[j]){
			count++
			i++
			j++
		  }else{
			if (len(a)-i>len(b)-j){
				i++
			}else{
				j++
			}
		  }
	 }
return count          
}
func main() {
	// result:=solution("abcdef","ace")
	// fmt.Print(result)

	// result=solution("ace","ace")
	// fmt.Print(result)

	// result=solution("ace","ad")
	// fmt.Print(result)

	// result=solution("ace","pqr")
	// fmt.Print(result)

	// result=solution(" ","babbab")
	// fmt.Print(result)

	// result:=solution("abc","bac")
	// fmt.Print(result)
	testLCS()
}
func lcs(a, b string) int {
	n := len(a)
	m := len(b)

	// Create a 2D dp array to store the lengths of LCS for substrings
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Fill dp array using dynamic programming
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// The value in dp[n][m] will be the length of the LCS
	return dp[n][m]
}

func testLCS() {
	testCases := []struct {
		a, b   string
		result int
	}{
		{"abc", "abc", 3},                     // Case 1: Completely matching strings
		{"abc", "def", 0},                     // Case 2: No common subsequence
		{"abc", "bac", 2},                     // Case 3: Partial match
		{"axbxc", "abc", 3},                   // Case 4: Subsequence in the middle
		{"abcdef", "acf", 3},                  // Case 5: One string is a subsequence of the other
		{"", "abc", 0},                        // Case 6: One string is empty
		{"abcdef", "ab", 2},                   // Case 7: Different length strings
		{"aggtab", "gxtxayb", 4},              // Case 8: Subsequence spread across
		{"aaa", "aa", 2},                      // Case 9: Repeated characters
		{"abcdefghijabcdefghij", "klmnopqrstklmnopqrstu", 0}, // Case 10: Very long strings
	}

	for _, tc := range testCases {
		result := lcs(tc.a, tc.b)
		fmt.Printf("LCS of \"%s\" and \"%s\" = %d (Expected: %d)\n", tc.a, tc.b, result, tc.result)
	}
}