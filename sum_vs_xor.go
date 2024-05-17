// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-sum-vs-xor/problem
func sumXor(n int64) *big.Int {
    var count int64
    for n > 0 {
        if n%2 == 0 {
            count++
        }
        n /= 2
    }
    
    result := big.NewInt(1)
    result.Lsh(result, uint(count))
    return result
}
