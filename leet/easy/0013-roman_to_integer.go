// https://leetcode.com/problems/roman-to-integer/

// This solution uses letter map and check for corner cases

func romanToInt(s string) int {
    romanMap := map[rune]int {
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    
    sum := 0
    var lastLetter rune
    for _, rl := range s {
        sum += romanMap[rl]
        if lastLetter == 'I' && (rl == 'V' || rl == 'X') {
            sum -= 2
        }
        if lastLetter == 'X' && (rl == 'L' || rl == 'C') {
            sum -= 20
        }
        if lastLetter == 'C' && (rl == 'D' || rl == 'M') {
            sum -= 200
        }
        lastLetter = rl
        
    }
    return sum
}
