package main

import "strconv"
import "fmt"
import "strings"

func dfs(num string, target int, position int, runningSum int64, runningString string,
	multiplyAfterLastAdd int64, results *[]string) {

	if position == len(num) {
		// fmt.Printf("%s , %d\n", runningString, runningSum)
		if runningSum == int64(target) {
			*results = append(*results, runningString)
		}
		return
	}

	for i := position; i < len(num); i++ {
		if i > position && num[position] == '0' {
			break
		}
		strPartial := num[position : i+1]
		a, _ := strconv.Atoi(strPartial)
		strPartialAsInt := int64(a)

		if len(runningString) == 0 {
			dfs(num, target, i+1, runningSum+strPartialAsInt, strPartial, strPartialAsInt, results)

		} else {
			dfs(num, target, i+1, runningSum+strPartialAsInt, runningString+"+"+strPartial, strPartialAsInt, results)
			dfs(num, target, i+1, runningSum-strPartialAsInt, runningString+"-"+strPartial, -strPartialAsInt, results)
			dfs(num, target, i+1, runningSum-multiplyAfterLastAdd+multiplyAfterLastAdd*strPartialAsInt,
				runningString+"*"+strPartial, multiplyAfterLastAdd*strPartialAsInt, results)
		}
	}

}
func addOperators(num string, target int) []string {
	// Similar to Combination sum DFS
	// Loop through the entire string
	// Create numbers at each point , and either add , subtract, or multiply to the running value and also the running string value
	// If we reach end , check if teh running value is equal to the target
	// We are dealing with '*', so need to handle priority of multiply also , so keep track of multiplyAfterAdd
	// At each point of multipling , you remove the multiplyAfterAdd from running sum, recalculate multiplyAfterAddand then add it back
	var results []string
	dfs(num, target, 0, int64(0), "", int64(0), &results)
	return results
}

func main() {
	// should
	results := addOperators("12", 3)
	fmt.Println(strings.Join(results, ", "))

	// should 1+2+3, 1*2*3
	results = addOperators("123", 6)
	fmt.Println(strings.Join(results, ", "))
	// should 2+3*2, 2*3+2
	results = addOperators("232", 8)
	fmt.Println(strings.Join(results, ", "))

	// should 1*0+5, 10 -5
	results = addOperators("105", 5)
	fmt.Println(strings.Join(results, ", "))

	// Should print 0+0, 0 -0, 0*0
	results = addOperators("00", 0)
	fmt.Println(strings.Join(results, ", "))

	//Should be empty
	results = addOperators("3456237490", 9191)
	fmt.Println(strings.Join(results, ", "))
}
