class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
        sort(nums.begin(), nums.end());
        int n = nums.size();
        for (int i = 0; i < n-1; ++i) {
            int low = i + 1;
            int high = n - 1;
            if (i > 0 && nums[i] == nums[i-1])   continue;
            while (low < high) {
                int a = nums[i];
                int b = nums[low];
                int c = nums[high];
                if (a + b + c == 0) {
                    vector<int> v;
                    v.push_back(a);
                    v.push_back(b);
                    v.push_back(c);
                    result.push_back(v);
                    while (nums[low] == nums[low + 1])
                        low++;
                    while (nums[high] == nums[high -1])
                        high--;
                    low++;
                    high--;
                } else if (a + b + c > 0) {
                    while (nums[high] == nums[high-1])
                        high--;
                    high--;
                } else if (a + b + c < 0) {
                    while (nums[low] == nums[low + 1])
                        low++;
                    low++;
                }
            }
        }
        return result;
    }
};