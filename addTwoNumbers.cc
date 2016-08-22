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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
         ListNode *result = NULL;
         ListNode **p = &result;
         int val1 = 0, val2 = 0;
         int add = 0;
         while (l1 != NULL || l2 != NULL) {
             if (l1 != NULL) {
                 val1 = l1->val;
                 l1 = l1->next;
             } else {
                 val1 = 0;
             }
             if (l2 != NULL) {
                 val2 = l2->val;
                 l2 = l2->next;
             } else {
                 val2 = 0;
             }
             add = add + val1 + val2;
             ListNode *q = new ListNode(add%10);
             *p = q;
             p = &(q->next);
             add /= 10;
         }
         if (add > 0) {
             ListNode *q = new ListNode(add%10);
             *p = q;
             p = &(q->next);
         }
         return result;
    }
};