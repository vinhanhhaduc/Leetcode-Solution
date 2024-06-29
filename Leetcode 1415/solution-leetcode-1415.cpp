class Solution {
public:
    string getHappyString(int n, int k) {
        vector<string> happyStrings;
        string s;
        vector<pair<string, int>> stack;
        stack.push_back({"", 0});

        while (!stack.empty()) {
            auto [s, len] = stack.back();
            stack.pop_back();

            if (len == n) {
                happyStrings.push_back(s);
                continue;
            }

            for (char ch : {'c', 'b', 'a'}) {
                if (s.empty() || s.back() != ch) {
                    stack.push_back({s + ch, len + 1});
                }
            }
        }

        if (happyStrings.size() < k) {
            return "";
        }
        
        sort(happyStrings.begin(), happyStrings.end());
        return happyStrings[k - 1];
    }
};