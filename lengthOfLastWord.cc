class Solution {
public:
    int lengthOfLastWord(string s) {
        size_t c = s.find_last_not_of(" ");
        if (c != s.size() - 1) {
            s = s.substr(0, c + 1);
        }
        size_t space = s.find_last_of(" ");
        if (space == string::npos) {
            return s.size();
        }
        return (s.substr(space + 1)).size();
    }
};
