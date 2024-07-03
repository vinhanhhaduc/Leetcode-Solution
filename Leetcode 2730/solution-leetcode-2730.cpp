class Solution {
public:
    int longestSemiRepetitiveSubstring(string s) {
        int i = 0, j = 0, n = s.length();
        int cnt = 0;
        char val;
        int res = 0;
        
        while (j < n) {
            if (j > 0 && s[j] == s[j - 1]) {
                cnt++;
                if (cnt == 1) {
                    val = s[j];
                }
            }
            
            if (cnt > 1) {
                while (i <= j && cnt > 1) {
                    if (s[i] == s[i + 1] && s[i] == val) {
                        cnt--;
                    }
                    i++;
                }
                val = s[j];
            }
            
            res = max(res, j - i + 1);
            j++;
        }
        
        return res;
    }
};
