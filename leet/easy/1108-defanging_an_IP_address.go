// https://leetcode.com/problems/defanging-an-ip-address/

// Concatinating strings

func defangIPaddr(address string) string {
    out := ""

    for _, ch := range address {
        if ch == '.' {
            out += "[.]"
            continue
        }
        out += string(ch)
    }
    
    return out
}
