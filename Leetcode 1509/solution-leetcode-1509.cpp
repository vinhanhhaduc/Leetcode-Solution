class Solution {
public:
    int minDifference(vector<int>& nums) {
        int length = nums.size();
        if (length < 5) return 0;
        sort(nums.begin(), nums.end());
        int j = length - 4;
        int i = 0;
        int res = INT_MAX;
        while (j  < length) {
            res = min(res, nums[j] - nums[i]);
            j++;
            i++;
        }
        return res;
    }
};