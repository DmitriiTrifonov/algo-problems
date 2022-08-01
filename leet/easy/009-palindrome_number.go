// https://leetcode.com/problems/palindrome-number/

// This solution return false if number is negative
// The main thing is to reverse number and compare it with original
// If original and reversed are the same the result is good overwise it is bad


func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    var reminder, reverse int
    number := x
    for x != 0 {
        reminder = x % 10
        reverse = reverse * 10 + reminder
        x /= 10
    }
    if reverse != number {
        return false
    }
    return true
}
