class Solution {
public:
    int maxArea(vector<int>& height) {
        int l = 0;
        int r = height.size() - 1;
        int maxW = 0;
        while (l < r) {
            int w = r - l;
            int h = min(height[l], height[r]);
            int c = w * h;
            maxW = max(maxW, c);
            if (height[l] < height[r]) {
                l++;
            } else {
                r--;
            }
        }
        return maxW;
    }
};
