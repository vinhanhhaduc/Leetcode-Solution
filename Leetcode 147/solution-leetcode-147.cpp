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
    ListNode* insertionSortList(ListNode* head) {
        if (!head || !head->next) {
            return head;
        }
        
        ListNode *d = new ListNode(0);
        d->next = head;
        
        ListNode *curr = head->next;
        ListNode *prev = head;
        
        while (curr) {
            if (curr->val < prev->val) {
                ListNode *temp = d;
                while (temp->next->val < curr->val) {
                    temp = temp->next;
                }
                prev->next = curr->next;
                curr->next = temp->next;
                temp->next = curr;
                curr = prev->next;
            } else {
                prev = curr;
                curr = curr->next;
            }
        }
        
        ListNode *res = d->next;
        delete d;
        return res;
    }
};