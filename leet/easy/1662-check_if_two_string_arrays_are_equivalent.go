// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

// Make two strings and compare them

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    w1 := ""
    w2 := ""
    
    for _, p := range word1 {
        w1 += p
    }
    
    for _, p := range word2 {
        w2 += p
    }
    
    if w1 == w2 {
        return true
    }
    return false
}
