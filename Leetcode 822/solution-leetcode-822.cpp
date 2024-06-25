class Solution {
public:
    int flipgame(vector<int>& fronts, vector<int>& backs) {
        int res = INT_MAX;
        int arr[2001];
        for (int i = 0; i < 2001; i++) {
            arr[i] = 0;
        } 
        for (int i = 0; i < fronts.size(); i++) {
            if (fronts[i] == backs[i]) {
                arr[fronts[i]] = 1;
            }
        }
        for(int i = 0; i < fronts.size(); i++) {
            if(arr[fronts[i]] == 0) {
                res = min(res, fronts[i]);
            }
            if(arr[backs[i]] == 0) {
                res = min(res, backs[i]);
            }
        }
        if (res == INT_MAX) {
            return 0;
        } else {
            return res;
        }
    }
};
