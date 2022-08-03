// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

// Max sets every step

func kidsWithCandies(candies []int, extraCandies int) []bool {
    out := make([]bool, 0, len(candies))
    max := 0
    for _, c := range candies {
        if c + extraCandies > max {
            max = c + extraCandies
            out = append(out, true)
            continue
        }
        out = append(out, false)
        max = c + extraCandies
    }
    return out
}
