class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        string max, temp;
        for(auto c : s) {
            for (auto t : temp) {
                if (t == c) {
                    if (temp.size() > max.size()) {
                        max = temp;
                    }
                    auto a = temp.find(t);
                    temp.erase(0, a+1);
                    break;
                }
            }
            temp.push_back(c);
        }
        if (temp.size() > max.size()) {
            max = temp;
        }
        return max.size();
    }
};