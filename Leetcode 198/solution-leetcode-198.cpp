class Solution {
public:
    int rob(vector<int>& nums) {
        if (nums.size() == 1) {
            return nums[0];
        }
        int arr1 = nums[0], arr2 = nums[1];
        int maxn = arr1;
        for (int i = 2; i < nums.size(); i++) {
            int sum = nums[i] + max(arr1, maxn);
            maxn = max(maxn, arr2);
            arr1 = arr2;
            arr2 = sum;
        }
        return max(maxn, arr2);
    }
};