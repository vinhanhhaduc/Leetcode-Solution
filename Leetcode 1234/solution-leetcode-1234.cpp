class Solution {
public:
    int balancedString(string s) {
        int count[128] = {0};
        int n = s.length();
        int res = n, i = 0, k = n / 4;
        for (int j = 0; j < n; ++j) {
            count[s[j]]++;
        }
        for (int j = 0; j < n; ++j) {
            count[s[j]]--;
            while (i < n && count['Q'] <= k && count['W'] <= k && count['E'] <= k && count['R'] <= k) {
                res = min(res, j - i + 1);
                count[s[i++]] += 1;
            }
        }
        
        return res;
    }

};