// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/

// Counts spaces and corner cases

func mostWordsFound(sentences []string) int {
    maxWords := 0
    for _, sentence := range sentences {
        sentenceWords := 0
        for _, char := range sentence {
            if char == ' ' {
                sentenceWords++
            }
        }
        if sentenceWords > 0 || len(sentence) > 0 {
            sentenceWords++
        }
        if maxWords < sentenceWords {
            maxWords = sentenceWords
        }
    }
    return maxWords
}
