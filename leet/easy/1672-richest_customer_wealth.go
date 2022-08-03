// https://leetcode.com/problems/richest-customer-wealth/

// Two arrays O(n)

func maximumWealth(accounts [][]int) int {
    max := 0
    for _, account := range accounts {
        accountWealth := 0
        for _, num := range account {
            accountWealth += num
        }
        if accountWealth > max {
            max = accountWealth
        }
    }
    return max
}
