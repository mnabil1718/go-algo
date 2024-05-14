// https://www.hackerrank.com/challenges/three-month-preparation-kit-an-interesting-game-1/problem
func gamingArray(arr []int32) string {
    // Write your code here
    var steps int32
    for len(arr) > 0 {
        var max int32
        var max_position int32
        for idx, val := range arr {
            if val > max {
                max = val
                max_position = int32(idx)
            }
        }
        
        arr = arr[:max_position]  
        steps++ 
    }
    
    if steps % 2 == 0 {
        return "ANDY"
    }
    
    return "BOB"
}
