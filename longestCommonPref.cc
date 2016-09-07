class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string result;
        if (strs.size() < 1) {
            return result;
        }
        string temp = strs[0];
        for (int i = 1; i <= temp.size(); ++i) {
            string w = temp.substr(0, i);
            bool match = true;
            for (auto s : strs) {
                if (i > s.size() || s.substr(0, i) != w) {
                    match = false;
                    break;
                }
            }
            if (!match) {
                break;
            }
            result = w;
        }
        return result;
    }
};