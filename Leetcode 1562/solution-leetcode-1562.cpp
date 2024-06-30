class Solution {
public:
    int findLatestStep(vector<int>& arr, int m) {
        vector<int> length(arr.size()+2, 0);
        vector<int> count(arr.size()+1, 0);
        int res = -1;

        for (int i = 0; i < arr.size(); i++) {
            int pos = arr[i];
            int left = length[pos - 1];
            int right = length[pos + 1];
            int newl = left + right + 1;
            if (left > 0) count[left]--;
            if (right > 0) count[right]--;

            count[newl]++;

            length[pos - left] = newl;
            length[pos + right] = newl;

            if (count[m] > 0) {
                res = i + 1;
            }
        }
        return res;
    }
};
