class Solution {
public:
    bool isPalindrome(int x) {
        vector<int> v;
        int n = 0;
        if (x < 0) {
            return false;
        }
        while (x != 0) {
            n = x % 10;
            x /= 10;
            v.push_back(n);
        }
        int low = 0;
        int high = v.size() - 1;
        bool result = false;
        while (low < high) {
            if (v[low] == v[high]) {
                ++low;
                --high;
            } else {
                break;
            }
        }
        if (low >= high) {
            result = true;
        }
        return result;
    }
};