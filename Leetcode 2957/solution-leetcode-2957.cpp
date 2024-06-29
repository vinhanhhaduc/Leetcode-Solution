class Solution {
public:
    int removeAlmostEqualCharacters(string word) {
        int res = 0;
        for(int i = 1 ; i < word.size() ; i++){
            if(word[i] == word[i - 1] || word[i] + 1 == word[i - 1] || word[i - 1] + 1 == word[i]){
                res++;
                i++;
            }
        }
        return res;
    }
};