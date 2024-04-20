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

func maximumPerimeterTriangle(sticks []int32) []int32 {
    // Write your code here
    var sortedArray []int32 = insertSort(sticks, len(sticks), false)
    
    
   for i := len(sortedArray) - 1; i >= 2; i-- {
       if sortedArray[i] < sortedArray[i - 1] + sortedArray[i - 2] {
           return []int32{sortedArray[i - 2], sortedArray[i - 1], sortedArray[i]}
       }
   }
    return []int32{-1}

}
