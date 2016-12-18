/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        ListNode *p = head;
        while (p != NULL) {
            if (p->next != NULL && p->val == p->next->val) {
                ListNode *del = p->next;
                p->next = del->next;
                delete del;
            } else {
                p = p->next;
            }
        }
        return head;
    }
};
