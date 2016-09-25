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
    ListNode* swapPairs(ListNode* head) {
        ListNode *p = head;
        while (p != NULL && p->next != NULL) {
            ListNode *q = p->next;
            if (p == head) {
                p->next = q->next;
                q->next = p;
                head = q;
            } else if (q->next != NULL) {
                ListNode *r = q->next;
                p->next = r;
                q->next = r->next;
                r->next = q;
                p = q;
            } else {
                p = p->next;
            }
        }
        return head;
    }
};