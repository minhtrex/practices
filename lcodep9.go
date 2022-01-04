package main

func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    
    y := 0
    for {
        y = (y * 10) + (x % 10)
        x /= 10
        if x <= y {
            break
        }
    }

    return y == x || y / 10 == x
}
