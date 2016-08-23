class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        vector<int> nums3;
        double median = 0;
        int v1 = 0;
        int v2 = 0;
        while (v1 < nums1.size() && v2 < nums2.size()) {
            if (nums1[v1]  < nums2[v2]) {
                nums3.push_back(nums1[v1]);
                ++v1;
            } else {
                nums3.push_back(nums2[v2]);
                ++v2;
            }
        }
        
        if (v1 < nums1.size()) {
            vector<int>::iterator iter = nums1.begin() + v1;
            nums3.insert(nums3.end(), iter, nums1.end());
        } else if (v2 < nums2.size()) {
            vector<int>::iterator iter = nums2.begin() + v2;
            nums3.insert(nums3.end(), iter, nums2.end());
        }
        if (nums3.size() % 2) {
            median = nums3[nums3.size() / 2];
        } else {
            int low = nums3.size() / 2 - 1;
            int high = nums3.size() / 2;
            median = (nums3[low] + nums3[high]) / 2.0;
        }
        for (auto a : nums3) {
            cout << a << endl;
        }
        return median;
    }
};