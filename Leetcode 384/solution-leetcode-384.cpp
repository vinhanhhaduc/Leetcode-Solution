class Solution {
public:
    vector<int> k;
    int length;

    Solution(vector<int>& nums) {
        k = nums;
        length = nums.size();
    }
    
    vector<int> reset() {
        return k;
    }
    
    vector<int> shuffle() {
        vector<int> a;
        a = k;
        int t = length, temp;
        for (int i = 0; i < t - 1; i++) {
            int random = rand()%t;
            temp = a[random];
            a[random] = a[i];
            a[i] = temp;
        }
        return a;

    }
};

