class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int n = nums.size();
        if (n <= 1) {
            return n;
        }
        int pos = 0;
        for (int i = 0; i < n-1; ++i) {
            if (nums[i] != nums[i+1]) {
                ++pos;
                nums[pos] = nums[i+1];
            }
        }
        return pos + 1;
    }
};