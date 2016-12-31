class Solution {
public:
    bool canJump(vector<int>& nums) {
        int last = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (i > last) {
                break;
            } else if (i + nums[i] > last) {
                last = i + nums[i];
            }
        }
        return last >= nums.size() -1;
    }
};
