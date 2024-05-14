// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-misere-nim-1/problem
func misereNim(s []int32) string {
    // Write your code here
    var min int32
    for _, val := range s {
        if val > min {
            min = val
        }
    }
    
    if min == 1 {
        if len(s) % 2 == 0 {
            return "First"
        } else {
            return "Second"
        }
    }
    
    xor_result := s[0]
    
    for i := 1; i < len(s); i++ {
        xor_result ^= s[i]
    }
    
    if xor_result == 0 {
        return "Second"
    } 
    
    return "First"

}
