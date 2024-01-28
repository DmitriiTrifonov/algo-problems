// https://leetcode.com/problems/integer-to-roman

func intToRoman(num int) string {
    nums := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
    romans := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}

    out := ""

    for i := 0;  num > 0; i++ {
        for num >= nums[i] {
            out += romans[i]
            num -= nums[i]
        }
    }

    return out
}

// Steps

/*
 1. Define possible ints
 2. Define possible roman letters
 3. Create a result string
 4. Make a cycle for convert numbers
 5. If number consists of multiple values make a cycle for it
 6. Subtract the num value
 7. Break the cycle if the num is zero or below
*/
