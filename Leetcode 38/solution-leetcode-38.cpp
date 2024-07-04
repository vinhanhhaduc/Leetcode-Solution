class Solution {
public:
    string countAndSay(int n) {
        if (n == 1) return "1";
        string s = countAndSay(n - 1);
        string res;
        int cnt = 1;
        char c = s[0];
        for (int i = 1; i < s.length(); ++i) {
            if (s[i] == c) {
                cnt++;
            } else {
                res += to_string(cnt) + c;
                c = s[i];
                cnt = 1;
            }
        }
        res += to_string(cnt) + c;
        return res;
    }
};