class Solution {
public:
    int numWaterBottles(int numBottles, int numExchange) {
        int cnt = 0, k = 0;
        while (numBottles != 0) {
            cnt += numBottles;
            k += numBottles;
            numBottles = k / numExchange;
            k %= numExchange;

        }
        return cnt; 
    }
};