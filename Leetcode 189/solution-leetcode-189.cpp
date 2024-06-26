class Solution {
public:
    void dao(vector<int>& nums, int i, int j) {
        while (i < j) {
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
            i++;
            j--;
        }
    }
    void rotate(vector<int>& nums, int k) {
        int length = nums.size();
        k = k % length;
        dao(nums, 0, length - 1);
        dao(nums, 0, k - 1);
        dao(nums, k, length - 1);
    }
};
