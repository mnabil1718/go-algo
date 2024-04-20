func insertSort (arr []int32, n int, desc bool) []int32 {
    for i := 1; i < n; i++ {
        j := i
        
        if desc {
            // Descending
            for j > 0 && arr[j] > arr[j - 1] {
                temp := arr[j]
                arr[j] = arr[j - 1]
                arr[j - 1] = temp
                j--
            }
        } else {
            // Ascending
            for j > 0 && arr[j] < arr[j - 1] {
                temp := arr[j]
                arr[j] = arr[j - 1]
                arr[j - 1] = temp
                j--
            }
        }
    }
    
    return arr
}
