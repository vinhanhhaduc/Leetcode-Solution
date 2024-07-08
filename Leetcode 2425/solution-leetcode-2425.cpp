class Solution {
public:
    int xorAllNums(vector<int>& nums1, vector<int>& nums2) {
        int res1 = 0, res2 = 0;
        for (int i = 0; i < nums1.size(); i++) {
            res1 ^= nums1[i];
        }
        for (int j = 0; j < nums2.size(); j++) {
            res2 ^= nums2[j];
        }
        if(nums1.size() % 2 == 0 &&  nums2.size()%2 == 0){
            return 0;
        }
        else if(nums1.size() % 2 && nums2.size() % 2 == 0){
            return res2;
        }
        else if (nums1.size() % 2 == 0 &&  nums2.size() % 2){
            return res1;
        }
        return res1^res2;
    }
};