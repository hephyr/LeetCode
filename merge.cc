class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        vector<int> v;
        int i = 0;
        int p = 0, q = 0;
        while (p < m && q < n) {
            if (nums1[p] > nums2[q]) {
                v.push_back(nums2[q]);
                ++q;
            } else {
                v.push_back(nums1[p]);
                ++p;
            }
            ++i;
        }
        if (p == m) {
            while (i < m+n) {
                v.push_back(nums2[q]);
                ++q;
                ++i;
            }
        } else if (q == n) {
            while (i < m+n) {
                v.push_back(nums1[p]);
                ++p;
                ++i;
            }
        }
        nums1 = v;
    }
};
