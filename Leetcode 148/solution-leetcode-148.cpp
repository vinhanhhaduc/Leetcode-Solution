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
    ListNode* sortList(ListNode* head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        
        ListNode* slow = head;
        ListNode* fast = head->next;
        
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
        }
        
        ListNode* s = slow->next;
        slow->next = nullptr;
        ListNode* s1 = sortList(head);
        ListNode* s2 = sortList(s);
        
        return merge(s1, s2);
    }
private:
    ListNode* merge(ListNode* l1, ListNode* l2) {
        ListNode *head = new ListNode(0);
        ListNode *current = head;
        
        while (l1 != nullptr && l2 != nullptr) {
            if (l1->val < l2->val) {
                current->next = l1;
                l1 = l1->next;
            } else {
                current->next = l2;
                l2 = l2->next;
            }
            current = current->next;
        }
        
        if (l1 != nullptr) {
            current->next = l1;
        }
        
        if (l2 != nullptr) {
            current->next = l2;
        }
        
        return head->next;
    }
};
