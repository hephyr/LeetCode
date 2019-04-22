#include <iostream>
#include <string>
#include <map>
using namespace std;
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int result = 0;
        map<char, int> m;
        int start = 0;
        for (int i = 0; i < s.size(); ++i) {
            auto it = m.find(s[i]);
            if (it != m.end() && it->second >= start) {
                start = it->second + 1;
            }
            int temp = i - start + 1;
            result = result < temp ? temp : result;
            m[s[i]] = i;
            if (s.size() - start + 1 <= result) break;
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