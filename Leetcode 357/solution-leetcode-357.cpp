class Solution {
public:
    int countNumbersWithUniqueDigits(int n) {
        if (n == 0) {
            return 1;
        }
        
        int res = 1;
        for (int i = 1; i <= n; i++) {
            if (i == 1) {
                res += 9;
            } else if (i == 2) {
                res += 9 * 9;
            } else {
                int temp = 9;
                for (int j = 2; j <= i; j++) {
                    temp *= (11 - j);
                }
                res += temp;
            }
        }
        return res;
    }
};