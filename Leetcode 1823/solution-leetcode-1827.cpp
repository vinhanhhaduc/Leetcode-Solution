class Solution {
public:
    int findTheWinner(int n, int k) {
        int res = 1;
        for (int i = 1; i <= n; i++) {
            res = (res + k) % i;
        }
        return res + 1;
    }
};