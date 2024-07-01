class Solution {
public:
    bool validPartition(vector<int>& nums) {
        int n = nums.size();
        if (n == 1) return false;

        vector<bool> dp(n + 1, false);
        dp[0] = true;

        for (int i = 2; i <= n; ++i) {
            if (i >= 2 && nums[i-1] == nums[i-2]) {
                dp[i] = dp[i] || dp[i-2];
            }
            if (i >= 3 && nums[i-1] == nums[i-2] && nums[i-2] == nums[i-3]) {
                dp[i] = dp[i] || dp[i-3];
            }
            if (i >= 3 && nums[i-1] == nums[i-2] + 1 && nums[i-2] == nums[i-3] + 1) {
                dp[i] = dp[i] || dp[i-3];
            }
        }

        return dp[n];
    }
};
