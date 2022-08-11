// https://leetcode.com/problems/string-to-integer-atoi/

// All corner cases

func myAtoi(s string) int {
	var out int
    isNegative := false
    isPositive := false
    isDigit := true
    isTrailingZero := false
	  for i, ch := range s {
        if out > (1 << 31) {
            out = (1 << 31)
            break
            } 
        if ch == ' ' {
            if out > 0 {
                break
            }
            if isPositive || isNegative || isTrailingZero {
                break
            }
            continue
        }
        if ch == '+' {
            if isNegative || isPositive {
                break
            }
            isPositive = true
            continue
        }
        if ch == '.' {
            break
        }
        if ch == '-' && i == len(s) - 1 {
            break
        }
        if !isNegative && ch == '-' {
            if isTrailingZero {
                break
            }
            isNegative = true
            continue
        }
        if isNegative && ch == '-' {
            break
        }
        if isPositive && ch == '+' {
            return 0
        }
        if isPositive && isNegative {
            return out
        }
        if i == 0 && ch == '0' {
            isTrailingZero = true
        }
		ch -= '0'
		if ch > 9 && isDigit{
            isDigit = false
            break
		}
		out = out*10 + int(ch)
	}
	if isNegative {
		out = -out
	}
  if out < -(1 << 31) {   
     out = -(1 << 31)
  }
  if out > (1 << 31) - 1 {
     out = (1 << 31) - 1

    } 
	return out
}
