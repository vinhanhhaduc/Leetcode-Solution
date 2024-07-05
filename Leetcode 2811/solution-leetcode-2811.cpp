class Solution {
public:
    bool canSplitArray(vector<int>& nums, int m) {
        if (nums.size() - 1 <= 1) return true;
        for(int i = 0;i < nums.size()-1; i++){
            if(nums[i] + nums[i + 1] >= m) {
                return true;
            }
        }
        return false;
    }
};