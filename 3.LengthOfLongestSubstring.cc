#include <iostream>
#include <string>
#include <map>
using namespace std;
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int result = 0;
        string unique;
        for (auto c : s) {
            auto pos = unique.find_first_of(c);
            if (pos != unique.npos) {
                unique.erase(0, pos+1);
            }
            unique.push_back(c);
            if (unique.size() > result) {
                result = unique.size();
            }
        }
        return result;
    }
};

int main() {
    Solution s;
    cout << s.lengthOfLongestSubstring("nfvbiywbasfbu") << endl;
    cout << s.lengthOfLongestSubstring("abcabc") << endl;
    return 0;
}