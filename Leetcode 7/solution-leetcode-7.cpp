class Solution {
public:
    int reverse(int x) {
        long long res = 0;
        long long t = 0;
        while (x != 0) {
            t = x % 10;
            res = res * 10 + t;
            x /= 10 ;
        }
        if (res > INT_MAX || res < INT_MIN) {
            return 0;
        }
        return res;
    }
};
