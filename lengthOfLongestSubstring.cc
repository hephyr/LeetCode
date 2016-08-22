class Solution {
public:
    // first write
    // int lengthOfLongestSubstring(string s) {
    //     string max, temp;
    //     for(auto c : s) {
    //         for (auto t : temp) {
    //             if (t == c) {
    //                 if (temp.size() > max.size()) {
    //                     max = temp;
    //                 }
    //                 auto a = temp.find(t);
    //                 temp.erase(0, a+1);
    //                 break;
    //             }
    //         }
    //         temp.push_back(c);
    //     }
    //     if (temp.size() > max.size()) {
    //         max = temp;
    //     }
    //     return max.size();
    // }
    // second write
    int lengthOfLongestSubstring(string s) {
        map<char, int> m;
        int maxSize = 0;
        int lastRepeatPos = -1;
        for (int i = 0; i < s.size(); ++i) {
            if (m.find(s[i]) != m.end() && lastRepeatPos < m[s[i]]) {
                lastRepeatPos = m[s[i]];
            }
            int len = i - lastRepeatPos;
            if (len > maxSize) {
                maxSize = len;
            }
            m[s[i]] = i;
        }
        return maxSize;
    }
};