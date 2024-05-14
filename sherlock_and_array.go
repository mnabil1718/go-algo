func sum(arr []int32) int32 {
    sum := int32(0)
    for _, val := range arr {
        sum += val   
    }
    
    return sum
}

// https://www.hackerrank.com/challenges/three-month-preparation-kit-sherlock-and-array/problem
func balancedSums(arr []int32) string {
    // Write your code here
    if len(arr) == 1 {
        return "YES"
    } 
    
    if len(arr) == 2 {
        return "NO"
    }
    
    totalSum := sum(arr)
    
    leftSum := int32(0)
    for i := 0; i < len(arr); i++ {
        if leftSum == (totalSum - leftSum - arr[i]) {
            return "YES"
        }
        
        leftSum += arr[i]
    }
    
    return "NO"
}
