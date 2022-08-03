// https://leetcode.com/problems/defanging-an-ip-address/

// Simply make a rune slice and replace it for changing

func defangIPaddr(address string) string {
    replace := []rune{'[', '.', ']'}
    out := make([]rune, 0, len(address) + 9)
    for _, char := range address {
        if char == '.' {
            out = append(out, replace...)
            continue
        }
        out = append(out, char)
    }
    return string(out)
}
