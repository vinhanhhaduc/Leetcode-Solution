class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        vector<int> res;
        unordered_map<int, int> k;
        for (int i = 0; i < nums1.size(); i++) {
            k[nums1[i]]++;
        }
        for (int i = 0; i < nums2.size(); i++) {
            if (k[nums2[i]]) {
                res.push_back(nums2[i]);
                k[nums2[i]]--;
            };
        }
        return res;
    }
};