// https://leetcode.com/problems/number-of-1-bits/

// Simply counting 1 for number
// https://en.wikipedia.org/wiki/Hamming_weight


func hammingWeight(num uint32) int {
    bitsCounter := 0
    for num > 0 {
        if num % 2 != 0 {
            bitsCounter++
            num /= 2
            continue
        }
        num /= 2
    }
    return bitsCounter
}
