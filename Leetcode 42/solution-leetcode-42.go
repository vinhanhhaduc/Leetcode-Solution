func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    left, right := 0, len(height) - 1
    leftM, rightM := height[left], height[right]
    water := 0

    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftM {
                leftM = height[left]
            } else {
                water += leftM - height[left]
            }
            left++
        } else {
            if height[right] >= rightM {
                rightM = height[right]
            } else {
                water += rightM - height[right]
            }
            right--
        }
    }

    return water
}