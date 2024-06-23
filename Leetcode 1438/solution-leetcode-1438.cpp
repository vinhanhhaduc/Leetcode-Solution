#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    int longestSubarray(vector<int>& arr, int limit) {
        int i = 0, j = 0, res = 0;
        vector<int> min_nums; 
        vector<int> max_nums;

        for (; i < arr.size(); i++) {
            while (!min_nums.empty() && arr[min_nums.back()] > arr[i])
                min_nums.pop_back();
            min_nums.push_back(i);


            while (!max_nums.empty() && arr[max_nums.back()] < arr[i])
                max_nums.pop_back();
            max_nums.push_back(i);
            while (!min_nums.empty() && !max_nums.empty() && arr[max_nums.front()] - arr[min_nums.front()] > limit) {
                if (min_nums.front() == j)
                    min_nums.erase(min_nums.begin());
                if (max_nums.front() == j)
                    max_nums.erase(max_nums.begin());
                j++;
            }

            res = max(res, i - j + 1);
        }
        return res;
    }
};

