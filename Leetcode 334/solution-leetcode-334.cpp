class Solution {
public:
    bool increasingTriplet(vector<int>& nums) {
        int length = nums.size();

        int a = INT_MAX;
        int b = nums[0];
        for(int i = 1; i < length; i++){
            if(nums[i] > a) return true;
            if(b < nums[i]){
                a = min(a, nums[i]);
            }
            b = min(b, nums[i]);
        }
        return false;
    }
};
