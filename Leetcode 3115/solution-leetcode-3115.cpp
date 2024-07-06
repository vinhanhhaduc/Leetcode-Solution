class Solution {
public:
    bool isPrime(int n) {
        if(n <= 1) {
            return false;
        }
        for (int i = 2; i <= sqrt(n); i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return true;
    }
    int maximumPrimeDifference(vector<int>& nums) {
        int maxn = INT_MIN, minn = INT_MAX;

        for (int i = 0; i < nums.size(); ++i) {
            if (isPrime(nums[i])) {
                maxn = max(maxn, i);
                minn = min(minn, i);
            }
        }

        return abs(maxn - minn);
    }
};