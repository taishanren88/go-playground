package main

import "fmt"

func uniquePaths(m int, n int) int {
        // We can only go right or down, which means we can come from the left or come from the top
        // Create a matrix : dp [m][n] and set dp[0][0] (start point) == 1, 
        // Edge cases: dp[0][j] = 1 since , dp [i][0] = 1
        // For the rest : At each point dp[i][j] = dp[i-1][j] + dp[i][j-1]
        // Return dp [m][n]
        if m == 0  || n == 0{
        	return 0
        }
        dp := make([][]int, m)
        for i := range dp {
        	dp[i] = make([]int, n)
        	dp[i][0] = 1
        }
        for j := range dp[0] {
        	dp[0][j] = 1
        }
        dp[0][0] = 1
        for i:= 1; i < m; i++ {
        	for j:= 1; j <n ; j++ {
        		dp[i][j] = dp[i-1][j] + dp[i][j-1];
        	}
        }

        return dp[m-1][n-1];
}

func main () {
	fmt.Println("hello World")
	var e int = 10;
	e = 10;

	x:=15
	fmt.Println(e);
	fmt.Println(x);
	// fmt.Println(uniquePaths(7, 3))
}