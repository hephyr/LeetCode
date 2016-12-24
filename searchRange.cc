class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        vector<int> result = {-1, -1};
        int low = 0;
        int high = nums.size() - 1;
        int middle = high / 2;
        while (low < high && nums[middle] != target) {
            if (nums[middle] < target) {
                low = middle + 1;
            } else {
                high = middle - 1;
            }
            middle = (low + high) / 2;
        }
        if (nums[middle] != target) {
            return result;
        } else {
            result[0] = result[1] = middle;
            low = middle - 1;
            high = middle + 1;
            while (nums[low] == target && low >= 0) {
                result[0] = low;
                --low;
            }
            while (nums[high] == target && high < nums.size()) {
                result[1] = high;
                ++high;
            }
        }
        return result;
    }
};
