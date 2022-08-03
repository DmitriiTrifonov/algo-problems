// https://leetcode.com/problems/decode-xored-array/

// XOR original i - 1 and encoded i

func decode(encoded []int, first int) []int {
    out := make([]int, len(encoded) + 1)
    out[0] = first
    for i := range encoded {
        out[i + 1] = out[i] ^ encoded[i]
    }
    return out
}
