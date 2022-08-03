// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

// Make a map with values and simply add them

func finalValueAfterOperations(operations []string) int {
    x := 0
    opMap := map[string]int{
        "--X": -1,
        "X--": -1,
        "++X": 1,
        "X++": 1,
    }
    for _, op := range operations {
        x += opMap[op]
    }
    return x
}
