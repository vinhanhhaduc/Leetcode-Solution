class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int c[256] = {0};
        int l = 0;
        int r = 0;
        int m = 0;
        
        while (l < s.length()) {
            int val = s[l];
            if (c[val] > r) {
                r = c[val];
            }
            m = max(m, l - r + 1);
            c[val] = l + 1;
            l++;
        }
        
        return m;
    }
};
