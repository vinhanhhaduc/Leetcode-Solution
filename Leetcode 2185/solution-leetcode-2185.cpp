class Solution {
public:
    int prefixCount(vector<string>& words, string pref) {
        int n = pref.size();
        int cnt = 0;
        for (int i = 0; i < words.size(); i++) {
            string s = words[i].substr(0, n);
            if (s == pref) {
                cnt++;
            }
        }
        return cnt;
    }
};