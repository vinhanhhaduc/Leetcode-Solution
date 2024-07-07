class Solution {
public:
    long long minimumCost(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        long long lt = find(nums,  nums.size() / 2, -1);
        long long rt = find(nums,  nums.size() / 2, 1);
        
        long long ls = calculate(nums, lt);
        long long rs = calculate(nums, rt);
        return min(ls, rs);
    }
    long long find(vector<int>& nums, int in, int direction) {
        long long med = nums[in];
        
        for (long long i = med; i >= 1 && i <= 1e9; i += direction) {
            if (isPalindrome(i))
                return i;
        }
        return -1;
    }
    bool isPalindrome(long long num) {
        string s = to_string(num);
        string t = s;
        reverse(t.begin(), t.end());
        return s == t;
    }
    
    long long calculate(vector<int>& nums, long long p) {
        long long c = 0;
        for (int i = 0; i < nums.size(); i++) {
            c += abs(nums[i] - p);
        }
        return c;
    }
};
