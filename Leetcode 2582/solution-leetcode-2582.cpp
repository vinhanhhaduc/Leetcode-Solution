class Solution {
public:
    int passThePillow(int n, int time) {
        int i = 1;
        bool f = true;
        while (time--) {
            if (f) {
                i++;
            } else {
                i--;
            }
            if (i == n) {
                f = false;
            }
            if (i == 1) {
                f = true;
            }
        }
        return i;
    }
};