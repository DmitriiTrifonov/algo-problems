// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

// Two cycles: one for max, one for extra

func kidsWithCandies(candies []int, extraCandies int) []bool {
    out := make([]bool, 0, len(candies))
    max := 0
    for _, c := range candies {
        if c > max {
            max = c
        }
    }
    for _, c := range candies {
        num := c + extraCandies
        if num >= max {
            out = append(out, true)
            continue
        }
        out = append(out, false)
    }
    return out
}
