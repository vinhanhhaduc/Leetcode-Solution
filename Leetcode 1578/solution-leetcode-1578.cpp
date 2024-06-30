class Solution {
public:
    int minCost(string colors, vector<int>& neededTime) {
        int min = 0;
        for (int i = 0; i < colors.size() -  1; i++) {
            if (colors[i] == colors[i + 1]) {
                if (neededTime[i] <= neededTime[i + 1]) {
                    min += neededTime[i];
                } else{
                    min += neededTime[i+1];
                    neededTime[i + 1] = neededTime[i];
                    colors[i + 1] = colors[i];
                }
            }
        }
        return min;
    }
};
