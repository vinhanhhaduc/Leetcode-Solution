class Solution {
public:
    int minTimeToType(string word) {
        int time = 0;
        char chara = 'a';
        for (int i = 0; i < word.length(); i++) {
            int n = abs(word[i] - chara);
            time += min(n, 26 - n);
            chara = word[i];
        }
        return time + word.length();
    }
};