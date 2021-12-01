package main

const (
    maxNum = 1000001
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    
    if m == 0 && n == 0{
        return 0
    }
    
    if m == 0 {
        return findMedianSortedArray(nums2)
    }
    
    if n == 0 {
        return findMedianSortedArray(nums1)
    }
    
    idx1 := 0
    idx2 := 0
    mn := m + n
    mid := mn / 2

    // store the latest 2 values
    vArr := make([]int, 2)
    v1 := nums1[0]
    v2 := nums2[0]
    v := v1

    for i := 0; i <= mid; i++ {
        if v1 > v2 {
            v = v2
            idx2++
            if idx2 >= n {
                v2 = maxNum
            } else {
                v2 = nums2[idx2]
            }
        } else {
            v = v1
            idx1++
            if idx1 >= m {
                v1 = maxNum
            } else {
                v1 = nums1[idx1]
            }
        }

        vArr[0] = vArr[1]
        vArr[1] = v
    }
    
    if mn % 2 != 0 {
        return float64(vArr[1])
    }

    return (float64(vArr[0]) + float64(vArr[1])) / 2
}

func findMedianSortedArray(nums []int) float64 {
    n := len(nums)
    if n == 1 {
        return float64(nums[0])
    } else if n % 2 != 0 {
        return float64(nums[n/2])
    }

    return (float64(nums[n/2 - 1]) + float64(nums[n/2])) / 2
}
