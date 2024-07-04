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
    ListNode* mergeNodes(ListNode* head) {
        ListNode* d = new ListNode(0);
        d->next = head;
        ListNode* prev = d;

        while (head) {
            head = head->next;
            int sum = 0;
            while (head && head->val != 0) {
                sum += head->val;
                head = head->next;
            }
            if (head) {
                prev->next = new ListNode(sum);
                prev = prev->next;
            }
        }

        return d->next;
    }

};