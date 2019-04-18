#include <iostream>
#include <vector>
#include <map>
using namespace std;
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> result;
        map<int, int> m;
        for (int i = 0; i < nums.size(); i++) {
			int diff = target - nums[i];
			auto it = m.find(diff);
			if (it != m.end()) {
				result.push_back(it->second);
				result.push_back(i);
				return result;
			}
			m[nums[i]] = i;
        }
		return result;
    }
};

int main() {
	Solution a = Solution();
	vector<int> v = {3, 3};
	auto vec = a.twoSum(v, 6);
	for (auto i : vec) {
		cout << i << endl;
	}
	return 0;
}