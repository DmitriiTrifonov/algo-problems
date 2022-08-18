// https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/

// (x2-x1)^2 + (y2-y1)^2 = R^2

func countPoints(points [][]int, queries [][]int) []int {
    out := make([]int, 0, len(queries))
    // (x2-x1)^2 + (y2-y1)^2 = R^2
    for _, query := range queries {
        x1 := query[0]
        y1 := query[1]
        r := query[2]
        d := r * r
        num := 0
        for _, point := range points {
            x2 := point[0]
            y2 := point[1]
            x := x2 - x1
            y := y2 - y1
            actual := x * x + y * y
            if actual <= d {
                num++
            }
        }
        out = append(out, num)
    }
    return out
}
