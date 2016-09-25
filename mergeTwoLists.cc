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
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *result = NULL;
        ListNode *p1, *p2, *tail;
        p1 = l1;
        p2 = l2;
        int val1;
        int val2;
        while (p1 != NULL || p2 != NULL) {
            val1 = (p1 != NULL) ? (p1->val) : 0X7FFFFFFF;
            val2 = (p2 != NULL) ? (p2->val) : 0X7FFFFFFF;
            int val;
            if (val1 < val2) {
                val = val1;
                p1 = p1->next;
            } else {
                val = val2;
                p2 = p2->next;
            }
            ListNode *p = new ListNode(val);
            if (result == NULL) {
                result = p;
                tail = result;
            } else {
                tail->next = p;
                tail = p;
            }
        }
        return result;
    }
};