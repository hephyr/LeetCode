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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode *p = head;
        int count = 0;
        while (p != NULL) {
            ++count;
            p = p->next;
        }
        p = head;
        ListNode *pre;
        while (p != NULL && count != n) {
            --count;
            pre = p;
            p = p->next;
        }
        ListNode *del = p;
        if (p != head) {
            pre->next = p->next;
        } else {
            head = p->next;
        }
        delete del;
        del = NULL;
        return head;
    }
};
