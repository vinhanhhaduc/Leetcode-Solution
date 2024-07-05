/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    vector<int> nodesBetweenCriticalPoints(ListNode* head) {
        vector<int> res;
        ListNode* prev = head;
        head = head->next;

        if (!head || !head->next || !head->next->next) {
            return {-1, -1};
        }

        int i = 1;
        while (head->next != NULL) {
            if ((prev->val < head->val && head->val > head->next->val) ||
                (prev->val > head->val && head->val < head->next->val)) {
                res.push_back(i);
            }
            prev = head;
            head = head->next;
            i++;
        }

        if (res.size() < 2) {
            return {-1, -1};
        }

        int minn = INT_MAX;
        int maxn = res.back() - res.front();
        
        for (int i = 1; i < res.size(); i++) {
            minn = min(minn, res[i] - res[i - 1]);
        }

        return {minn, maxn};
    }
};